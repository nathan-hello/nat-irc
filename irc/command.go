package irc

import (
	"fmt"
)

type ParsedCommand struct {
	Prefix  string
	FullCmd string
}

type Command struct {
	*ParsedCommand
	Validator       func(string) bool
	ErrorHandler    func(string) any
	Callback        func(string) any
	PotentialErrors []string
}

type PrivmsgParams struct {
	AutomaticReply bool
	Author         string
	Recipient      string
	Message        string
}

func Privmsg(p PrivmsgParams) (*ParsedCommand, *IrcError) {
	return &ParsedCommand{}, nil

}

type NickParams struct {
	PreviousName string
	NewName      string
}

func Nick(p NickParams) (*ParsedCommand, *IrcError) {
	if p.NewName == "" {
		return nil, &IrcError{Message: "no new nickname given"}
	}
	cmd := ""
	if p.PreviousName != "" {
		cmd += fmt.Sprintf(":%v ", p.PreviousName)
	}
	cmd += fmt.Sprintf("NICK %v;", p.NewName)

	return &ParsedCommand{
		Prefix:  "NICK",
		FullCmd: cmd,
	}, nil
}

type JoinParams struct {
	Author  string
	Channel string
	Key     string
}

func Join(p []JoinParams) (*ParsedCommand, *IrcError) {
	cmd := ""
	author := ""
	channels := []string{}
	keys := []string{}

	for _, j := range p {
		if author == "" {
			author = j.Author
		} else if author != j.Author {
			return nil, &IrcError{
				Code:    ErrTooManyAuthors,
				Message: fmt.Sprintf("too many authors in JoinParams struct - %#v", p),
			}
		}

		channels = append(channels, j.Channel)
		keys = append(channels, j.Key)
	}

	if author != "" {
		cmd += fmt.Sprintf(":%v ", author)
	}

	for i, v := range channels {
		cmd += fmt.Sprintf("%v ", v)
		if i < len(channels) {
			cmd += ","
		}
	}

	if len(keys) > 0 {
		cmd += " "
	}

	for i, v := range keys {
		cmd += fmt.Sprintf("%v ", v)
		if i < len(channels) {
			cmd += ","
		}
	}

	return &ParsedCommand{
		Prefix:  "JOIN",
		FullCmd: cmd,
	}, nil
}
