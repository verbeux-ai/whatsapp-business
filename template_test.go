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
