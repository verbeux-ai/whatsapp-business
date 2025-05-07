package listener

import (
	"io"
)

type listener struct {
	chError chan error

	textMessageListener     *TextMessageListener
	audioMessageListener    *AudioMessageListener
	imageMessageListener    *ImageMessageListener
	documentMessageListener *DocumentMessageListener
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
	OnAudioMessage(AudioMessageListener)
	OnImageMessage(ImageMessageListener)
	OnDocumentMessage(DocumentMessageListener)
	ReadBodyAsync(rawBody io.ReadCloser) error
	ReadBodySync(rawBody io.ReadCloser) error
}

func (s *listener) OnTextMessage(listener TextMessageListener) {
	s.textMessageListener = &listener
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
