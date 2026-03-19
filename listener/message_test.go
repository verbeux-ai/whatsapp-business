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

func TestListener_OnQuickReplyMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnButtonMessage(func(message *listener.ButtonMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.Message)
		require.NotEmpty(t, message.Payload)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.Contacts)
		require.NotEmpty(t, message.Contacts[0].Name)
		require.NotEmpty(t, message.Contacts[0].WaID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(quickReplyMessage))))
	require.NoError(t, err)
}

func TestListener_OnListQuickReplyMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnButtonMessage(func(message *listener.ButtonMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.Message)
		require.NotEmpty(t, message.Payload)
		require.NotEmpty(t, message.ID)
		require.NotEmpty(t, message.Time)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.Contacts)
		require.NotEmpty(t, message.Contacts[0].Name)
		require.NotEmpty(t, message.Contacts[0].WaID)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(listMessage))))
	require.NoError(t, err)
}

func TestListener_OnHistoryMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnHistoryMessage(func(message *listener.HistoryMessage) error {
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.ToPhoneNumberId)
		require.NotEmpty(t, message.History)

		block := message.History[0]
		require.Equal(t, 1, block.Phase)
		require.Equal(t, 131, block.ChunkOrder)
		require.Equal(t, 30, block.Progress)
		require.NotEmpty(t, block.Threads)

		thread := block.Threads[0]
		require.Equal(t, "1234567890", thread.ID)
		require.Len(t, thread.Messages, 2)

		msg1 := thread.Messages[0]
		require.Equal(t, "16505551111", msg1.From)
		require.Equal(t, "ABGGFlA5Fpa", msg1.ID)
		require.NotEmpty(t, msg1.Time)
		require.Equal(t, "media_placeholder", msg1.Type)
		require.Equal(t, "read", msg1.Status)
		require.True(t, msg1.FromMe)
		require.Empty(t, msg1.Text)

		msg2 := thread.Messages[1]
		require.Equal(t, "16315551181", msg2.From)
		require.Equal(t, "text", msg2.Type)
		require.Equal(t, "this is a text message", msg2.Text)
		require.Equal(t, "delivered", msg2.Status)
		require.False(t, msg2.FromMe)

		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(historyMessage))))
	require.NoError(t, err)
}

func TestListener_OnContactSyncMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnContactSyncMessage(func(message *listener.ContactSyncMessage) error {
		require.NotEmpty(t, message)
		require.Equal(t, "123123123", message.ToPhoneNumberId)
		require.Len(t, message.Events, 1)

		event := message.Events[0]
		require.Equal(t, "contact", event.Type)
		require.Equal(t, "add", event.Action)
		require.Equal(t, int64(1773945449482), event.Time.UnixMilli())
		require.Equal(t, 0, event.Version)

		require.NotNil(t, event.Contact)
		require.Equal(t, "John Doe", event.Contact.FullName)
		require.Equal(t, "John", event.Contact.FirstName)
		require.Equal(t, "3213213214", event.Contact.PhoneNumber)

		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(contactSyncMessage))))
	require.NoError(t, err)
}

func TestListener_OnMessageEcho(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnMessageEcho(func(message *listener.MessageEcho) error {
		require.NotEmpty(t, message)
		require.Equal(t, "558594138387", message.From)
		require.Equal(t, "558599386799", message.To)
		require.Equal(t, "wamid.HBgMNTU4NTk0MTM4Mzg3FQIAERgUMkExOUE4REQ3NTU2RkQ0MzE3OEQA", message.ID)
		require.NotEmpty(t, message.Time)
		require.Equal(t, "audio", message.Type)
		require.Equal(t, "282416054963178", message.ToPhoneNumberId)

		require.NotNil(t, message.Audio)
		require.Equal(t, "26243872851911592", message.Audio.AudioID)
		require.Equal(t, "audio/ogg; codecs=opus", message.Audio.Mimetype)
		require.True(t, message.Audio.Voice)

		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(messageEchoMessage))))
	require.NoError(t, err)
}

func TestListener_OnAccountUpdateMessage(t *testing.T) {
	client := listener.NewMessageListener()

	client.OnAccountUpdateMessage(func(message *listener.AccountUpdateMessage) error {
		require.NotEmpty(t, message)
		require.Equal(t, "16505551111", message.PhoneNumber)
		require.Equal(t, listener.AccountUpdateVerified, message.Event)
		return nil
	})

	err := client.ReadBodySync(io.NopCloser(bytes.NewReader([]byte(accountUpdateMessage))))
	require.NoError(t, err)
}
