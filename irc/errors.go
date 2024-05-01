package irc

import "fmt"

const (
	ErrCodeParseMsg = iota
	ErrCodeConnection
	ErrTooManyAuthors
)

type IrcError struct {
	Message string
	Code    int
}

func (e *IrcError) Error() string {
	return fmt.Sprintf("err: code %d | msg: %s", e.Code, e.Message)
}

func NewError(msg string, code int) *IrcError {
	return &IrcError{
		Message: msg,
		Code:    code,
	}
}
