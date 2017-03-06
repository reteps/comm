package comm

import "net"
import "fmt"
import "bufio"

type Server struct {
	Port string
}

func (s Server) ServStart() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", s.Port)
	fmt.Println("Listening for connection...")
	s.Connection = ln.Accept()
	fmt.Println("Connection Found!")
}
func (s Server) ServRead() {
	message, _ := bufio.NewReader(s.Connection).ReadString("\nEND")
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
