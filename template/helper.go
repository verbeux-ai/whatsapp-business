package template

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	whatsapp_business "github.com/verbeux-ai/whatsapp-business"
	"github.com/verbeux-ai/whatsapp-business/listener"
)

func replaceParams(text string, params map[string]string) string {
	if params == nil || text == "" {
		return text
	}
	re := regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)\}\}`)

	replacer := func(match string) string {
		key := strings.Trim(match, "{}")
		key = strings.TrimSpace(key)
		if val, ok := params[key]; ok {
			return val
		}
		return match
	}

	return re.ReplaceAllStringFunc(text, replacer)
}

func (s *MessageBuilder) buildReplacedText(componentType whatsapp_business.TemplateMessageComponentType, templateMessage *whatsapp_business.TemplateMessage) (string, error) {
	var originalText string
	hasImage := false
	for _, component := range s.TemplateData.Components {
		if component.Format == "IMAGE" {
			hasImage = true
			continue
		}
		templateDataComponentType := string(componentType)
		if component.Type == templateDataComponentType {
			originalText = component.Text
			break
		}
	}
	if len(originalText) <= 0 && !hasImage {
		err := fmt.Errorf("Component not found")
		return "", err
	}

	paramsMap := make(map[string]string)
	for _, component := range templateMessage.Components {
		if component.Type == componentType {
			for _, param := range component.Parameters {
				if param.Type == whatsapp_business.TemplateMessageComponentParameterText && param.ParamText != nil {
					paramsMap[param.ParameterName] = *param.ParamText
				}
			}
			break
		}
	}

	return replaceParams(originalText, paramsMap), nil
}

func (s *MessageBuilder) buildText(templateMessage *whatsapp_business.TemplateMessage) (*listener.TextMessage, error) {
	hasHeader := false
	for _, component := range s.TemplateData.Components {
		if component.Type == "HEADER" {
			hasHeader = true
		}
	}
	var headerText string
	if hasHeader {
		text, err := s.buildReplacedText(whatsapp_business.TemplateMessageComponentTypeHeader, templateMessage)
		if err != nil {
			return nil, err
		}
		headerText = text
	}
	bodyText, err := s.buildReplacedText(whatsapp_business.TemplateMessageComponentTypeBody, templateMessage)
	if err != nil {
		return nil, err
	}
	finalMessage := bodyText
	if len(headerText) > 0 {
		finalMessage = headerText + "\n\n" + bodyText
	}

	return &listener.TextMessage{
		Message: finalMessage,
		Time:    time.Now(),
	}, nil
}

func (s *MessageBuilder) buildImage(templateMessage *whatsapp_business.TemplateMessage) (*listener.ImageMessage, error) {
	finalCaption, err := s.buildReplacedText(whatsapp_business.TemplateMessageComponentTypeBody, templateMessage)
	if err != nil {
		return nil, err
	}

	return &listener.ImageMessage{
		Caption: finalCaption,
		Time:    time.Now(),
	}, nil
}
