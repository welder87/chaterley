package core

import (
	"errors"
	"fmt"
)

var (
	ErrNameEmpty       = errors.New("name is empty")
	ErrStartsWithDigit = errors.New("cannot start with a digit")
	ErrNameUnchanged   = errors.New("name must be different from current")
	ErrMemberNotFound  = errors.New("member not found")
	ErrMemberCount  = errors.New("member count")
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

type PermissionError struct {
	Field  string
	Reason string
	Err    error
}

func (e PermissionError) Error() string {
	return fmt.Sprintf("permission failed on %s: %s", e.Field, e.Reason)
}

func (e PermissionError) Unwrap() error {
	return e.Err
}
