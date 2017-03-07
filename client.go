package main

import (
	"bufio"
	c "comm"
	"fmt"
	"os"
)

func main() {
	conn := c.Connect("45.55.91.237", "8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		conn.CliSend(text[:len(text)-1])
		message := conn.Read()
		fmt.Println(message)
	}
}
