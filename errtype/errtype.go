package main

import (
	"errors"
	"fmt"
)

type ErrorType string

type errWithType struct {
	msg string
	t   ErrorType
}

func NewErrorf(t ErrorType, format string, args ...any) error {
	return &errWithType{
		msg: fmt.Sprintf(format, args...),
		t:   t,
	}
}

func (e *errWithType) Error() string {
	return e.msg
}

func GetType(err error) ErrorType {
	var e *errWithType
	if errors.As(err, &e) {
		return e.t
	}
	return ""
}
