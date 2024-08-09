package whatsapp

import (
	"encoding/json"

	"github.com/verbeux-ai/whatsapp-business/domain/entities"
)

type Business interface {
	Auth(token string) error
	RawMessage(data json.RawMessage) (*entities.RawMessage, error)
	// Message (data json.RawMessage) (entities.Message, error)
	// in future we will do a better treatment of return

	// SendMessage
}
