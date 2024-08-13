package sdk

import (
	"encoding/json"

	"github.com/verbeux-ai/whatsapp-business/domain/entities"
	"github.com/verbeux-ai/whatsapp-business/pkg/whatsapp"
)

type wpp struct {
	token string
}

// Auth verify if token is right
func (s *wpp) Auth(token string) error {
	if token != s.token {
		return ErrWrongToken
	}

	return nil
}

func (s *wpp) RawMessage(data json.RawMessage) (*entities.RawMessage, error) {
	var result entities.RawMessage
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *wpp) TextMessageParallel(data json.RawMessage, handler func(messages entities.Message) error) error {
	for message {
		go func() {
			handler(message)
		}()
	}

	wait()

	return nil
}

func NewWhatsappBusiness(token string) whatsapp.Business {
	return &wpp{
		token: token,
	}
}
