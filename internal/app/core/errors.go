package core

import (
	"errors"
	"fmt"
)

type ErrCode int

const (
	EmptyName ErrCode = iota
	NameUnchanged
	StartsWithDigit
	MemberNotFound
	MaxMemberCount
	MinMemberCount
	MemberIsExists
	MemberIsNotExists
	PasswordLength
	GenPasswordSalt
	GenPasswordHash
	Unknown
	InvalidPassword
	DecodePasswordFromB64
)

func (k ErrCode) String() string {
	return errorsByCode[k].Error()
}

var (
	ErrNameEmpty             = errors.New("name is empty")
	ErrNameUnchanged         = errors.New("name must be different from current")
	ErrStartsWithDigit       = errors.New("cannot start with a digit")
	ErrMemberNotFound        = errors.New("member not found")
	ErrMaxMemberCount        = errors.New("max member count")
	ErrMinMemberCount        = errors.New("min member count")
	ErrMemberIsExists        = errors.New("member is exists")
	ErrMemberIsNotExists     = errors.New("member is not exists")
	ErrPasswordLength        = errors.New("password length less than minimum simbols count")
	ErrGenPasswordSalt       = errors.New("generate salt is not possible")
	ErrGenPasswordHashed     = errors.New("generate hash is not possible")
	ErrInvalidPassword       = errors.New("input password is not equal current password")
	ErrDecodePasswordFromB64 = errors.New("passsword is not decode from base64")
)

type ValidationError struct {
	Field string
	Code  ErrCode
	Err   error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %d", e.Field, e.Code)
}

func (e ValidationError) Unwrap() error {
	if e.Err == nil {
		return errorsByCode[e.Code]
	}
	return e.Err
}

var errorsByCode = []error{
	EmptyName:             ErrNameEmpty,
	NameUnchanged:         ErrNameUnchanged,
	StartsWithDigit:       ErrStartsWithDigit,
	MemberNotFound:        ErrMemberNotFound,
	MaxMemberCount:        ErrMaxMemberCount,
	MinMemberCount:        ErrMinMemberCount,
	MemberIsExists:        ErrMemberIsExists,
	MemberIsNotExists:     ErrMemberIsNotExists,
	PasswordLength:        ErrPasswordLength,
	GenPasswordSalt:       ErrGenPasswordSalt,
	GenPasswordHash:       ErrGenPasswordHashed,
	InvalidPassword:       ErrInvalidPassword,
	DecodePasswordFromB64: ErrDecodePasswordFromB64,
	Unknown:               nil,
}
