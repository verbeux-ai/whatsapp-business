package whatsapp_business_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	whatsapp_business "github.com/verbeux-ai/whatsapp-business"
)

func TestSendTextMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTextMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.TextMessage{
		PreviewUrl: nil,
		Body:       os.Getenv("TEXT_TEST_CONTENT"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTemplateMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTemplateMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.TemplateMessage{
		Name: "hello_world",
		Language: whatsapp_business.TemplateLanguageCode{
			Code: "en_US",
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTextMessageWithOption(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTextMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.TextMessage{
		Body: os.Getenv("TEXT_TEST_CONTENT"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Greater(t, len(result.Messages), 0)

	result2, err2 := client.SendTextMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.TextMessage{
		Body: os.Getenv("TEXT_TEST_CONTENT"),
	}, whatsapp_business.WithReplyMessage(result.Messages[0].Id))
	require.NoError(t, err2)
	require.NotEmpty(t, result2)
}

func TestSendImageMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendImageMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.ImageMessage{
		Link:    os.Getenv("IMAGE_TEST_CONTENT"),
		Caption: "Test caption",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendAudioMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendAudioMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.AudioMessage{
		Link: os.Getenv("AUDIO_TEST_CONTENT"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendDocumentMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendDocumentMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.DocumentMessage{
		Link:     os.Getenv("DOCUMENT_TEST_CONTENT"),
		Caption:  "Caption test",
		Filename: "Filename test",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendInteractiveMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendInteractiveMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.InteractiveMessage{
		Type: whatsapp_business.InteractiveMessageTypeList,
		Header: whatsapp_business.InteractiveMessageHeader{
			Type: whatsapp_business.InteractiveMessageHeaderTypeText,
			Text: "Test Header",
		},
		Body: whatsapp_business.InteractiveMessageBody{
			Text: "Test body text",
		},
		Footer: whatsapp_business.InteractiveMessageFooter{
			Text: "Test footer text",
		},
		Action: whatsapp_business.InteractiveMessageAction{
			Button: "Choose an option",
			Sections: []whatsapp_business.InteractiveMessageSection{
				{
					Title: "Section 1",
					Rows: []whatsapp_business.InteractiveMessageRow{
						{
							Id:          "row1",
							Title:       "Row 1",
							Description: "Row 1 description",
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
