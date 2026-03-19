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
				switch change.Field {
				case "messages":
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

				case "history":
					if s.historyMessageListener != nil {
						wg.Add(1)
						go func(change RawChange) {
							defer wg.Done()
							if err := s.treatHistory(change); err != nil {
								s.chError <- err
							}
						}(change)
					}

				case "smb_app_state_sync":
					if s.contactSyncMessageListener != nil {
						wg.Add(1)
						go func(change RawChange) {
							defer wg.Done()
							if err := s.treatContactSync(change); err != nil {
								s.chError <- err
							}
						}(change)
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
			switch change.Field {
			case "messages":
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

			case "history":
				if s.historyMessageListener != nil {
					if err := s.treatHistory(change); err != nil {
						return err
					}
				}

				if s.messageEchoListener != nil {
					for _, echo := range change.Value.MessageEchoes {
						if err := s.treatMessageEcho(echo, change.Value.Metadata); err != nil {
							return err
						}
					}
				}

			case "smb_message_echoes":
				if s.messageEchoListener != nil {
					for _, echo := range change.Value.MessageEchoes {
						if err := s.treatMessageEcho(echo, change.Value.Metadata); err != nil {
							return err
						}
					}
				}

			case "smb_app_state_sync":
				if s.contactSyncMessageListener != nil {
					if err := s.treatContactSync(change); err != nil {
						return err
					}
				}

			case "account_update":
				if s.accountUpdateMessageListener != nil {
					if err := s.treatAccountUpdate(change); err != nil {
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

func (s *listener) treatHistory(change RawChange) error {
	msg, err := s.treatHistoryMessage(change)
	if err != nil {
		return err
	}
	if err := (*s.historyMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatHistoryMessage(change RawChange) (*HistoryMessage, error) {
	var blocks []HistoryBlock
	for _, raw := range change.Value.History {
		var threads []Thread
		for _, rawThread := range raw.Threads {
			var messages []ThreadMessage
			for _, rawMsg := range rawThread.Messages {
				threadMsg, err := rawToThreadMessage(rawMsg)
				if err != nil {
					return nil, err
				}
				messages = append(messages, threadMsg)
			}

			threads = append(threads, Thread{
				ID:       rawThread.ID,
				Messages: messages,
			})
		}

		blocks = append(blocks, HistoryBlock{
			Phase:      raw.Metadata.Phase,
			ChunkOrder: raw.Metadata.ChunkOrder,
			Progress:   raw.Metadata.Progress,
			Threads:    threads,
		})
	}

	return &HistoryMessage{
		ToPhoneNumberId: change.Value.Metadata.PhoneNumberID,
		History:         blocks,
	}, nil
}

func rawToThreadMessage(rawMsg RawMessageContent) (ThreadMessage, error) {
	messageTimeInt, err := strconv.ParseInt(rawMsg.Timestamp, 10, 64)
	if err != nil {
		return ThreadMessage{}, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}
	msgTime := time.Unix(messageTimeInt, 0)

	var status string
	var fromMe bool
	if rawMsg.HistoryContext != nil {
		status = rawMsg.HistoryContext.Status
		fromMe = rawMsg.HistoryContext.FromMe
	}

	threadMsg := ThreadMessage{
		From:   rawMsg.From,
		ID:     rawMsg.ID,
		Time:   msgTime,
		Type:   rawMsg.Type,
		Status: status,
		FromMe: fromMe,
	}

	if rawMsg.Text != nil {
		threadMsg.Text = rawMsg.Text.Body
	}
	if rawMsg.Audio != nil {
		threadMsg.Audio = &AudioMessage{
			ID:       rawMsg.ID,
			AudioID:  rawMsg.Audio.ID,
			Mimetype: rawMsg.Audio.MimeType,
			Sha256:   rawMsg.Audio.Sha256,
			Voice:    rawMsg.Audio.Voice,
			Time:     msgTime,
		}
	}
	if rawMsg.Image != nil {
		threadMsg.Image = &ImageMessage{
			ID:       rawMsg.ID,
			ImageID:  rawMsg.Image.ID,
			Mimetype: rawMsg.Image.MimeType,
			Sha256:   rawMsg.Image.Sha256,
			Caption:  rawMsg.Image.Caption,
			Time:     msgTime,
		}
	}
	if rawMsg.Document != nil {
		threadMsg.Document = &DocumentMessage{
			ID:         rawMsg.ID,
			DocumentID: rawMsg.Document.ID,
			Mimetype:   rawMsg.Document.MimeType,
			Sha256:     rawMsg.Document.Sha256,
			Filename:   rawMsg.Document.Filename,
			Caption:    rawMsg.Document.Caption,
			Time:       msgTime,
		}
	}
	if rawMsg.Video != nil {
		threadMsg.Video = &VideoMessage{
			ID:       rawMsg.Video.ID,
			Mimetype: rawMsg.Video.MimeType,
			Sha256:   rawMsg.Video.Sha256,
		}
	}
	if rawMsg.Sticker != nil {
		threadMsg.Sticker = &StickerMessage{
			ID:       rawMsg.Sticker.ID,
			Mimetype: rawMsg.Sticker.MimeType,
			Sha256:   rawMsg.Sticker.Sha256,
		}
	}
	if rawMsg.Location != nil {
		threadMsg.Location = &LocationMessage{
			Latitude:  rawMsg.Location.Latitude,
			Longitude: rawMsg.Location.Longitude,
			Name:      rawMsg.Location.Name,
			Address:   rawMsg.Location.Address,
		}
	}
	if rawMsg.Reaction != nil {
		threadMsg.Reaction = &ReactionMessage{
			MessageID: rawMsg.Reaction.MessageID,
			Emoji:     rawMsg.Reaction.Emoji,
		}
	}
	if rawMsg.Contacts != nil {
		for _, c := range *rawMsg.Contacts {
			sc := SharedContact{
				FirstName:     c.Name.FirstName,
				FormattedName: c.Name.FormattedName,
			}
			for _, p := range c.Phones {
				sc.Phones = append(sc.Phones, SharedContactPhone{
					Phone: p.Phone,
					WaID:  p.WaID,
					Type:  p.Type,
				})
			}
			threadMsg.Contacts = append(threadMsg.Contacts, sc)
		}
	}
	if rawMsg.Errors != nil {
		for _, e := range *rawMsg.Errors {
			threadMsg.Errors = append(threadMsg.Errors, MessageError{
				Code:    e.Code,
				Title:   e.Title,
				Message: e.Message,
				Details: e.ErrorData.Details,
			})
		}
	}
	if rawMsg.Context != nil {
		threadMsg.Context = &MessageContext{
			From:      rawMsg.Context.From,
			ID:        rawMsg.Context.ID,
			Forwarded: rawMsg.Context.Forwarded,
		}
	}
	if rawMsg.Edit != nil {
		edit := &MessageEdit{
			OriginalMessageID: rawMsg.Edit.OriginalMessageID,
			Type:              rawMsg.Edit.Message.Type,
		}
		if rawMsg.Edit.Message.Text != nil {
			edit.Text = rawMsg.Edit.Message.Text.Body
		}
		threadMsg.Edit = edit
	}

	return threadMsg, nil
}

func (s *listener) treatContactSync(change RawChange) error {
	msg, err := s.treatContactSyncMessage(change)
	if err != nil {
		return err
	}
	if err := (*s.contactSyncMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatContactSyncMessage(change RawChange) (*ContactSyncMessage, error) {
	var events []ContactSyncEvent
	for _, raw := range change.Value.StateSync {
		messageTimeInt, err := strconv.ParseInt(raw.Metadata.Timestamp, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
		}
		eventTime := parseTimestamp(messageTimeInt)

		event := ContactSyncEvent{
			Type:    raw.Type,
			Action:  raw.Action,
			Time:    eventTime,
			Version: raw.Metadata.Version,
		}

		if raw.Contact != nil {
			event.Contact = &ContactSyncDetail{
				FullName:    raw.Contact.FullName,
				FirstName:   raw.Contact.FirstName,
				PhoneNumber: raw.Contact.PhoneNumber,
			}
		}

		events = append(events, event)
	}

	return &ContactSyncMessage{
		ToPhoneNumberId: change.Value.Metadata.PhoneNumberID,
		Events:          events,
	}, nil
}

func (s *listener) treatAccountUpdate(change RawChange) error {
	msg := &AccountUpdateMessage{
		PhoneNumber: change.Value.PhoneNumber,
		Event:       AccountUpdateEvent(change.Value.Event),
	}
	if err := (*s.accountUpdateMessageListener)(msg); err != nil {
		return err
	}
	return nil
}

func (s *listener) treatMessageEcho(raw RawMessageContent, metadata RawMetadata) error {
	messageTimeInt, err := strconv.ParseInt(raw.Timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}
	msgTime := time.Unix(messageTimeInt, 0)

	msg := &MessageEcho{
		ID:              raw.ID,
		From:            raw.From,
		To:              raw.To,
		Time:            msgTime,
		Type:            raw.Type,
		ToPhoneNumberId: metadata.PhoneNumberID,
	}

	if raw.Text != nil {
		msg.Text = raw.Text.Body
	}
	if raw.Audio != nil {
		msg.Audio = &AudioMessage{
			ID:       raw.ID,
			AudioID:  raw.Audio.ID,
			Mimetype: raw.Audio.MimeType,
			Sha256:   raw.Audio.Sha256,
			Voice:    raw.Audio.Voice,
			Time:     msgTime,
		}
	}
	if raw.Image != nil {
		msg.Image = &ImageMessage{
			ID:       raw.ID,
			ImageID:  raw.Image.ID,
			Mimetype: raw.Image.MimeType,
			Sha256:   raw.Image.Sha256,
			Caption:  raw.Image.Caption,
			Time:     msgTime,
		}
	}
	if raw.Document != nil {
		msg.Document = &DocumentMessage{
			ID:         raw.ID,
			DocumentID: raw.Document.ID,
			Mimetype:   raw.Document.MimeType,
			Sha256:     raw.Document.Sha256,
			Filename:   raw.Document.Filename,
			Caption:    raw.Document.Caption,
			Time:       msgTime,
		}
	}
	if raw.Video != nil {
		msg.Video = &VideoMessage{
			ID:       raw.Video.ID,
			Mimetype: raw.Video.MimeType,
			Sha256:   raw.Video.Sha256,
		}
	}
	if raw.Sticker != nil {
		msg.Sticker = &StickerMessage{
			ID:       raw.Sticker.ID,
			Mimetype: raw.Sticker.MimeType,
			Sha256:   raw.Sticker.Sha256,
		}
	}

	if err := (*s.messageEchoListener)(msg); err != nil {
		return err
	}
	return nil
}

func parseTimestamp(ts int64) time.Time {
	if ts > 9999999999 {
		return time.UnixMilli(ts)
	}
	return time.Unix(ts, 0)
}
