package listener

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
)

func (s *listener) treatTextMessage(text rawMessageContent) (*TextMessage, error) {
	var content string
	if text.Text != nil {
		return nil, ErrEmptyMessage
	}

	messageTimeInt, err := strconv.ParseInt(text.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTimestampInvalid, err)
	}

	messageTime := time.Unix(messageTimeInt, 0)
	return &TextMessage{
		From:    text.From,
		ID:      text.ID,
		Message: content,
		Time:    messageTime,
	}, nil
}

func (s *listener) ReadBodyAsync(rawBody io.ReadCloser) error {
	var data RawMessage
	if err := json.NewDecoder(rawBody).Decode(&data); err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	for _, entry := range data.Entry {
		for _, change := range entry.Changes {
			for _, message := range change.Value.Messages {
				if message.Text != nil {
					if s.textMessageListener != nil {
						wg.Add(1)
						go func() {
							defer wg.Done()
							msg, err := s.treatTextMessage(message)
							if err != nil {
								s.chError <- err
								return
							}
							if err := (*s.textMessageListener)(msg); err != nil {
								s.chError <- err
							}
						}()
					}
				}
			}
		}
	}

	wg.Wait()

	return nil
}
