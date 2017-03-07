package comm

import "net"
import "fmt"
import "bufio"

type Server struct {
	Connection net.Conn
}

func ServStart(port string) {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", port)
	fmt.Println("Listening for connection...")
	connection, _ := ln.Accept()
	fmt.Println("Connection Found!")
	return Server{connection}
}
func (s Server) ServRead() {
	message, _ := bufio.NewReader(s.Connection).ReadString('\n')
	return string(message)
}
func (s Server) ServSend(message string) {

	s.Connection.Write([]byte(message + "\nEND"))
}

type Client struct {
	Ip   string
	Port string
}

func (c Client) CliConnect() {
	fmt.Println("Looking for a server")
	c.Connection, _ = net.Dial("tcp", Ip+":"+Port)
	fmt.Println("Connection Found!")
}
func (c Client) CliRead() {
	message, _ := bufio.NewReader(c.Connection).ReadString("\nEND")
	return string(message)
}
func (c Client) CliSend(message string) {
	fmt.Fprintf(c.Connection, message+"\nEND")
}
