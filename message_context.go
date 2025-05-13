package whatsapp_business

type SendMessageOption func(*MessageContext)

func WithReplyMessage(messageID string) SendMessageOption {
	return func(context *MessageContext) {
		context.MessageId = messageID
	}
}
