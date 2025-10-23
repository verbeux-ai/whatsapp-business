package whatsapp_business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceId string `json:"fbtrace_id"`
	} `json:"error"`
}

func (s *Client) metaRequestWithToken(ctx context.Context, reqBody any, method, route string) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		marshalledBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(marshalledBody)
	}

	url := fmt.Sprintf("%s/%s", s.baseUrl, route)

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Content-Type", "application/json")

	return s.httpClient.Do(req)
}

func (s *Client) metaMultipartRequestWithToken(ctx context.Context, body io.Reader, contentType, method, route string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.baseUrl, route)

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Content-Type", contentType)

	return s.httpClient.Do(req)
}

func (s *Client) metaDownloadFile(ctx context.Context, mediaURL string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mediaURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.token)

	return s.httpClient.Do(req)
}

func (s *Client) metaRequest(ctx context.Context, reqBody any, method, route string) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		marshalledBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(marshalledBody)
	}

	url := fmt.Sprintf("%s/%s", s.baseUrl, route)

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return s.httpClient.Do(req)
}
