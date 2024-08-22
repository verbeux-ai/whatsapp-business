package whatsapp_business

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PhoneNumberResponse struct {
	VerifiedName           string                                   `json:"verified_name"`
	CodeVerificationStatus string                                   `json:"code_verification_status"`
	DisplayPhoneNumber     string                                   `json:"display_phone_number"`
	QualityRating          string                                   `json:"quality_rating"`
	PlatformType           string                                   `json:"platform_type"`
	Throughput             PhoneNumberResponseThroughput            `json:"throughput"`
	LastOnboardedTime      string                                   `json:"last_onboarded_time"`
	WebhookConfiguration   PhoneNumberResponseWhatsappConfiguration `json:"webhook_configuration"`
	Id                     string                                   `json:"id"`
	*ErrorResponse
}

type PhoneNumberResponseThroughput struct {
	Level string `json:"level"`
}

type PhoneNumberResponseWhatsappConfiguration struct {
	WhatsappBusinessAccount string `json:"whatsapp_business_account"`
	Application             string `json:"application"`
}

func (s *Client) GetPhoneNumber(phoneID string) (*PhoneNumberResponse, error) {
	res, err := s.metaRequestWithToken(nil, http.MethodGet, fmt.Sprintf("%s", phoneID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn PhoneNumberResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %v", toReturn.ErrorResponse.Error.Message, toReturn)
	}

	return &toReturn, nil
}

type RegisterPhoneNumberRequest struct {
	MessagingProduct messagingProductType `json:"messaging_product"`
	Pin              string               `json:"pin"`
}

type RegisterPhoneNumberResponse struct {
	Success bool `json:"success"`
	*ErrorResponse
}

func (s *Client) RegisterPhoneNumber(phoneID string, pin string) (*RegisterPhoneNumberResponse, error) {
	res, err := s.metaRequestWithToken(RegisterPhoneNumberRequest{
		MessagingProduct: whatsappMessagingProduct,
		Pin:              pin,
	}, http.MethodPost, fmt.Sprintf(phoneNumberRegister, phoneID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn RegisterPhoneNumberResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %v", toReturn.ErrorResponse.Error.Message, toReturn)
	}

	return &toReturn, nil
}

type SetPhoneNumberWebhookRequest struct {
	WebhookConfiguration SetPhoneNumberWebhookConfig `json:"webhook_configuration"`
}

type SetPhoneNumberWebhookConfig struct {
	OverrideCallbackUri string `json:"override_callback_uri"`
	VerifyToken         string `json:"verify_token"`
}

type SetPhoneNumberWebhookResponse struct {
	Success bool `json:"success"`
	*ErrorResponse
}

func (s *Client) SetPhoneNumberWebhook(phoneID string, request *SetPhoneNumberWebhookConfig) (*SetPhoneNumberWebhookResponse, error) {
	res, err := s.metaRequestWithToken(SetPhoneNumberWebhookRequest{
		WebhookConfiguration: *request,
	}, http.MethodPost, phoneID)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn SetPhoneNumberWebhookResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %v", toReturn.ErrorResponse.Error.Message, toReturn)
	}

	return &toReturn, nil
}
