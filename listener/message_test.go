package listener_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/whatsapp-business/listener"
)

func TestListener_OnTextMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnTextMessage(func(message *listener.TextMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.Message)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.Contacts)
		require.NotEmpty(t, message.Contacts[0].Name)
		require.NotEmpty(t, message.Contacts[0].WaID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(textMessage))))
	require.NoError(t, err)
}

func TestListener_OnAudioMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnAudioMessage(func(message *listener.AudioMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.Contacts)
		require.NotEmpty(t, message.Sha256)
		require.NotEmpty(t, message.Voice)
		require.NotEmpty(t, message.Contacts[0].Name)
		require.NotEmpty(t, message.Contacts[0].WaID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(audioMessage))))
	require.NoError(t, err)
}

func TestListener_OnImageMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnImageMessage(func(message *listener.ImageMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.Contacts)
		require.NotEmpty(t, message.Sha256)
		require.NotEmpty(t, message.Caption)
		require.NotEmpty(t, message.Contacts[0].Name)
		require.NotEmpty(t, message.Contacts[0].WaID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(imageMessage))))
	require.NoError(t, err)
}

func TestListener_OnDocumentMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnDocumentMessage(func(message *listener.DocumentMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.Contacts)
		require.NotEmpty(t, message.Sha256)
		require.NotEmpty(t, message.Caption)
		require.NotEmpty(t, message.Filename)
		require.NotEmpty(t, message.Contacts[0].Name)
		require.NotEmpty(t, message.Contacts[0].WaID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(documentMessage))))
	require.NoError(t, err)
}

func TestListener_OnStatusMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnStatusMessage(func(message *listener.StatusMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.Status)
		require.NotEmpty(t, message.Origin)
		require.NotEmpty(t, message.Origin.Type)
		require.NotEmpty(t, message.WaID)
		require.NotEmpty(t, message.ConversationID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(statusesMessage))))
	require.NoError(t, err)
}
