package template

import (
	"slices"

	"github.com/verbeux-ai/whatsapp-business"
)

type MessageBuilder struct {
	TemplateMessage *whatsapp_business.TemplateMessage
	TemplateData    *whatsapp_business.TemplateData
	Opts            []whatsapp_business.TemplateOption
}

func NewTemplateMessageBuilder(message *whatsapp_business.TemplateMessage) *MessageBuilder {
	return &MessageBuilder{
		TemplateMessage: message,
	}
}

func (s *MessageBuilder) WithTemplate(template *whatsapp_business.TemplateData) *MessageBuilder {
	s.TemplateData = template
	return s
}

func (s *MessageBuilder) WithOpts(opts []whatsapp_business.TemplateOption) *MessageBuilder {
	newOpts := slices.Clone(opts)
	s.Opts = newOpts
	return s
}

func (s *MessageBuilder) buildOpts() *whatsapp_business.TemplateMessage {
	templateMessage := *s.TemplateMessage
	for _, opt := range s.Opts {
		opt(&templateMessage)
	}
	return &templateMessage
}

func (s *MessageBuilder) Build() (any, error) {
	templateType := s.GetType()
	templateMessage := s.buildOpts()
	switch templateType {
	case whatsapp_business.TextMessageType:
		return s.buildText(templateMessage)
	case whatsapp_business.ImageMessageType:
		return s.buildImage(templateMessage)
	}
	return nil, nil
}

func (s *MessageBuilder) GetType() whatsapp_business.MessageType {
	for _, component := range s.TemplateMessage.Components {
		if component.Type == whatsapp_business.TemplateMessageComponentTypeHeader {
			if len(component.Parameters) <= 0 {
				continue
			}
			for _, param := range component.Parameters {
				switch param.Type {
				case whatsapp_business.TemplateMessageComponentParameterImage:
					return whatsapp_business.ImageMessageType
				case whatsapp_business.TemplateMessageComponentParameterVideo:
					return whatsapp_business.VideoMessageType
				case whatsapp_business.TemplateMessageComponentParameterDocument:
					return whatsapp_business.DocumentMessageType
				}
			}
		}
	}

	return whatsapp_business.TextMessageType
}
