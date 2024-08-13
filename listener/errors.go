package listener

import (
	"errors"
)

var (
	ErrEmptyMessage     = errors.New("empty message from meta")
	ErrTimestampInvalid = errors.New("invalid timestamp from meta")
)
