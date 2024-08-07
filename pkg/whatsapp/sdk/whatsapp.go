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

func (s *wpp) Message(data json.RawMessage) (*entities.RawMessage, error) {

}

func NewWhatsappBusiness(token string) whatsapp.Business {
	return &wpp{
		token: token,
	}
}
