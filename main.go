package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nathan-hello/nat-irc/irc"
)

func main() {
	conn, rw := irc.Connect(&irc.ConnectParams{Server: "irc.reluekiss.com", Port: "6697"})
	defer conn.Close()
	// go irc.TestConnection(conn, rw)
	go irc.OutputServerMessages(rw, os.Stdout)

	for {
		stdin := bufio.NewReader(os.Stdin)
		fmt.Println("CMD:")
		text, err := stdin.ReadString('\n')
		if err != nil {
			panic(err)
		}

		irc.SendMessage(conn, rw, text)

	}

}
