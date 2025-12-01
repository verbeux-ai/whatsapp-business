package whatsapp_business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ListTemplateResponse struct {
	Data   []TemplateData `json:"data"`
	Paging Paging         `json:"paging"`
	*ErrorResponse
}

type GetTemplateResponse struct {
	Data   *TemplateData `json:"data"`
	Paging Paging        `json:"paging"`
	*ErrorResponse
}

type TemplateData struct {
	Name            string                  `json:"name"`
	ParameterFormat string                  `json:"parameter_format"`
	Components      []TemplateDataComponent `json:"components"`
	Language        string                  `json:"language"`
	Status          string                  `json:"status"`
	Category        string                  `json:"category"`
	SubCategory     string                  `json:"sub_category,omitempty"`
	Id              string                  `json:"id"`
}

type TemplateDataComponent struct {
	Type    string           `json:"type"`
	Format  string           `json:"format,omitempty"`
	Text    string           `json:"text,omitempty"`
	Buttons []TemplateButton `json:"buttons,omitempty"`
	Example *TemplateExample `json:"example,omitempty"`
}

type TemplateButton struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type TemplateExample struct {
	HeaderTextNamedParams []NamedParam `json:"header_text_named_params,omitempty"`
	BodyTextNamedParams   []NamedParam `json:"body_text_named_params,omitempty"`
	HeaderHandle          []string     `json:"header_handle,omitempty"`
}

type NamedParam struct {
	ParamName string `json:"param_name"`
	Example   string `json:"example"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type ListTemplateFilter struct {
	Limit uint64 `json:"limit"`
	After string `json:"after"`
}

func (s *Client) ListTemplates(ctx context.Context, filter ListTemplateFilter) (*ListTemplateResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s/message_templates", s.businessID))
	if err != nil {
		return nil, err
	}

	q := u.Query()

	if filter.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", filter.Limit))
	}

	if filter.After != "" {
		q.Set("after", filter.After)
	}

	u.RawQuery = q.Encode()

	resp, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn *ListTemplateResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return toReturn, nil
}

func (s *Client) GetTemplate(ctx context.Context, name string, language string) (*TemplateData, error) {
	queryString := fmt.Sprintf("message_templates?name=%s&language=%s&limit=1", name, language)
	resp, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, fmt.Sprintf("%s/%s", s.businessID, queryString))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var toReturn ListTemplateResponse
	if err = json.Unmarshal(bodyBytes, &toReturn); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	if len(toReturn.Data) == 0 {
		return nil, errors.New("template not found with specified name and language")
	}

	return &toReturn.Data[0], nil
}
