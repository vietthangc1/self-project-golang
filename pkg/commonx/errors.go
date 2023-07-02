package commonx

import (
	"errors"
	"fmt"
)

var (
	ErrItemNotFound   error = errors.New("item not found")
	ErrUnknown        error = errors.New("unknown")
	ErrKeyNotFound    error = errors.New("key not found")
	ErrWrongMethod    error = errors.New("wrong method")
	ErrNotFoundParams error = errors.New("params not found")
)

func ErrorMessages(err error, msg string) error {
	return fmt.Errorf("%s, err: %w", msg, err)
}
