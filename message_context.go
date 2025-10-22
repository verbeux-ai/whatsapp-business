package whatsapp_business

type SendMessageOption func(*MessageContext)

func WithReplyMessage(messageID string) SendMessageOption {
	return func(context *MessageContext) {
		context.MessageId = messageID
	}
}

type TemplateOption func(*TemplateMessage)

func NewTextParameter(text string) TemplateComponentParameter {
	return TemplateComponentParameter{
		Type: TemplateComponentParameterText,
		Text: &text,
	}
}

func NewCurrencyParameter(fallback, code string, amount int) TemplateComponentParameter {
	return TemplateComponentParameter{
		Type: TemplateComponentParameterCurrency,
		Currency: &TemplateComponentParameterTypeCurrency{
			FallbackValue: fallback,
			Code:          code,
			Amount1000:    amount,
		},
	}
}

func NewDateTimeParameter(fallback string) TemplateComponentParameter {
	return TemplateComponentParameter{
		Type: TemplateComponentParameterDateTime,
		DateTime: &TemplateComponentParameterTypeDateTime{
			FallbackValue: fallback,
		},
	}
}

func WithHeaderText(text string) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type: TemplateComponentTypeHeader,
			Parameters: []TemplateComponentParameter{
				{
					Type: TemplateComponentParameterText,
					Text: &text,
				},
			},
		})
	}
}

func WithHeaderImage(link string) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type: TemplateComponentTypeHeader,
			Parameters: []TemplateComponentParameter{
				{
					Type:  TemplateComponentParameterImage,
					Image: &TemplateComponentParameterTypeImage{Link: link},
				},
			},
		})
	}
}

func WithHeaderVideo(link string) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type: TemplateComponentTypeHeader,
			Parameters: []TemplateComponentParameter{
				{
					Type:  TemplateComponentParameterVideo,
					Video: &TemplateComponentParameterTypeVideo{Link: link},
				},
			},
		})
	}
}

func WithHeaderDocument(link, filename string) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type: TemplateComponentTypeHeader,
			Parameters: []TemplateComponentParameter{
				{
					Type:     TemplateComponentParameterDocument,
					Document: &TemplateComponentParameterTypeDocument{Link: link, Filename: filename},
				},
			},
		})
	}
}

func WithBody(params ...TemplateComponentParameter) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type:       TemplateComponentTypeBody,
			Parameters: params,
		})
	}
}

func WithButtonURL(index, urlText string) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type:    TemplateComponentTypeButton,
			SubType: TemplateComponentButtonSubTypeURL,
			Index:   index,
			Parameters: []TemplateComponentParameter{
				{
					Type: TemplateComponentParameterText,
					Text: &urlText,
				},
			},
		})
	}
}

func WithButtonQuickReply(index, payload string) TemplateOption {
	return func(tmpl *TemplateMessage) {
		tmpl.Components = append(tmpl.Components, TemplateComponents{
			Type:    TemplateComponentTypeButton,
			SubType: TemplateComponentButtonSubTypeQuickReply,
			Index:   index,
			Parameters: []TemplateComponentParameter{
				{
					Type:    TemplateComponentParameterPayload,
					Payload: &payload,
				},
			},
		})
	}
}
