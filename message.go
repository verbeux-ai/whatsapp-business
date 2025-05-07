package whatsapp_business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type messageRecipientType string

const (
	individualRecipientType messageRecipientType = "individual"
)

type messagingProductType string

const (
	whatsappMessagingProduct messagingProductType = "whatsapp"
)

type messageStatus string

const (
	messageStatusRead messageStatus = "read"
)

type messageRequestType string

const (
	textMessageType        messageRequestType = "text"
	imageMessageType       messageRequestType = "image"
	audioMessageType       messageRequestType = "audio"
	videoMessageType       messageRequestType = "video"
	documentMessageType    messageRequestType = "document"
	interactiveMessageType messageRequestType = "interactive"
)

type baseMessageRequest struct {
	MessagingProduct messagingProductType `json:"messaging_product"`
	RecipientType    messageRecipientType `json:"recipient_type"`
	To               string               `json:"to"`
	Type             messageRequestType   `json:"type"`
	Context          *MessageContext      `json:"context,omitempty"`
}
type MessageContext struct {
	MessageId string `json:"message_id,omitempty"`
}

func newBaseMessageRequest(to string, t messageRequestType, options ...SendMessageOption) baseMessageRequest {
	result := baseMessageRequest{
		MessagingProduct: whatsappMessagingProduct,
		RecipientType:    individualRecipientType,
		To:               to,
		Type:             t,
	}

	if len(options) > 0 {
		result.Context = &MessageContext{}
	}

	for _, opt := range options {
		opt(result.Context)
	}

	return result
}

type textMessageRequest struct {
	baseMessageRequest
	Text TextMessage `json:"text"`
}

type TextMessage struct {
	PreviewUrl *bool  `json:"preview_url,omitempty"`
	Body       string `json:"body"`
}

type contactMessageResponse struct {
	Input string `json:"input"`
	WaId  string `json:"wa_id"`
}

type contentMessageResponse struct {
	Id string `json:"id"`
}

type MessageResponse struct {
	MessagingProduct string                   `json:"messaging_product"`
	Contacts         []contactMessageResponse `json:"contacts"`
	Messages         []contentMessageResponse `json:"messages"`
	*ErrorResponse
}

func (s *Client) SendTextMessage(ctx context.Context, to string, d TextMessage, options ...SendMessageOption) (*MessageResponse, error) {
	body := textMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, textMessageType, options...),
		Text:               d,
	}
	return s.messageRequest(ctx, body, http.MethodPost)
}

type ImageMessage struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
	ID      string `json:"id"`
}

type imageMessageRequest struct {
	baseMessageRequest
	Image ImageMessage `json:"image"`
}

func (s *Client) SendImageMessage(ctx context.Context, to string, d ImageMessage, options ...SendMessageOption) (*MessageResponse, error) {
	body := imageMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, imageMessageType, options...),
		Image:              d,
	}
	return s.messageRequest(ctx, body, http.MethodPost)
}

type VideoMessage struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
	ID      string `json:"id"`
}

type videoMessageRequest struct {
	baseMessageRequest
	Video VideoMessage `json:"video"`
}

func (s *Client) SendVideoMessage(ctx context.Context, to string, d VideoMessage, options ...SendMessageOption) (*MessageResponse, error) {
	body := videoMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, videoMessageType, options...),
		Video:              d,
	}
	return s.messageRequest(ctx, body, http.MethodPost)
}

type AudioMessage struct {
	Link string `json:"link"`
	ID   string `json:"id"`
}

type audioMessageRequest struct {
	baseMessageRequest
	Audio AudioMessage `json:"audio"`
}

func (s *Client) SendAudioMessage(ctx context.Context, to string, d AudioMessage, options ...SendMessageOption) (*MessageResponse, error) {
	body := audioMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, audioMessageType, options...),
		Audio:              d,
	}
	return s.messageRequest(ctx, body, http.MethodPost)
}

type DocumentMessage struct {
	Link     string `json:"link"`
	ID       string `json:"id"`
	Caption  string `json:"caption"`
	Filename string `json:"filename"`
}

type DocumentMessageRequest struct {
	baseMessageRequest
	Document DocumentMessage `json:"document"`
}

func (s *Client) SendDocumentMessage(ctx context.Context, to string, d DocumentMessage, options ...SendMessageOption) (*MessageResponse, error) {
	body := DocumentMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, documentMessageType, options...),
		Document:           d,
	}
	return s.messageRequest(ctx, body, http.MethodPost)
}

type InteractiveMessageRequest struct {
	baseMessageRequest
	Interactive InteractiveMessage `json:"interactive"`
}

func (s *Client) SendInteractiveMessage(ctx context.Context, to string, d InteractiveMessage, options ...SendMessageOption) (*MessageResponse, error) {
	body := InteractiveMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, interactiveMessageType, options...),
		Interactive:        d,
	}
	return s.messageRequest(ctx, body, http.MethodPost)
}

type InteractiveMessageType string

const (
	InteractiveMessageTypeList InteractiveMessageType = "LIST"
)

type InteractiveMessage struct {
	Type   InteractiveMessageType   `json:"type"`
	Header InteractiveMessageHeader `json:"header"`
	Body   InteractiveMessageBody   `json:"body"`
	Footer InteractiveMessageFooter `json:"footer"`
	Action InteractiveMessageAction `json:"action"`
}
type InteractiveMessageHeaderType string

const (
	InteractiveMessageHeaderTypeText InteractiveMessageHeaderType = "text"
)

type InteractiveMessageHeader struct {
	Type InteractiveMessageHeaderType `json:"type"`
	Text string                       `json:"text"`
}

type InteractiveMessageBody struct {
	Text string `json:"text"`
}

type InteractiveMessageFooter struct {
	Text string `json:"text"`
}

type InteractiveMessageAction struct {
	Button   string                      `json:"button"`
	Sections []InteractiveMessageSection `json:"sections"`
}

type InteractiveMessageSection struct {
	Title string                  `json:"title"`
	Rows  []InteractiveMessageRow `json:"rows"`
}

type InteractiveMessageRow struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type readMessageRequest struct {
	MessagingProduct messagingProductType `json:"messaging_product"`
	Status           messageStatus        `json:"status"`
	MessageId        string               `json:"message_id"`
}

func (s *Client) ReadMessage(ctx context.Context, messageID string) (*MessageResponse, error) {
	body := readMessageRequest{
		MessagingProduct: whatsappMessagingProduct,
		Status:           messageStatusRead,
		MessageId:        messageID,
	}

	return s.messageRequest(ctx, body, http.MethodPut)
}

func (s *Client) messageRequest(ctx context.Context, body any, method string) (*MessageResponse, error) {
	resp, err := s.metaRequestWithToken(ctx, body, method, fmt.Sprintf("%s/%s", s.phoneNumberID, messagesEndpoint))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn MessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return &toReturn, nil
}
