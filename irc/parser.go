package irc

import "github.com/nathan-hello/nat-irc/utils"

type PrivmsgParams struct {
	Command
	Args struct {
		Recipient string
		Message   string
	}
}

func ParsePrivmsg(p PrivmsgParams) (string, utils.IrcError) {

	return "", utils.ErrParseMsg("")
}
