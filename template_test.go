package whatsapp_business_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	whatsapp_business "github.com/verbeux-ai/whatsapp-business"
)

func TestClient_ListTemplates(t *testing.T) {
	t.Run("Without filter", func(t *testing.T) {
		result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{})
		require.NoError(t, err)
		require.NotEmpty(t, result.Data)
	})

	t.Run("With limit", func(t *testing.T) {
		result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
			Limit: 1,
		})
		require.NoError(t, err)
		require.Len(t, result.Data, 1)
	})

	t.Run("With after cursor", func(t *testing.T) {
		result, err := client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
			Limit: 1,
		})
		require.NoError(t, err)
		require.Len(t, result.Data, 1)
		require.NotEmpty(t, result.Paging.Cursors.After)

		result, err = client.ListTemplates(t.Context(), whatsapp_business.ListTemplateFilter{
			Limit: 1,
			After: result.Paging.Cursors.After,
		})

		require.NoError(t, err)
		require.Len(t, result.Data, 1)
	})
}

func TestClient_GetTemplate(t *testing.T) {
	result, err := client.GetTemplate(t.Context(), "hello_world", "en_US")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
