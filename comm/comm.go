package comm

import "net"
import "fmt"
import "bufio"
import "os"

type Connection struct {
	Link net.Conn
	IsClient bool
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
	return Connection{connection,false}
}
func (c Connection) Read() string {
	message, err := bufio.NewReader(c.Link).ReadString(byte('`'))
	if err != nil {
		fmt.Println("The connection was interrupted by a loss of connection or control + c.")
		os.Exit(2)
	}
	return message[:len(message)-1] //no strange character at end
}
func (c Connection) Send(text string) {
	if IsClient {
		fmt.Fprintf(c.Link, text + "`")
	} else {
		c.Link.Write([]byte(text + "`"))
	}
}

func Connect(ip, port string) Connection {
	connection, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Connecting - The server has either not been started, or you have the wrong port")
		os.Exit(1)
	}
	fmt.Println("Connection Found!")
	return Connection{connection,true}
}

