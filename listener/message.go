package listener

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
)

func (s *listener) ReadBodyAsync(RawBody io.ReadCloser) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		var data RawMessage
		if err := json.NewDecoder(RawBody).Decode(&data); err != nil {
			s.chError <- err
			return
		}

		for _, entry := range data.Entry {
			for _, change := range entry.Changes {
				var contacts []Contact
				for _, contact := range change.Value.Contacts {
					contacts = append(contacts, Contact{
						Name: contact.Profile.Name,
						WaID: contact.WaID,
					})
				}

				for _, message := range change.Value.Messages {
					if message.Text != nil && s.textMessageListener != nil {
						wg.Add(1)
						go func(message RawMessageContent, change RawChange, contacts []Contact) {
							defer wg.Done()
							if err := s.treatText(message, change, contacts); err != nil {
								s.chError <- err
							}
						}(message, change, contacts)
					}
					if message.Button != nil && s.buttonMessageListener != nil {
						wg.Add(1)
						go func(message RawMessageContent, change RawChange, contacts []Contact) {
							defer wg.Done()
							if err := s.treatButton(message, change, contacts); err != nil {
								s.chError <- err
							}
						}(message, change, contacts)
					}
					if message.Interactive != nil && s.buttonMessageListener != nil {
						wg.Add(1)
						go func(message RawMessageContent, change RawChange, contacts []Contact) {
							defer wg.Done()
							if err := s.treatList(message, change, contacts); err != nil {
								s.chError <- err
							}
						}(message, change, contacts)
					}
					if message.Audio != nil && s.audioMessageListener != nil {
						wg.Add(1)
						go func(message RawMessageContent, change RawChange, contacts []Contact) {
							defer wg.Done()
							if err := s.treatAudio(message, change, contacts); err != nil {
								s.chError <- err
							}
						}(message, change, contacts)
					}
					if message.Image != nil && s.imageMessageListener != nil {
						wg.Add(1)
						go func(message RawMessageContent, change RawChange, contacts []Contact) {
							defer wg.Done()
							if err := s.treatImage(message, change, contacts); err != nil {
								s.chError <- err
							}
						}(message, change, contacts)
					}
					if message.Document != nil && s.documentMessageListener != nil {
						wg.Add(1)
						go func(message RawMessageContent, change RawChange, contacts []Contact) {
							defer wg.Done()
							if err := s.treatDocument(message, change, contacts); err != nil {
								s.chError <- err
							}
						}(message, change, contacts)
					}
				}

				if s.statusMessageListener != nil {
					for _, status := range change.Value.Statuses {
						wg.Add(1)
						go func(status RawStatus) {
							defer wg.Done()
							if err := s.treatStatus(status); err != nil {
								s.chError <- err
							}
						}(status)
					}
				}
			}
		}
	}()

	return &wg
}

