package listener

import (
	"io"
	"sync"
)

type listener struct {
	chError chan error

	textMessageListener          *TextMessageListener
	buttonMessageListener        *ButtonMessageListener
	audioMessageListener         *AudioMessageListener
	imageMessageListener         *ImageMessageListener
	documentMessageListener      *DocumentMessageListener
	statusMessageListener        *StatusMessageListener
	historyMessageListener       *HistoryMessageListener
	contactSyncMessageListener   *ContactSyncMessageListener
	accountUpdateMessageListener *AccountUpdateMessageListener
	messageEchoListener          *MessageEchoListener
}

func NewMessageListener() MessageListener {
	return &listener{
		chError: make(chan error),
	}
}

func (s *listener) HandleErrors(f func(error)) (closer func()) {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-s.chError:
				f(err)
			case <-done:
				return
			}

		}
	}()

	return func() {
		done <- struct{}{}
	}
}

type MessageListener interface {
	HandleErrors(f func(error)) (closer func())
	OnTextMessage(TextMessageListener)
	OnButtonMessage(ButtonMessageListener)
	OnAudioMessage(AudioMessageListener)
	OnImageMessage(ImageMessageListener)
	OnDocumentMessage(DocumentMessageListener)
	OnStatusMessage(StatusMessageListener)
	OnHistoryMessage(HistoryMessageListener)
	OnContactSyncMessage(ContactSyncMessageListener)
	OnAccountUpdateMessage(AccountUpdateMessageListener)
	OnMessageEcho(MessageEchoListener)
	ReadBodyAsync(io.ReadCloser) *sync.WaitGroup
	ReadBodySync(io.ReadCloser) error
}

func (s *listener) OnTextMessage(listener TextMessageListener) {
	s.textMessageListener = &listener
}

func (s *listener) OnButtonMessage(listener ButtonMessageListener) {
	s.buttonMessageListener = &listener
}

func (s *listener) OnAudioMessage(listener AudioMessageListener) {
	s.audioMessageListener = &listener
}

func (s *listener) OnImageMessage(listener ImageMessageListener) {
	s.imageMessageListener = &listener
}

func (s *listener) OnDocumentMessage(listener DocumentMessageListener) {
	s.documentMessageListener = &listener
}

func (s *listener) OnStatusMessage(listener StatusMessageListener) {
	s.statusMessageListener = &listener
}

func (s *listener) OnHistoryMessage(listener HistoryMessageListener) {
	s.historyMessageListener = &listener
}

func (s *listener) OnContactSyncMessage(listener ContactSyncMessageListener) {
	s.contactSyncMessageListener = &listener
}

func (s *listener) OnAccountUpdateMessage(listener AccountUpdateMessageListener) {
	s.accountUpdateMessageListener = &listener
}

func (s *listener) OnMessageEcho(listener MessageEchoListener) {
	s.messageEchoListener = &listener
}
