package main

import (
	"bufio"
	c "comm"
	"fmt"
	"net"
)

func main() {
	conn := c.Server(8081)
	conn.ServStart()
	for {
		message := c.ServRead()
		fmt.Println("Got this from client" + message)
		c.ServSend(message + "-FROM_SERVER")
	}
}
