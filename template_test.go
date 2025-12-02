package whatsapp_business_test

import (
	"testing"

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
		Category: whatsapp_business.MarketingCategory,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result.Data)

	for _, template := range result.Data {
		require.Equal(t, string(whatsapp_business.MarketingCategory), template.Category)
	}
}

func TestClient_GetTemplate(t *testing.T) {
	result, err := client.GetTemplate(t.Context(), "hello_world", "en_US")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
