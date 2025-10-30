package whatsapp_business

type SendMessageOption func(*MessageContext)

func WithReplyMessage(messageID string) SendMessageOption {
	return func(context *MessageContext) {
		context.MessageId = messageID
	}
}

type TemplateOption func(*TemplateMessage)

func WithHeaderTextNamed(name, text string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateComponentTypeHeader, TemplateComponentParameter{
			Type:          TemplateComponentParameterText,
			Text:          &text,
			ParameterName: name,
		})
	}
}

func WithHeaderImageLink(link string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateComponentTypeHeader, TemplateComponentParameter{
			Type:  TemplateComponentParameterImage,
			Image: &TemplateComponentParameterTypeImage{Link: link},
		})
	}
}

func WithHeaderDocumentLink(link string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateComponentTypeHeader, TemplateComponentParameter{
			Type:     TemplateComponentParameterDocument,
			Document: &TemplateComponentParameterTypeDocument{Link: link},
		})
	}
}

func WithHeaderVideoLink(name, link string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateComponentTypeHeader, TemplateComponentParameter{
			Type:          TemplateComponentParameterVideo,
			Video:         &TemplateComponentParameterTypeVideo{Link: link},
			ParameterName: name,
		})
	}
}

func WithBodyTextNamed(name, text string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateComponentTypeBody, TemplateComponentParameter{
			Type:          TemplateComponentParameterText,
			Text:          &text,
			ParameterName: name,
		})
	}
}

func pushComponents(components []TemplateComponents, templateType TemplateComponentType, toInsert ...TemplateComponentParameter) []TemplateComponents {
	foundIndex := -1
	for i, compt := range components {
		if compt.Type == templateType {
			foundIndex = i
			break
		}
	}

	var actualParams []TemplateComponentParameter
	if foundIndex != -1 {
		actualParams = components[foundIndex].Parameters
	}

	actualParams = append(actualParams, toInsert...)
	newComponent := TemplateComponents{
		Type:       templateType,
		Parameters: actualParams,
	}

	if foundIndex == -1 {
		components = append(components, newComponent)
	} else {
		components[foundIndex] = newComponent
	}

	return components
}
