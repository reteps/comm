package main

import (
	"bufio"
	c "comm"
	"fmt"
	"os"
)

func main() {
	conn := c.Client()
	conn.CliConnect()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		conn.CliSend(text)
		conn.Cli
