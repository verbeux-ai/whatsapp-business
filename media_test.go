package whatsapp_business_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetMedia(t *testing.T) {
	ctx := context.Background()
	result, err := client.GetMedia(ctx, os.Getenv("MEDIA_ID_TEST"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.ID)
	require.NotEmpty(t, result.URL)
	require.NotEmpty(t, result.Sha256)
	require.NotEmpty(t, result.FileSize)
}
