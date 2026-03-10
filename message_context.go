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
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateMessageComponentTypeHeader, TemplateMessageComponentParameter{
			Type:          TemplateMessageComponentParameterText,
			ParamText:     &text,
			ParameterName: name,
		})
	}
}

func WithHeaderImageLink(link string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateMessageComponentTypeHeader, TemplateMessageComponentParameter{
			Type:  TemplateMessageComponentParameterImage,
			Image: &TemplateMessageComponentParameterTypeImage{Link: link},
		})
	}
}

func WithHeaderDocumentLink(link string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateMessageComponentTypeHeader, TemplateMessageComponentParameter{
			Type:     TemplateMessageComponentParameterDocument,
			Document: &TemplateMessageComponentParameterTypeDocument{Link: link},
		})
	}
}

func WithHeaderVideoLink(name, link string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateMessageComponentTypeHeader, TemplateMessageComponentParameter{
			Type:          TemplateMessageComponentParameterVideo,
			Video:         &TemplateMessageComponentParameterTypeVideo{Link: link},
			ParameterName: name,
		})
	}
}

func WithBodyTextNamed(name, text string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateMessageComponentTypeBody, TemplateMessageComponentParameter{
			Type:          TemplateMessageComponentParameterText,
			ParamText:     &text,
			ParameterName: name,
		})
	}
}

func WithBodyText(text string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = pushComponents(templateMessage.Components, TemplateMessageComponentTypeBody, TemplateMessageComponentParameter{
			Type:      TemplateMessageComponentParameterText,
			ParamText: &text,
		})
	}
}

func WithURLButton(index, text string) TemplateOption {
	return func(templateMessage *TemplateMessage) {
		templateMessage.Components = append(templateMessage.Components, TemplateMessageComponent{
			Type:    TemplateMessageComponentTypeSingleButton,
			SubType: TemplateMessageComponentButtonSubTypeURL,
			Index:   index,
			Parameters: []TemplateMessageComponentParameter{
				{
					Type:      TemplateMessageComponentParameterText,
					ParamText: &text,
				},
			},
		})
	}
}

func pushComponents(components []TemplateMessageComponent, templateType TemplateMessageComponentType, toInsert ...TemplateMessageComponentParameter) []TemplateMessageComponent {
	foundIndex := -1
	for i, compt := range components {
		if compt.Type == templateType {
			foundIndex = i
			break
		}
	}

	var actualParams []TemplateMessageComponentParameter
	if foundIndex != -1 {
		actualParams = components[foundIndex].Parameters
	}

	actualParams = append(actualParams, toInsert...)
	newComponent := TemplateMessageComponent{
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
