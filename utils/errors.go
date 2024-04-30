package utils

import (
	"fmt"
)

type IrcError struct {
	s    string
	Code int
}

func (e IrcError) Error() string {
	return e.s
}

const (
	ErrCodeParseMsg = iota
	ErrCodeConnection
)

func ErrParseMsg(msg string) IrcError {
	return IrcError{
		s:    fmt.Sprintf("invalid PRIVMSG - %#v", msg),
		Code: ErrCodeParseMsg,
	}

}
