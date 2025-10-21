package whatsapp_business_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetMedia(t *testing.T) {
	result, err := client.GetMedia(t.Context(), os.Getenv("MEDIA_ID_TEST"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.ID)
	require.NotEmpty(t, result.URL)
	require.NotEmpty(t, result.Sha256)
	require.NotEmpty(t, result.FileSize)
}

func TestClient_DownloadMedia(t *testing.T) {
	result, err := client.GetMedia(t.Context(), os.Getenv("MEDIA_ID_TEST"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.ID)
	require.NotEmpty(t, result.URL)
	require.NotEmpty(t, result.Sha256)
	require.NotEmpty(t, result.FileSize)

	media, err := client.DownloadMedia(t.Context(), result.URL)
	require.NoError(t, err)
	require.NotEmpty(t, media)

	f, err := os.Create(media.GuessFileName)
	require.NoError(t, err)
	defer f.Close()
	_, err = io.Copy(f, media.Body)
	require.NoError(t, err)
}
