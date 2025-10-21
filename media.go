package whatsapp_business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"path"
	"strings"
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

type DownloadMediaResponse struct {
	*http.Response
	GuessFileName string `json:"guess_file_name"`
}

func (s *Client) DownloadMedia(ctx context.Context, mediaURL string) (*DownloadMediaResponse, error) {
	resp, err := s.metaDownloadFile(ctx, mediaURL)
	if err != nil {
		return nil, err
	}

	fileName := guessFilename(resp, mediaURL, "download")

	return &DownloadMediaResponse{
		resp,
		fileName,
	}, nil
}

func guessFilename(resp *http.Response, urlStr, fallbackBase string) string {
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if fn := params["filename"]; fn != "" {
				return fn
			}
			if fn := params["filename*"]; fn != "" { // RFC 5987
				return fn
			}
		}
	}
	if base := path.Base(urlStr); base != "" && !strings.HasSuffix(base, "/") {
		return base
	}

	ext := ""
	if ct := resp.Header.Get("Content-Type"); ct != "" {
		if exts, _ := mime.ExtensionsByType(ct); len(exts) > 0 {
			ext = exts[0]
		}
	}
	if ext == "" {
		ext = ".bin"
	}
	if fallbackBase == "" {
		fallbackBase = "download"
	}

	return fallbackBase + ext
}
