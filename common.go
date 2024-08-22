package whatsapp_business

import (
	"bytes"
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

func (s *Client) metaRequestWithToken(reqBody any, method, endpoint string) (*http.Response, error) {
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
	req.Header.Set("Content-Type", "application/json")

	return s.httpClient.Do(req)
}

func (s *Client) metaRequest(reqBody any, method, endpoint string) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		marshalledBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(marshalledBody)
	}

	url := fmt.Sprintf("%s/%s", s.baseUrl, endpoint)

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return s.httpClient.Do(req)
}
