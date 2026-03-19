package whatsapp_business_test

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	whatsapp_business "github.com/verbeux-ai/whatsapp-business"
)

func TestClient_ListTemplatesWithoutFilter(t *testing.T) {
	result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{})
	require.NoError(t, err)
	require.NotEmpty(t, result.Data)
}

func TestClient_ListTemplatesWithLimit(t *testing.T) {
	result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
		Limit: 1,
	})
	require.NoError(t, err)
	require.Len(t, result.Data, 1)
}

func TestClient_ListTemplatesWithAfterCursor(t *testing.T) {
	result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
		Limit: 1,
	})

	require.NoError(t, err)
	require.Len(t, result.Data, 1)
	if result.Paging.Cursors.After == "" {
		t.Skip("Not enough templates to test pagination, skipping.")
	}
	require.NotEmpty(t, result.Paging.Cursors.After)

	result, err = client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
		Limit: 1,
		After: result.Paging.Cursors.After,
	})

	require.NoError(t, err)
	require.Len(t, result.Data, 1)
}

func TestClient_ListTemplatesWithCategory(t *testing.T) {
	result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
		Category: whatsapp_business.MarketingCategoryFilter,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result.Data)

	for _, template := range result.Data {
		require.Equal(t, string(whatsapp_business.MarketingCategory), template.Category)
	}
}

func TestClient_GetTemplate(t *testing.T) {
	result, err := client.GetTemplate(t.Context(), "auth", "pt_BR")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestClient_CreateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("test_template_%d", time.Now().UnixNano())

	createReq := whatsapp_business.CreateTemplateRequest{
		Name:     templateName,
		Category: whatsapp_business.MarketingCategory,
		Language: "en_US",
		Components: []whatsapp_business.TemplateDataComponent{
			{
				Type: "BODY",
				Text: "Test body",
			},
		},
	}

	createResp, err := client.CreateTemplate(t.Context(), createReq)
	require.NoError(t, err)
	require.NotEmpty(t, createResp)
	require.NotEmpty(t, createResp.Id)
}

func TestClient_CreateTemplateWithNamedParameters(t *testing.T) {
	templateName := fmt.Sprintf("test_template_named_%d", time.Now().UnixNano())

	createReq := whatsapp_business.CreateTemplateRequest{
		Name:            templateName,
		Category:        whatsapp_business.MarketingCategory,
		Language:        "en_US",
		ParameterFormat: whatsapp_business.NamedParameterFormat,
		Components: []whatsapp_business.TemplateDataComponent{
			{
				Type: "BODY",
				Text: "Hello {{first_name}}, your order is confirmed.",
				Example: &whatsapp_business.TemplateExample{
					BodyTextNamedParams: []whatsapp_business.NamedParam{
						{
							ParamName: "first_name",
							Example:   "John",
						},
					},
				},
			},
		},
	}

	createResp, err := client.CreateTemplate(t.Context(), createReq)
	require.NoError(t, err)
	require.NotEmpty(t, createResp)
	require.NotEmpty(t, createResp.Id)
}

func TestClient_CreateTemplateWithImageHeader(t *testing.T) {
	// Create a temporary file for the test
	file, err := os.Create("test_image.png")
	require.NoError(t, err)
	// Decode the base64 string to bytes
	sEnc := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	require.NoError(t, err)
	_, err = file.Write(sDec)
	require.NoError(t, err)
	file.Close()

	t.Cleanup(func() {
		os.Remove(file.Name())
	})

	templateName := fmt.Sprintf("test_template_image_%d", time.Now().UnixNano())

	uploadResp, err := client.UploadPermanentImage(t.Context(), file.Name())
	require.NoError(t, err)
	require.NotEmpty(t, uploadResp.Handle)

	createReq := whatsapp_business.CreateTemplateRequest{
		Name:     templateName,
		Category: whatsapp_business.MarketingCategory,
		Language: "en_US",
		Components: []whatsapp_business.TemplateDataComponent{
			{
				Type:   "HEADER",
				Format: "IMAGE",
				Example: &whatsapp_business.TemplateExample{
					HeaderHandle: []string{uploadResp.Handle},
				},
			},
			{
				Type: "BODY",
				Text: "Check out our new product!",
			},
		},
	}

	createResp, err := client.CreateTemplate(t.Context(), createReq)
	require.NoError(t, err)
	require.NotEmpty(t, createResp)
	require.NotEmpty(t, createResp.Id)
}

func TestClient_UpdateTemplate(t *testing.T) {
	templateName := os.Getenv("APPROVED_TEMPLATE_NAME_TEST")
	if templateName == "" {
		t.Skip("APPROVED_TEMPLATE_NAME_TEST is not set, skipping.")
	}

	template, err := client.GetTemplate(t.Context(), templateName, "en_US")
	require.NoError(t, err)
	require.NotEmpty(t, template)

	updateReq := whatsapp_business.UpdateTemplateRequest{
		Category: whatsapp_business.MarketingCategory,
	}

	updateResp, err := client.UpdateTemplate(t.Context(), template.Id, updateReq)
	require.NoError(t, err)
	require.NotEmpty(t, updateResp)
}

func TestClient_DeleteTemplate(t *testing.T) {
	templateName := fmt.Sprintf("test_template_%d", time.Now().UnixNano())
	createReq := whatsapp_business.CreateTemplateRequest{
		Name:     templateName,
		Category: whatsapp_business.UtilityCategory,
		Language: "en_US",
		Components: []whatsapp_business.TemplateDataComponent{
			{
				Type: "BODY",
				Text: "Test body",
			},
		},
	}

	createResp, err := client.CreateTemplate(t.Context(), createReq)
	require.NoError(t, err)
	require.NotEmpty(t, createResp)
	require.NotEmpty(t, createResp.Id)

	deleteResp, err := client.DeleteTemplate(t.Context(), templateName)
	require.NoError(t, err)
	require.True(t, deleteResp.Success)
}
