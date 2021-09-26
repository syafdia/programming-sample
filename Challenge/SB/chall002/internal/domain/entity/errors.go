package entity

import (
	"errors"
	"fmt"
)

// List of error.
var (
	ErrNotImplemented = errors.New("not implemented")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrBadGateway     = errors.New("bad gateway")
	ErrUnknown        = errors.New("unknown error")
)

type Error struct {
	err    error
	detail string
}

func NewError(err error, detail string) *Error {
	return &Error{err: err, detail: detail}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: %s", e.err)
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) GetDetail() string {
	return e.detail
}
