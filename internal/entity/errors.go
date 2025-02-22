package entity

import "errors"

var (
	ErrAirstripIsFull   = errors.New("airstrip is full")
	ErrAirstripNotFound = errors.New("airstrip not found")
)
