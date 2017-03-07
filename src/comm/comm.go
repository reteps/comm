package comm

import "net"
import "fmt"
import "bufio"
import "os"

type Connection struct {
	Link net.Conn
}

func Start(port string) Connection {

	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening for connection...")
	connection, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Found!")
	return Connection{connection}
}
func (c Connection) Read() string {
	message, err := bufio.NewReader(c.Link).ReadString(byte('¬'))
	if err != nil {
		panic(err)
	}
	return message[:len(message)-1] //no strange character at end
}
func (c Connection) ServSend(text string) {

	c.Link.Write([]byte(text + "¬"))
}

func Connect(ip, port string) Connection {
	connection, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Connecting - The server has either not been started, or you have the wrong port")
		os.Exit(1)
	}
	fmt.Println("Connection Found!")
	return Connection{connection}
}
func (c Connection) CliSend(message string) {
	fmt.Fprintf(c.Link, message+"¬")
}
