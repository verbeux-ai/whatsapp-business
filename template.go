package whatsapp_business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ListTemplateResponse struct {
	Data   []TemplateData `json:"data"`
	Paging Paging         `json:"paging"`
	*ErrorResponse
}

type TemplateData struct {
	Name            string              `json:"name"`
	ParameterFormat string              `json:"parameter_format"`
	Components      []TemplateComponent `json:"components"`
	Language        string              `json:"language"`
	Status          string              `json:"status"`
	Category        string              `json:"category"`
	SubCategory     string              `json:"sub_category,omitempty"`
	Id              string              `json:"id"`
}

type TemplateComponent struct {
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

func (s *Client) ListTemplates(ctx context.Context) ([]TemplateData, error) {
	resp, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, fmt.Sprintf("%s/message_templates", s.businessID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn ListTemplateResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return toReturn.Data, nil
}
