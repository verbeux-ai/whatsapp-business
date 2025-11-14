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

func TestSendAudioWithIdMessage(t *testing.T) {
	mediaResult, err := client.UploadFromURL(t.Context(), whatsapp_business.UploadFromURL{
		URL: os.Getenv("AUDIO_TEST_CONTENT"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, mediaResult)
	require.NotEmpty(t, mediaResult.ID)

	ctx := context.Background()
	result, err := client.SendAudioMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.AudioMessage{
		ID:    mediaResult.ID,
		Voice: true,
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

func TestSendTemplateMessage(t *testing.T) {
	ctx := context.Background()

	template := whatsapp_business.TemplateMessage{
		Name: "template_sunne_1",
		Language: whatsapp_business.TemplateLanguageCode{
			Code: "pt_BR",
		}}

	result, err := client.SendTemplateMessage(ctx, os.Getenv("NUMBER"), template)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTemplateButtonMessage(t *testing.T) {
	ctx := context.Background()

	template := whatsapp_business.TemplateMessage{
		Name: "test_button",
		Language: whatsapp_business.TemplateLanguageCode{
			Code: "pt_BR",
		}}

	result, err := client.SendTemplateMessage(ctx, os.Getenv("NUMBER"), template)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTemplateMessageHeaderWithVariables(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTemplateMessage(
		ctx,
		os.Getenv("NUMBER"),
		whatsapp_business.TemplateMessage{
			Name: "test_header_param",
			Language: whatsapp_business.TemplateLanguageCode{
				Code: "pt_BR",
			},
		},
		whatsapp_business.WithHeaderTextNamed("version", "v1.2"),
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTemplateMessageHeaderBodyWithVariables(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTemplateMessage(
		ctx,
		os.Getenv("NUMBER"),
		whatsapp_business.TemplateMessage{
			Name: "test_text_with_variables",
			Language: whatsapp_business.TemplateLanguageCode{
				Code: "pt_BR",
			},
		},
		whatsapp_business.WithHeaderTextNamed("atendente", "Bumba meu boi"),
		whatsapp_business.WithBodyTextNamed("cliente", "Ivao da Massa"),
		whatsapp_business.WithBodyTextNamed("valor", "34 milhoes"),
		whatsapp_business.WithBodyTextNamed("data", "Dia 10"),
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTemplateMessageImage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTemplateMessage(
		ctx,
		os.Getenv("NUMBER"),
		whatsapp_business.TemplateMessage{
			Name: "test_image",
			Language: whatsapp_business.TemplateLanguageCode{
				Code: "pt_BR",
			},
		},
		whatsapp_business.WithHeaderImageLink(os.Getenv("IMAGE_TEST_CONTENT")),
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendTemplateMessageDocument(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTemplateMessage(
		ctx,
		os.Getenv("NUMBER"),
		whatsapp_business.TemplateMessage{
			Name: "test_document",
			Language: whatsapp_business.TemplateLanguageCode{
				Code: "pt_BR",
			},
		},
		whatsapp_business.WithHeaderDocumentLink(os.Getenv("DOCUMENT_TEST_CONTENT")),
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

//
//func TestSendTemplateMessageWithHeaderDocument(t *testing.T) {
//	ctx := context.Background()
//	result, err := client.SendTemplateMessage(
//		ctx,
//		os.Getenv("NUMBER"),
//		whatsapp_business.TemplateMessage{
//			Name: "test_document",
//			Language: whatsapp_business.TemplateLanguageCode{
//				Code: "pt_BR",
//			},
//		},
//		whatsapp_business.WithHeaderDocumentNamed(whatsapp_business.HeaderDocumentOptions{
//			Link:     os.Getenv("DOCUMENT_TEST_CONTENT"),
//			Filename: "document.pdf",
//		}),
//	)
//	require.NoError(t, err)
//	require.NotEmpty(t, result)
//}
//
//func TestSendTemplateMessageWithHeaderVideo(t *testing.T) {
//	ctx := context.Background()
//	result, err := client.SendTemplateMessage(
//		ctx,
//		os.Getenv("NUMBER"),
//		whatsapp_business.TemplateMessage{
//			Name: "test_video",
//			Language: whatsapp_business.TemplateLanguageCode{
//				Code: "pt_BR",
//			},
//		},
//		whatsapp_business.WithHeaderVideoNamed(whatsapp_business.HeaderVideoOptions{
//			Link: os.Getenv("VIDEO_TEST_CONTENT"),
//		}),
//	)
//	require.NoError(t, err)
//	require.NotEmpty(t, result)
//}
//
//func TestSendTemplateMessageWithHeaderImage(t *testing.T) {
//	ctx := context.Background()
//	result, err := client.SendTemplateMessage(
//		ctx,
//		os.Getenv("NUMBER"),
//		whatsapp_business.TemplateMessage{
//			Name: "test_image",
//			Language: whatsapp_business.TemplateLanguageCode{
//				Code: "pt_BR",
//			},
//		},
//		whatsapp_business.WithHeaderImageNamed(whatsapp_business.HeaderImageOptions{
//			Link: os.Getenv("IMAGE_TEST_CONTENT"),
//		}),
//	)
//	require.NoError(t, err)
//	require.NotEmpty(t, result)
//}

//func TestSendTemplateMessageWithHeaderText(t *testing.T) {
//	ctx := context.Background()
//	result, err := client.SendTemplateMessage(
//		ctx,
//		os.Getenv("NUMBER"),
//		whatsapp_business.TemplateMessage{
//			Name: "test_text_with_variables",
//			Language: whatsapp_business.TemplateLanguageCode{
//				Code: "pt_BR",
//			},
//		},
//		whatsapp_business.WithHeaderTextNamed(whatsapp_business.HeaderTextOptions{Text: "Davi"}),
//		whatsapp_business.WithBodyNamed(whatsapp_business.BodyOptions{
//			Params: []whatsapp_business.TemplateComponentParameter{
//				whatsapp_business.NewTextParameterNamed(whatsapp_business.TextParamOptions{Text: "Nome do Cliente"}),
//				whatsapp_business.NewTextParameterNamed(whatsapp_business.TextParamOptions{Text: "1000"}),
//				whatsapp_business.NewDateTimeParameterNamed(whatsapp_business.DateTimeParamOptions{Fallback: "22/10/2025"}),
//			},
//		}),
//	)
//	require.NoError(t, err)
//	require.NotEmpty(t, result)
//}

func TestSendInteractiveMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendInteractiveMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.InteractiveMessage{
		Type: whatsapp_business.InteractiveMessageTypeList,
		Header: &whatsapp_business.InteractiveMessageHeader{
			Type: whatsapp_business.InteractiveMessageHeaderTypeText,
			Text: "Test Header",
		},
		Body: &whatsapp_business.InteractiveMessageBody{
			Text: "Test body text",
		},
		Footer: &whatsapp_business.InteractiveMessageFooter{
			Text: "Test footer text",
		},
		Action: &whatsapp_business.InteractiveMessageAction{
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

func TestSendButtonMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendInteractiveMessage(ctx, os.Getenv("NUMBER"), whatsapp_business.InteractiveMessage{
		Type: whatsapp_business.InteractiveMessageTypeButton,
		Header: &whatsapp_business.InteractiveMessageHeader{
			Type: whatsapp_business.InteractiveMessageHeaderTypeText,
			Text: "Test Header",
		},
		Body: &whatsapp_business.InteractiveMessageBody{
			Text: "Test body text",
		},
		Footer: &whatsapp_business.InteractiveMessageFooter{
			Text: "Test footer text",
		},
		Action: &whatsapp_business.InteractiveMessageAction{
			Buttons: []whatsapp_business.InteractiveMessageSectionButton{
				{
					Type: whatsapp_business.InteractiveMessageSectionButtonReplyType,
					Reply: whatsapp_business.InteractiveMessageSectionButtonReply{
						Id:    "1",
						Title: "Teste",
					},
				},
				{
					Type: whatsapp_business.InteractiveMessageSectionButtonReplyType,
					Reply: whatsapp_business.InteractiveMessageSectionButtonReply{
						Id:    "2",
						Title: "Bumbar!",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
