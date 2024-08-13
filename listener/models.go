package listener

import (
	"time"
)

type TextMessage struct {
	From    string
	ID      string
	Message string
	Time    time.Time
}

type TextMessageListener func(message *TextMessage) error
