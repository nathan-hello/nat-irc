package irc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"time"
)

type ConnectParams struct {
	Server string
	Port   string
}

func Connect(opts *ConnectParams) (*tls.Conn, *bufio.ReadWriter) {
	conn, err := tls.Dial("tcp", opts.Server+":"+opts.Port, &tls.Config{})
	if err != nil {
		panic(fmt.Sprint("Error connecting to IRC server: ", err))
	}

	wr := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	return conn, wr
}

func SendMessage(c *tls.Conn, rw *bufio.ReadWriter, cmd string) {
	fmt.Fprint(rw.Writer, cmd)
	rw.Writer.Flush()
}

func ServerMessages(rw *bufio.ReadWriter) {
	for {
		message, err := rw.Reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Print(message)

	}
}

func TestConnection(c *tls.Conn, rw *bufio.ReadWriter) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		fmt.Fprintf(rw.Writer, "PING :keepalive\r\n")
		rw.Writer.Flush()
		fmt.Println("Sent PING message to keep the connection alive.")
	}
}
