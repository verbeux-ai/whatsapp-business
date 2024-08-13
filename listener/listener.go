package listener

type listener struct {
	textMessageListener *TextMessageListener
}

func NewMessageListener() MessageListener {
	return &listener{}
}

type MessageListener interface {
	OnTextMessage(TextMessageListener)
}

func (s *listener) OnTextMessage(message TextMessageListener) {
	s.textMessageListener = &message
}
