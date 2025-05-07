package whatsapp_business_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetBusinessProfile(t *testing.T) {
	ctx := context.Background()
	result, err := client.GetBusinessProfile(ctx, os.Getenv("PHONE_ID"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
