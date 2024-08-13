package handlers

import (
	"encoding/json"
)

type HandledMessage struct {
	Object string         `json:"object"`
	Entry  []handledEntry `json:"entry"`
}

type handledEntry struct {
	ID      string          `json:"id"`
	Changes []handledChange `json:"changes"`
}

type handledChange struct {
	Value handledValue `json:"value"`
	Field string       `json:"field"`
}

type handledValue struct {
	MessagingProduct string           `json:"messaging_product"`
	Metadata         handledMetadata  `json:"metadata"`
	Contacts         []handledContact `json:"contacts"`
	Messages         []handledMessage `json:"messages"`
	Statuses         []handledStatus  `json:"statuses,omitempty"`
}

type handledMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type handledContact struct {
	Profile handledProfile `json:"profile"`
	WaID    string         `json:"wa_id"`
}

type handledProfile struct {
	Name string `json:"name"`
}

type handledMessage struct {
	From        string                   `json:"from"`
	ID          string                   `json:"id"`
	Timestamp   string                   `json:"timestamp"`
	Type        string                   `json:"type"`
	Text        *HandledText             `json:"text,omitempty"`
	Image       *handledImage            `json:"image,omitempty"`
	Sticker     *handledSticker          `json:"sticker,omitempty"`
	Location    *handledLocation         `json:"location,omitempty"`
	Contacts    *[]handledMessageContact `json:"contacts,omitempty"`
	Reaction    *handledReaction         `json:"reaction,omitempty"`
	Interactive *handledInteractive      `json:"interactive,omitempty"`
	Referral    *handledReferral         `json:"referral,omitempty"`
	Order       *handledOrder            `json:"order,omitempty"`
	System      *handledSystem           `json:"system,omitempty"`
	Errors      *[]handledError          `json:"errors,omitempty"`
	Context     *handledContext          `json:"context,omitempty"`
}

type HandledText struct {
	Body string `json:"body"`
}

type handledImage struct {
	Caption  string `json:"caption"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type handledSticker struct {
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type handledLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

type handledMessageContact struct {
	Addresses []handledAddress `json:"addresses,omitempty"`
	Birthday  string           `json:"birthday,omitempty"`
	Emails    []handledEmail   `json:"emails,omitempty"`
	Name      handledName      `json:"name,omitempty"`
	Org       handledOrg       `json:"org,omitempty"`
	Phones    []handledPhone   `json:"phones,omitempty"`
	Urls      []handledUrl     `json:"urls,omitempty"`
}

type handledAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Type        string `json:"type"`
	Zip         string `json:"zip"`
}

type handledEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type handledName struct {
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	Suffix        string `json:"suffix"`
	Prefix        string `json:"prefix"`
}

type handledOrg struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
}

type handledPhone struct {
	Phone string `json:"phone"`
	WaID  string `json:"wa_id"`
	Type  string `json:"type"`
}

type handledUrl struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type handledReaction struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

type handledInteractive struct {
	Type        string              `json:"type"`
	ListReply   *handledListReply   `json:"list_reply,omitempty"`
	ButtonReply *handledButtonReply `json:"button_reply,omitempty"`
}

type handledListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type handledButtonReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type handledReferral struct {
	SourceURL    string `json:"source_url"`
	SourceID     string `json:"source_id"`
	SourceType   string `json:"source_type"`
	Headline     string `json:"headline"`
	Body         string `json:"body"`
	MediaType    string `json:"media_type"`
	ImageURL     string `json:"image_url,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	CtwaClid     string `json:"ctwa_clid"`
}

type handledOrder struct {
	CatalogID    string               `json:"catalog_id"`
	ProductItems []handledProductItem `json:"product_items"`
	Text         string               `json:"text"`
}

type handledProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}

type handledSystem struct {
	Body    string `json:"body"`
	NewWaID string `json:"new_wa_id"`
	Type    string `json:"type"`
}

type handledError struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
	Title   string `json:"title"`
}

type handledContext struct {
	From            string                  `json:"from"`
	ID              string                  `json:"id"`
	ReferredProduct *handledReferredProduct `json:"referred_product,omitempty"`
}

type handledReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

type handledStatus struct {
	ID           string               `json:"id"`
	Status       string               `json:"status"`
	Timestamp    string               `json:"timestamp"`
	RecipientID  string               `json:"recipient_id"`
	Conversation *handledConversation `json:"conversation,omitempty"`
	Pricing      *handledPricing      `json:"pricing,omitempty"`
	Errors       *[]handledError      `json:"errors,omitempty"`
}

type handledConversation struct {
	ID                  string        `json:"id"`
	ExpirationTimestamp string        `json:"expiration_timestamp,omitempty"`
	Origin              handledOrigin `json:"origin"`
}

type handledOrigin struct {
	Type string `json:"type"`
}

type handledPricing struct {
	Billable     bool   `json:"billable"`
	PricingModel string `json:"pricing_model"`
	Category     string `json:"category"`
}

type TextMessageHandler func(message *HandledText) error

type listener struct {
	textMessageHandler *TextMessageHandler
}

func NewMessageHandler() MessageHandler {
	return &listener{}
}

type MessageHandler interface {
	OnTextMessage(TextMessageHandler)
}

func (s *listener) OnTextMessage(message TextMessageHandler) {
	s.textMessageHandler = &message
}

func (s *listener) ProcessJson(d []byte) error {
	var data HandledMessage
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	for _, entry := range data.Entry {
		for _, change := range entry.Changes {
			for _, message := range change.Value.Messages {

				if message.Text != nil {
					if s.textMessageHandler != nil {
						return (*s.textMessageHandler)(message.Text)
					}
				}

			}
		}
	}

	return nil
}
