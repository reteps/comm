package main

import (
	c "comm"
	"fmt"
)

func main() {
	conn := c.Start("8081")
	for {
		message := conn.Read()
		fmt.Println("Got this from client: " + message)
		conn.ServSend(message + "-FROM_SERVER")
	}
}
