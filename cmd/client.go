package main

import (
	"bufio"
	"fmt"
	c "github.com/reteps/comm/comm"
	"os"
)

func main() {
	conn := c.Connect("45.55.91.237", "8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		conn.Send(text[:len(text)-1])
		message := conn.Read()
		fmt.Println(message)
	}
}
