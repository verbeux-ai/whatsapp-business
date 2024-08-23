package whatsapp_business

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	*ErrorResponse
}

func (s *Client) GenerateAccessToken(code, clientID, clientSecret string) (*AuthResponse, error) {
	res, err := s.metaRequest(nil, http.MethodGet, fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s", oAuthAccessTokenEndpoint, code, clientID, clientSecret))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var toReturn AuthResponse
	if err = json.NewDecoder(res.Body).Decode(&toReturn); err != nil {
		return nil, err
	}
	if toReturn.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %v", toReturn.ErrorResponse.Error.Message, toReturn)
	}

	return &toReturn, nil
}
