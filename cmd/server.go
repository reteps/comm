package main

import (
	"fmt"
	c "github.com/reteps/comm/comm"
)

func main() {
	conn := c.Start("8081")
	for {
		message := conn.Read()
		fmt.Println("Got this from client: " + message)
		conn.Send(message + "-FROM_SERVER")
	}
}
