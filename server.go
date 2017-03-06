package main

import (
	"bufio"
	c "comm"
	"fmt"
	"net"
)

func main() {

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		fmt.Print("Message from server: " + message)
	}
}
func main() {
	conn := c.Server(8081)
	conn.ServStart()
}
