package whatsapp_business

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BusinessResponse struct {
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	TimezoneId               string `json:"timezone_id"`
	MessageTemplateNamespace string `json:"message_template_namespace"`
	*ErrorResponse
}

func (s *Client) GetBusiness(businessAccountId string) (*BusinessResponse, error) {
	res, err := s.metaRequestWithToken(nil, http.MethodGet, fmt.Sprintf("%s", businessAccountId))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn BusinessResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %v", toReturn.ErrorResponse.Error.Message, toReturn)
	}

	return &toReturn, nil
}

func (s *Client) SetBusinessWebhook(businessAccountId string, request *SetWebhookConfig) (*SetBusinessWebhookResponse, error) {
	res, err := s.metaRequestWithToken(SetBusinessWebhookRequest{
		WebhookConfiguration: *request,
	}, http.MethodPost, fmt.Sprintf(businessSubscribedApps, businessAccountId))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn SetBusinessWebhookResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %v", toReturn.ErrorResponse.Error.Message, toReturn)
	}

	return &toReturn, nil
}
