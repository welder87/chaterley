package core

import (
	"errors"
	"fmt"
)

var (
	ErrNameEmpty        = errors.New("name is empty")
	ErrStartsWithDigit  = errors.New("cannot start with a digit")
	ErrNameUnchanged    = errors.New("name must be different from current")
	ErrContentEmpty     = errors.New("content is empty")
	ErrContentTypeEmpty = errors.New("content type is empty")
)

type ValidationError struct {
	Field  string
	Reason string
	Err    error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Reason)
}

func (e ValidationError) Unwrap() error {
	return e.Err
}
