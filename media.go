package whatsapp_business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Media struct {
	MessagingProduct string `json:"messaging_product"`
	URL              string `json:"url"`
	MimeType         string `json:"mime_type"`
	Sha256           string `json:"sha256"`
	FileSize         uint64 `json:"file_size"`
	ID               string `json:"id"`
	*ErrorResponse
}

func (s *Client) GetMedia(ctx context.Context, mediaID string) (*Media, error) {
	resp, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, fmt.Sprintf("%s?phone_number_id=%s", mediaID, s.phoneNumberID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn Media
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return &toReturn, nil
}
