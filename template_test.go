package whatsapp_business_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_ListTemplates(t *testing.T) {
	result, err := client.ListTemplates(t.Context())
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestClient_GetTemplate(t *testing.T) {
	result, err := client.GetTemplate(t.Context(), "hello_world", "en_US")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
