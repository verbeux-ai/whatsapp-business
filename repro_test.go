package whatsapp_business

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReproduction_Marshal(t *testing.T) {
	jsonPayload := `{
    "name": "template_test_3",
    "category": "MARKETING",
    "language": "pt_BR",
    "components": [
        {
            "type": "HEADER",
            "format": "TEXT",
            "text": "Welcome to Verbeux"
        },
        {
            "type": "BODY",
            "text": "Hello {{1}}, welcome to our platform!",
            "example": {
                "body_text": [
                    [
                        "John Doe"
                    ]
                ]
            }
        },
        {
            "type": "FOOTER",
            "text": "Sent by Verbeux AI"
        },
        {
            "type": "BUTTONS",
            "buttons": [
                {
                    "type": "QUICK_REPLY",
                    "text": "Learn More"
                },
                {
                   "type": "QUICK_REPLY",
                    "text": "Stop"
                }
            ]
        }
    ]
}`

	var req CreateTemplateRequest
	err := json.Unmarshal([]byte(jsonPayload), &req)
	assert.NoError(t, err)

	marshalled, err := json.Marshal(req)
	assert.NoError(t, err)

	fmt.Printf("Marshalled: %s\n", string(marshalled))
}
