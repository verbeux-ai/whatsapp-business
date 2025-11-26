package template_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	whatsapp "github.com/verbeux-ai/whatsapp-business"
	"github.com/verbeux-ai/whatsapp-business/listener"
	"github.com/verbeux-ai/whatsapp-business/template"
)

func TestNewTemplateMessageBuilder(t *testing.T) {
	message := &whatsapp.TemplateMessage{}
	builder := template.NewTemplateMessageBuilder(message)

	assert.NotNil(t, builder)
	assert.Equal(t, message, builder.TemplateMessage)
}

func TestMessageBuilderWithTemplate(t *testing.T) {
	builder := template.NewTemplateMessageBuilder(&whatsapp.TemplateMessage{})
	templateData := &whatsapp.TemplateData{Name: "test_template"}
	builder.WithTemplate(templateData)

	assert.Equal(t, templateData, builder.TemplateData)
}

func TestMessageBuilderWithOpts(t *testing.T) {
	builder := template.NewTemplateMessageBuilder(&whatsapp.TemplateMessage{})
	opts := []whatsapp.TemplateOption{
		func(tm *whatsapp.TemplateMessage) {
			tm.Name = "test"
		},
	}
	builder.WithOpts(opts)

	assert.Equal(t, len(opts), len(builder.Opts))
}

func TestMessageBuilderGetType(t *testing.T) {
	t.Run("should return TextMessageType when no header", func(t *testing.T) {
		message := &whatsapp.TemplateMessage{
			Components: []whatsapp.TemplateMessageComponent{
				{
					Type: whatsapp.TemplateMessageComponentTypeBody,
				},
			},
		}
		builder := template.NewTemplateMessageBuilder(message)
		messageType := builder.GetType()
		assert.Equal(t, whatsapp.TextMessageType, messageType)
	})

	t.Run("should return TextMessageType when header is text", func(t *testing.T) {
		message := &whatsapp.TemplateMessage{
			Components: []whatsapp.TemplateMessageComponent{
				{
					Type: whatsapp.TemplateMessageComponentTypeHeader,
					Parameters: []whatsapp.TemplateMessageComponentParameter{
						{
							Type: whatsapp.TemplateMessageComponentParameterText,
						},
					},
				},
			},
		}
		builder := template.NewTemplateMessageBuilder(message)
		messageType := builder.GetType()
		assert.Equal(t, whatsapp.TextMessageType, messageType)
	})

	t.Run("should return ImageMessageType when header is image", func(t *testing.T) {
		message := &whatsapp.TemplateMessage{
			Components: []whatsapp.TemplateMessageComponent{
				{
					Type: whatsapp.TemplateMessageComponentTypeHeader,
					Parameters: []whatsapp.TemplateMessageComponentParameter{
						{
							Type: whatsapp.TemplateMessageComponentParameterImage,
						},
					},
				},
			},
		}
		builder := template.NewTemplateMessageBuilder(message)
		messageType := builder.GetType()
		assert.Equal(t, whatsapp.ImageMessageType, messageType)
	})

	t.Run("should return VideoMessageType when header is video", func(t *testing.T) {
		message := &whatsapp.TemplateMessage{
			Components: []whatsapp.TemplateMessageComponent{
				{
					Type: whatsapp.TemplateMessageComponentTypeHeader,
					Parameters: []whatsapp.TemplateMessageComponentParameter{
						{
							Type: whatsapp.TemplateMessageComponentParameterVideo,
						},
					},
				},
			},
		}
		builder := template.NewTemplateMessageBuilder(message)
		messageType := builder.GetType()
		assert.Equal(t, whatsapp.VideoMessageType, messageType)
	})

	t.Run("should return DocumentMessageType when header is document", func(t *testing.T) {
		message := &whatsapp.TemplateMessage{
			Components: []whatsapp.TemplateMessageComponent{
				{
					Type: whatsapp.TemplateMessageComponentTypeHeader,
					Parameters: []whatsapp.TemplateMessageComponentParameter{
						{
							Type: whatsapp.TemplateMessageComponentParameterDocument,
						},
					},
				},
			},
		}
		builder := template.NewTemplateMessageBuilder(message)
		messageType := builder.GetType()
		assert.Equal(t, whatsapp.DocumentMessageType, messageType)
	})
}

func TestMessageBuilderBuildText(t *testing.T) {
	paramText := "John"
	message := &whatsapp.TemplateMessage{
		Components: []whatsapp.TemplateMessageComponent{
			{
				Type:          whatsapp.TemplateMessageComponentTypeBody,
				ComponentText: "Hello {{name}}",
				Parameters: []whatsapp.TemplateMessageComponentParameter{
					{
						Type:          whatsapp.TemplateMessageComponentParameterText,
						ParameterName: "name",
						ParamText:     &paramText,
					},
				},
			},
		},
	}
	builder := template.NewTemplateMessageBuilder(message).WithTemplate(&whatsapp.TemplateData{Components: []whatsapp.TemplateDataComponent{{Type: "BODY", Text: "Hello {{name}}"}}})
	result, err := builder.Build()

	assert.NoError(t, err)
	assert.IsType(t, &listener.TextMessage{}, result)
	textMessage := result.(*listener.TextMessage)
	assert.Equal(t, "Hello John", textMessage.Message)
	assert.WithinDuration(t, time.Now(), textMessage.Time, time.Second)
}

func TestMessageBuilderBuildImage(t *testing.T) {
	paramText := "customer"
	message := &whatsapp.TemplateMessage{
		Components: []whatsapp.TemplateMessageComponent{
			{
				Type: whatsapp.TemplateMessageComponentTypeHeader,
				Parameters: []whatsapp.TemplateMessageComponentParameter{
					{
						Type: whatsapp.TemplateMessageComponentParameterImage,
						Image: &whatsapp.TemplateMessageComponentParameterTypeImage{
							Link: "http://example.com/image.png",
						},
					},
				},
			},
			{
				Type:          whatsapp.TemplateMessageComponentTypeBody,
				ComponentText: "Hello {{name}}",
				Parameters: []whatsapp.TemplateMessageComponentParameter{
					{
						Type:          whatsapp.TemplateMessageComponentParameterText,
						ParameterName: "name",
						ParamText:     &paramText,
					},
				},
			},
		},
	}
	builder := template.NewTemplateMessageBuilder(message).WithTemplate(&whatsapp.TemplateData{Components: []whatsapp.TemplateDataComponent{{Type: "BODY", Text: "Hello {{name}}"}}})
	result, err := builder.Build()

	assert.NoError(t, err)
	assert.IsType(t, &listener.ImageMessage{}, result)
	imageMessage := result.(*listener.ImageMessage)
	assert.Equal(t, "Hello customer", imageMessage.Caption)
	assert.WithinDuration(t, time.Now(), imageMessage.Time, time.Second)
}

func TestMessageBuilderBuildError(t *testing.T) {
	message := &whatsapp.TemplateMessage{
		Components: []whatsapp.TemplateMessageComponent{
			{
				Type: whatsapp.TemplateMessageComponentTypeBody,
			},
		},
	}
	builder := template.NewTemplateMessageBuilder(message).WithTemplate(&whatsapp.TemplateData{Components: []whatsapp.TemplateDataComponent{}})
	_, err := builder.Build()
	assert.Error(t, err)
	assert.EqualError(t, err, "Component not found")
}
