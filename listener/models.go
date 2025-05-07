package listener

import (
	"time"
)

type TextMessage struct {
	Contacts        []Contact
	ID              string
	Message         string
	Time            time.Time
	ToPhoneNumberId string
}

type TextMessageListener func(message *TextMessage) error

type AudioMessage struct {
	Contacts        []Contact
	ID              string
	AudioID         string
	Mimetype        string
	Sha256          string
	Voice           bool
	Time            time.Time
	ToPhoneNumberId string
}

type AudioMessageListener func(message *AudioMessage) error

type ImageMessage struct {
	Contacts        []Contact
	ID              string
	ImageID         string
	Mimetype        string
	Sha256          string
	Caption         string
	Time            time.Time
	ToPhoneNumberId string
}
type ImageMessageListener func(message *ImageMessage) error

type DocumentMessage struct {
	Contacts        []Contact
	ID              string
	DocumentID      string
	Mimetype        string
	Sha256          string
	Filename        string `json:"filename"`
	Caption         string
	Time            time.Time
	ToPhoneNumberId string
}
type DocumentMessageListener func(message *DocumentMessage) error

type Origin struct {
	Type string
}

type StatusType string

const (
	StatusTypeRead      StatusType = "read"
	StatusTypeDelivered StatusType = "delivered"
	StatusTypeSent      StatusType = "sent"
)

type StatusMessage struct {
	ID             string
	Status         StatusType
	Time           time.Time
	WaID           string
	ConversationID string
	Origin         Origin
}
type StatusMessageListener func(message *StatusMessage) error
