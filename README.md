# WhatsApp Business Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/verbeux-ai/whatsapp-business.svg)](https://pkg.go.dev/github.com/verbeux-ai/whatsapp-business)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A Go client for the [WhatsApp Cloud API](https://developers.facebook.com/docs/whatsapp/cloud-api). It provides typed helpers for sending messages, working with media and templates, managing WhatsApp Business resources, and decoding webhook payloads.

The client uses Graph API `v22.0` by default. Use `WithBaseUrl` if your application needs a different Graph API version or a test server.

## Features

- Send text, image, video, audio, document, interactive, and template messages.
- Reply to an existing message and mark messages as read.
- Upload media from a URL or an `io.Reader`, fetch media metadata, and download media.
- Create, list, retrieve, update, and delete message templates.
- Retrieve business, profile, and phone-number details; register a phone number; configure webhooks; and generate an access token.
- Decode incoming webhook payloads for text, media, button replies, statuses, history, contact sync, account updates, and message echoes.

## Requirements

- Go 1.22 or newer.
- A Meta app with the WhatsApp product configured.
- A WhatsApp access token and phone-number ID. Some operations also require a WhatsApp Business Account ID or Meta app ID.

See Meta's [Cloud API getting-started guide](https://developers.facebook.com/docs/whatsapp/cloud-api/get-started) to configure these credentials and the required permissions.

## Install

```bash
go get github.com/verbeux-ai/whatsapp-business
```

## Quick start

Create a client with the credentials required by the operations your application will use. Keep tokens in environment variables or a secrets manager—never commit them.

```go
package main

import (
	"context"
	"os"

	whatsapp "github.com/verbeux-ai/whatsapp-business"
)

func main() {
	client := whatsapp.NewClient(
		whatsapp.WithToken(os.Getenv("WHATSAPP_TOKEN")),
		whatsapp.WithPhoneNumberId(os.Getenv("WHATSAPP_PHONE_NUMBER_ID")),
		// Required for template-management operations.
		whatsapp.WithBusinessId(os.Getenv("WHATSAPP_BUSINESS_ACCOUNT_ID")),
		// Required by UploadPermanentImage.
		whatsapp.WithAppId(os.Getenv("META_APP_ID")),
	)

	_, err := client.SendTextMessage(
		context.Background(),
		"5511999999999",
		whatsapp.TextMessage{Body: "Olá! Enviada pela Cloud API."},
	)
	if err != nil {
		panic(err)
	}
}
```

Use a phone number in international format, with country code and no `+`, spaces, or punctuation.

### Send media or reply to a message

Pass a public media URL or a media ID returned by a previous upload. The `WithReplyMessage` option adds the Cloud API message context.

```go
response, err := client.SendImageMessage(ctx, "5511999999999", whatsapp.ImageMessage{
	Link:    "https://example.com/promo.png",
	Caption: "Confira a novidade",
})
if err != nil {
	return err
}

_, err = client.SendTextMessage(
	ctx,
	"5511999999999",
	whatsapp.TextMessage{Body: "Posso ajudar em algo mais?"},
	whatsapp.WithReplyMessage(response.Messages[0].Id),
)
```

For uploads, use `UploadFile` with any `io.Reader`, or `UploadFromURL` to stream a file from a URL. `GetMedia` returns the media URL and metadata; pass that URL to `DownloadMedia` to retrieve its response body.

### Send a template message

The template must already exist and be approved in Meta. Template options add named or positional parameters to the header, body, and URL buttons.

```go
_, err := client.SendTemplateMessage(
	ctx,
	"5511999999999",
	whatsapp.TemplateMessage{
		Name: "order_update",
		Language: whatsapp.TemplateLanguageCode{Code: "pt_BR"},
	},
	whatsapp.WithBodyTextNamed("customer_name", "Ana"),
	whatsapp.WithBodyTextNamed("order_id", "12345"),
)
```

Template management is available through `ListTemplates`, `GetTemplate`, `CreateTemplate`, `UpdateTemplate`, and `DeleteTemplate`. The `template` package also includes `NewTemplateMessageBuilder` for building a message from template metadata.

### Receive webhook events

The `listener` package decodes webhook POST bodies and dispatches typed callbacks. Register the handlers your application needs, then pass the request body to `ReadBodySync` (or `ReadBodyAsync`).

```go
func handleWebhook(w http.ResponseWriter, r *http.Request) {
	webhook := listener.NewMessageListener()

	webhook.OnTextMessage(func(message *listener.TextMessage) error {
		log.Printf("message from %s: %s", message.From, message.Message)
		return nil
	})

	if err := webhook.ReadBodySync(r.Body); err != nil {
		http.Error(w, "invalid webhook payload", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
```

Available handlers include `OnTextMessage`, `OnImageMessage`, `OnAudioMessage`, `OnDocumentMessage`, `OnButtonMessage`, `OnStatusMessage`, `OnHistoryMessage`, `OnContactSyncMessage`, `OnAccountUpdateMessage`, and `OnMessageEcho`.

Your HTTP handler remains responsible for Meta's webhook verification challenge and for validating request signatures before handing the body to the listener.

## Configuration reference

| Option | Used by |
| --- | --- |
| `WithToken` | All authenticated Cloud API requests |
| `WithPhoneNumberId` | Sending messages, media operations, and message sync |
| `WithBusinessId` | Template management |
| `WithAppId` | Permanent image uploads |
| `WithHttpClient` | Custom transport, timeouts, proxies, or tests |
| `WithBaseUrl` | Custom Graph API base URL or test server |

All client methods accept a `context.Context`; use it to apply deadlines and cancellation in production.

## Testing

Run the test suite with:

```bash
go test ./...
```

Some root-package tests exercise the real Cloud API and expect credentials and test fixtures in a local `.env` file. Keep `.env` private; it is ignored by Git.

## Contributing

Issues and pull requests are welcome. Please format Go changes with `gofmt` and run the relevant tests before submitting a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
