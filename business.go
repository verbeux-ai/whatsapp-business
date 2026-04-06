package whatsapp_business

import (
	"context"
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

func (s *Client) GetBusiness(ctx context.Context, businessAccountId string) (*BusinessResponse, error) {
	res, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, fmt.Sprintf("%s", businessAccountId))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn BusinessResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %+v", toReturn.ErrorResponse.Error.Message, toReturn.ErrorResponse.Error)
	}

	return &toReturn, nil
}

func (s *Client) SetBusinessWebhook(ctx context.Context, businessAccountId string, request *SetWebhookConfig) (*SetBusinessWebhookResponse, error) {
	res, err := s.metaRequestWithToken(ctx, SetBusinessWebhookRequest{
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
		return nil, fmt.Errorf("%s: %+v", toReturn.ErrorResponse.Error.Message, toReturn.ErrorResponse.Error)
	}

	return &toReturn, nil
}

type BusinessProfileResponse struct {
	Data []BusinessProfile `json:"data"`
	*ErrorResponse
}

type BusinessProfile struct {
	About             string   `json:"about"`
	Address           string   `json:"address"`
	Description       string   `json:"description"`
	Email             string   `json:"email"`
	MessagingProduct  string   `json:"messaging_product"`
	ProfilePictureUrl string   `json:"profile_picture_url"`
	Websites          []string `json:"websites"`
	Vertical          string   `json:"vertical"`
}

func (s *Client) GetBusinessProfile(ctx context.Context, phoneID string) (*BusinessProfile, error) {
	res, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, fmt.Sprintf(whatsappBusinessProfile, phoneID)+"?fields=about,address,description,email,profile_picture_url,websites,vertical")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn BusinessProfileResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %+v", toReturn.ErrorResponse.Error.Message, toReturn.ErrorResponse.Error)
	}

	if len(toReturn.Data) > 0 {
		return &toReturn.Data[0], nil
	}

	return nil, nil
}
