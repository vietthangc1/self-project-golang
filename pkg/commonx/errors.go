package commonx

import (
	"errors"
	"fmt"
)

var (
	ErrItemNotFound        error = errors.New("item not found")
	ErrUnknown             error = errors.New("unknown")
	ErrKeyNotFound         error = errors.New("key not found")
	ErrWrongMethod         error = errors.New("wrong method")
	ErrNotFoundParams      error = errors.New("params not found")
	ErrInsufficientDataGet error = errors.New("insufficient data get")
	ErrUnauthorized        error = errors.New("unaurhorized")
	ErrNotAuthenticated    error = errors.New("not authenticated")
)

func ErrorMessages(err error, msg string) error {
	return fmt.Errorf("%s, err: %w", msg, err)
}
