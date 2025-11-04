package whatsapp_business

import (
	"regexp"
	"strings"
	"time"

	"github.com/verbeux-ai/whatsapp-business/listener"
)

func substituteParams(text string, params map[string]string) string {
	if params == nil || text == "" {
		return text
	}
	re := regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)\}\} `)

	replacer := func(match string) string {
		key := strings.Trim(match, "{}")
		if val, ok := params[key]; ok {
			return val
		}
		return match
	}

	return re.ReplaceAllStringFunc(text, replacer)
}

func (s *TemplateMessage) buildSubstitutedText(componentType TemplateComponentType, opts ...TemplateOption) string {
	for _, opt := range opts {
		opt(s)
	}

	var originalText string
	for _, component := range s.Components {
		if component.Type == componentType {
			originalText = component.Text
			break
		}
	}

	paramsMap := make(map[string]string)
	for _, component := range s.Components {
		if component.Type == componentType {
			for _, param := range component.Parameters {
				if param.Type == TemplateComponentParameterText && param.Text != nil {
					paramsMap[param.ParameterName] = *param.Text
				}
			}
			break
		}
	}

	return substituteParams(originalText, paramsMap)
}

func (s *TemplateMessage) BuildTextMessage(opts []TemplateOption) *listener.TextMessage {
	headerText := s.buildSubstitutedText(TemplateComponentTypeHeader, opts...)
	bodyText := s.buildSubstitutedText(TemplateComponentTypeBody, opts...)

	finalMessage := bodyText
	if headerText != "" {
		finalMessage = headerText + "\n\n" + bodyText
	}

	return &listener.TextMessage{
		Message: finalMessage,
		Time:    time.Now(),
	}
}

func (s *TemplateMessage) BuildImageMessage(opts []TemplateOption) *listener.ImageMessage {
	finalCaption := s.buildSubstitutedText(TemplateComponentTypeBody, opts...)

	return &listener.ImageMessage{
		Caption: finalCaption,
		Time:    time.Now(),
	}
}

func (s *TemplateMessage) GetTemplateType(opts []TemplateOption) MessageType {
	for _, opt := range opts {
		opt(s)
	}
	for _, component := range s.Components {
		for _, param := range component.Parameters {
			switch param.Type {
			case TemplateComponentParameterText:
				return TextMessageType
			case TemplateComponentParameterImage:
				return ImageMessageType
			}
		}
	}
	return ""
}
