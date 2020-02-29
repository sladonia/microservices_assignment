package storage

import (
	"errors"
	"fmt"
)

var ItemNotFoundBasicError = errors.New("record not found")

type ItemNotFoundError struct {
	Err error
	Msg string
}

func NewItemNotFoundError(msg string) *ItemNotFoundError {
	return &ItemNotFoundError{
		Err: ItemNotFoundBasicError,
		Msg: msg,
	}
}

func (e *ItemNotFoundError) Error() string {
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s: %s", e.Err.Error(), e.Msg)
}

func (e *ItemNotFoundError) Unwrap() error {
	return e.Err
}