func (s *listener) ReadBodySync(RawBody io.ReadCloser) error {
	var data RawMessage
	if err := json.NewDecoder(RawBody).Decode(&data); err != nil {
		return err
	}

	for _, entry := range data.Entry {
		for _, change := range entry.Changes {
			var contacts []Contact
			for _, contact := range change.Value.Contacts {
				contacts = append(contacts, Contact{
					Name: contact.Profile.Name,
					WaID: contact.WaID,
				})
			}

			for _, message := range change.Value.Messages {
				if message.Button != nil && s.buttonMessageListener != nil {
					if err := s.treatButton(message, change, contacts); err != nil {
						return err
					}
				}
				if message.Interactive != nil && s.buttonMessageListener != nil {
					if err := s.treatList(message, change, contacts); err != nil {
						return err
					}
				}
				if message.Text != nil && s.textMessageListener != nil {
					if err := s.treatText(message, change, contacts); err != nil {
						return err
					}
				}
				if message.Audio != nil && s.audioMessageListener != nil {
					if err := s.treatAudio(message, change, contacts); err != nil {
						return err
					}
				}
				if message.Image != nil && s.imageMessageListener != nil {
					if err := s.treatImage(message, change, contacts); err != nil {
						return err
					}
				}
				if message.Document != nil && s.documentMessageListener != nil {
					if err := s.treatDocument(message, change, contacts); err != nil {
						return err
					}
				}
			}

			if s.statusMessageListener != nil {
				for _, status := range change.Value.Statuses {
					if err := s.treatStatus(status); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (s *listener) treatButton(message RawMessageContent, change RawChange, contacts []Contact) error {
	msg, err := s.treatButtonMessage(message, change.Value.Metadata)
	if err != nil {
		return err
	}
	msg.Contacts = contacts
	if err := (*s.buttonMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatButtonMessage(data RawMessageContent, metaData RawMetadata) (*ButtonMessage, error) {
	if data.Button == nil {
		return nil, ErrEmptyMessage
	}

	content := data.Button.Text
	payload := data.Button.Payload

	messageTimeInt, err := strconv.ParseInt(data.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}
	messageTime := time.Unix(messageTimeInt, 0)

	return &ButtonMessage{
		ID:              data.ID,
		Message:         content,
		Payload:         payload,
		Time:            messageTime,
		ToPhoneNumberId: metaData.PhoneNumberID,
	}, nil
}

func (s *listener) treatList(message RawMessageContent, change RawChange, contacts []Contact) error {
	msg, err := s.treatListMessage(message.Interactive, change.Value.Metadata, message.Timestamp, message.ID)
	if err != nil {
		return err
	}
	msg.Contacts = contacts
	if err := (*s.buttonMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatListMessage(data *RawInteractive, metaData RawMetadata, timestamp, id string) (*ButtonMessage, error) {
	if data.ListReply == nil {
		return nil, ErrEmptyMessage
	}

	content := data.ListReply.Title + "\n" + data.ListReply.Description

	messageTimeInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}
	messageTime := time.Unix(messageTimeInt, 0)

	return &ButtonMessage{
		ID:              id,
		Message:         content,
		Payload:         data.ListReply.ID,
		Time:            messageTime,
		ToPhoneNumberId: metaData.PhoneNumberID,
	}, nil
}

func (s *listener) treatText(message RawMessageContent, change RawChange, contacts []Contact) error {
	msg, err := s.treatTextMessage(message, change.Value.Metadata)
	if err != nil {
		return err
	}
	msg.Contacts = contacts
	if err := (*s.textMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatTextMessage(text RawMessageContent, metaData RawMetadata) (*TextMessage, error) {
	var content string
	if text.Text == nil {
		return nil, ErrEmptyMessage
	} else {
		content = text.Text.Body
	}

	messageTimeInt, err := strconv.ParseInt(text.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}

	messageTime := time.Unix(messageTimeInt, 0)
	return &TextMessage{
		ID:              text.ID,
		Message:         content,
		Time:            messageTime,
		ToPhoneNumberId: metaData.PhoneNumberID,
		Referral:        text.Referral,
	}, nil
}

func (s *listener) treatAudio(message RawMessageContent, change RawChange, contacts []Contact) error {
	msg, err := s.treatAudioMessage(message, change.Value.Metadata)
	if err != nil {
		return err
	}
	msg.Contacts = contacts
	if err := (*s.audioMessageListener)(msg); err != nil {
		return err
	}

	return nil
}

func (s *listener) treatAudioMessage(rawMessage RawMessageContent, metaData RawMetadata) (*AudioMessage, error) {
	var audioID string
	if rawMessage.Audio == nil {
		return nil, ErrEmptyMessage
	} else {
		audioID = rawMessage.Audio.ID
	}

	messageTimeInt, err := strconv.ParseInt(rawMessage.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}

	messageTime := time.Unix(messageTimeInt, 0)
	return &AudioMessage{
		Contacts:        nil,
		ID:              rawMessage.ID,
		AudioID:         audioID,
		Mimetype:        rawMessage.Audio.MimeType,
		Sha256:          rawMessage.Audio.Sha256,
		Voice:           rawMessage.Audio.Voice,
		Time:            messageTime,
		ToPhoneNumberId: metaData.PhoneNumberID,
	}, nil
}

func (s *listener) treatImage(message RawMessageContent, change RawChange, contacts []Contact) error {
	msg, err := s.treatImageMessage(message, change.Value.Metadata)
	if err != nil {
		return err
	}
	msg.Contacts = contacts
	if err := (*s.imageMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatImageMessage(rawMessage RawMessageContent, metaData RawMetadata) (*ImageMessage, error) {
	var imageID string
	if rawMessage.Image == nil {
		return nil, ErrEmptyMessage
	} else {
		imageID = rawMessage.Image.ID
	}

	messageTimeInt, err := strconv.ParseInt(rawMessage.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}

	messageTime := time.Unix(messageTimeInt, 0)
	return &ImageMessage{
		Contacts:        nil,
		ID:              rawMessage.ID,
		ImageID:         imageID,
		Mimetype:        rawMessage.Image.MimeType,
		Sha256:          rawMessage.Image.Sha256,
		Time:            messageTime,
		ToPhoneNumberId: metaData.PhoneNumberID,
		Caption:         rawMessage.Image.Caption,
	}, nil
}

func (s *listener) treatStatus(rawStatus RawStatus) error {
	msg, err := s.treatStatusMessage(rawStatus)
	if err != nil {
		return err
	}
	if err := (*s.statusMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatStatusMessage(rawStatus RawStatus) (*StatusMessage, error) {
	messageTimeInt, err := strconv.ParseInt(rawStatus.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}

	messageTime := time.Unix(messageTimeInt, 0)
	return &StatusMessage{
		ID:             rawStatus.ID,
		Status:         StatusType(rawStatus.Status),
		Time:           messageTime,
		WaID:           rawStatus.RecipientID,
		ConversationID: rawStatus.Conversation.ID,
		Origin: Origin{
			Type: rawStatus.Conversation.Origin.Type,
		},
	}, nil
}

func (s *listener) treatDocument(message RawMessageContent, change RawChange, contacts []Contact) error {
	msg, err := s.treatDocumentMessage(message, change.Value.Metadata)
	if err != nil {
		return err
	}
	msg.Contacts = contacts
	if err := (*s.documentMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatDocumentMessage(rawMessage RawMessageContent, metaData RawMetadata) (*DocumentMessage, error) {
	var documentID string
	if rawMessage.Document == nil {
		return nil, ErrEmptyMessage
	} else {
		documentID = rawMessage.Document.ID
	}

	messageTimeInt, err := strconv.ParseInt(rawMessage.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}

	messageTime := time.Unix(messageTimeInt, 0)
	return &DocumentMessage{
		Contacts:        nil,
		ID:              rawMessage.ID,
		DocumentID:      documentID,
		Mimetype:        rawMessage.Document.MimeType,
		Sha256:          rawMessage.Document.Sha256,
		Filename:        rawMessage.Document.Filename,
		Time:            messageTime,
		ToPhoneNumberId: metaData.PhoneNumberID,
		Caption:         rawMessage.Document.Caption,
	}, nil
}
