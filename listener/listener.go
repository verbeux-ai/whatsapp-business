package listener

type listener struct {
	chError chan error

	textMessageListener *TextMessageListener
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
}

func (s *listener) OnTextMessage(message TextMessageListener) {
	s.textMessageListener = &message
}
