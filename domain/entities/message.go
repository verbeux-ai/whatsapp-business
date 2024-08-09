package entities

type RawMessage struct {
	Object string     `json:"object"`
	Entry  []RawEntry `json:"entry"`
}

type RawEntry struct {
	Id      string      `json:"id"`
	Changes []RawChange `json:"changes"`
}

type RawChange struct {
	Value RawValue `json:"value"`
	Field string   `json:"field"`
}

type RawValue struct {
	MessagingProduct string              `json:"messaging_product"`
	Metadata         RawMetadata         `json:"metadata"`
	Contacts         []RawContact        `json:"contacts"`
	Messages         []RawMessageContent `json:"messages"`
}

type RawMessageContent struct {
	From      string `json:"from"`
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
	// Text represents a text message, it should be nil if empty
	Text *RawText `json:"text"`
	Type string   `json:"type"`
}

type RawText struct {
	Body string `json:"body"`
}

type RawContact struct {
	Profile RawProfile `json:"profile"`
	WaId    string     `json:"wa_id"`
}

type RawProfile struct {
	Name string `json:"name"`
}

type RawMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}
