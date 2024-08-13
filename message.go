package whatsapp_business

import (
	"bytes"
	"encoding/json"
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

type messageRequestType string

const (
	textMessageType messageRequestType = "text"
)

type baseMessageRequest struct {
	MessagingProduct messagingProductType `json:"messaging_product"`
	RecipientType    messageRecipientType `json:"recipient_type"`
	To               string               `json:"to"`
	Type             messageRequestType   `json:"type"`
}

func newBaseMessageRequest(to string, t messageRequestType) baseMessageRequest {
	return baseMessageRequest{
		MessagingProduct: whatsappMessagingProduct,
		RecipientType:    individualRecipientType,
		To:               to,
		Type:             t,
	}
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
}

func (s *Client) SendMessage(to string, txt string) error {
	_, err := s.SendTextMessage(to, TextMessage{Body: txt})
	return err
}

func (s *Client) SendTextMessage(to string, d TextMessage) (*MessageResponse, error) {
	body := textMessageRequest{
		baseMessageRequest: newBaseMessageRequest(to, textMessageType),
		Text:               d,
	}
	return s.messageRequest(body)
}

func (s *Client) messageRequest(body any) (*MessageResponse, error) {
	resp, err := s.metaRequest(body, http.MethodPost, fmt.Sprintf("%s/%s", s.phoneNumberID, "messages"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn MessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	return &toReturn, nil
}

func (s *Client) metaRequest(reqBody any, method, endpoint string) (*http.Response, error) {
	marshalledBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(marshalledBody)

	url := fmt.Sprintf("%s/%s", s.baseUrl, endpoint)

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.token)

	return s.httpClient.Do(req)
}
