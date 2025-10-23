package whatsapp_business_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	whatsapp_business "github.com/verbeux-ai/whatsapp-business"
)

var client *whatsapp_business.Client

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")

	client = whatsapp_business.NewClient(
		whatsapp_business.WithToken(os.Getenv("TOKEN")),
		whatsapp_business.WithPhoneNumberId(os.Getenv("PHONE_ID")),
		whatsapp_business.WithBusinessId(os.Getenv("WA_BUSINESS_ID")),
	)

	os.Exit(m.Run())
}
