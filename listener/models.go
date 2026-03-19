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
	Referral        *RawReferral
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
	StatusTypeError     StatusType = "error"
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

type ButtonMessage struct {
	Contacts        []Contact
	ID              string
	Message         string
	Payload         string
	Time            time.Time
	ToPhoneNumberId string
}

type ButtonMessageListener func(message *ButtonMessage) error

type HistoryMessage struct {
	ToPhoneNumberId string
	History         []HistoryBlock
}

type HistoryBlock struct {
	Phase      int
	ChunkOrder int
	Progress   int
	Threads    []Thread
}

type Thread struct {
	ID       string
	Messages []ThreadMessage
}

type ThreadMessage struct {
	From   string
	ID     string
	Time   time.Time
	Type   string
	Status string
	FromMe bool

	// Message content by type
	Text     string
	Audio    *AudioMessage
	Image    *ImageMessage
	Document *DocumentMessage
	Video    *VideoMessage
	Sticker  *StickerMessage
	Location *LocationMessage
	Reaction *ReactionMessage
	Contacts []SharedContact
	Errors   []MessageError
	Context  *MessageContext
	Edit     *MessageEdit
}

type SharedContact struct {
	FirstName     string
	FormattedName string
	Phones        []SharedContactPhone
}

type SharedContactPhone struct {
	Phone string
	WaID  string
	Type  string
}

type MessageError struct {
	Code    int
	Title   string
	Message string
	Details string
}

type MessageContext struct {
	From      string
	ID        string
	Forwarded bool
}

type MessageEdit struct {
	OriginalMessageID string
	Type              string
	Text              string
}

type StickerMessage struct {
	ID       string
	Mimetype string
	Sha256   string
}

type LocationMessage struct {
	Latitude  float64
	Longitude float64
	Name      string
	Address   string
}

type ReactionMessage struct {
	MessageID string
	Emoji     string
}

type HistoryMessageListener func(message *HistoryMessage) error

type ContactSyncMessage struct {
	ToPhoneNumberId string
	Events          []ContactSyncEvent
}

type ContactSyncEvent struct {
	Type    string
	Action  string
	Contact *ContactSyncDetail
	Time    time.Time
	Version int
}

type ContactSyncDetail struct {
	FullName    string
	FirstName   string
	PhoneNumber string
}

type ContactSyncMessageListener func(message *ContactSyncMessage) error

type AccountUpdateEvent string

const (
	AccountUpdateVerified AccountUpdateEvent = "VERIFIED_ACCOUNT"
)

type AccountUpdateMessage struct {
	PhoneNumber string
	Event       AccountUpdateEvent
}

type AccountUpdateMessageListener func(message *AccountUpdateMessage) error

type MessageEcho struct {
	ID              string
	From            string
	To              string
	Time            time.Time
	Type            string
	ToPhoneNumberId string

	Text     string
	Audio    *AudioMessage
	Image    *ImageMessage
	Document *DocumentMessage
	Video    *VideoMessage
	Sticker  *StickerMessage
}

type VideoMessage struct {
	ID       string
	Mimetype string
	Sha256   string
}

type MessageEchoListener func(message *MessageEcho) error
