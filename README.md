# go-comm
a simple server + client connection

This program was based off of `https://systembash.com/a-simple-go-tcp-server-and-tcp-client/`.

Commands:

`Connect(ip,port)`
 + connects to a server
`Start(port)`
 + starts a servver
`message := Read()`
 + listens for a message and reads it
`CliSend(message)`
 + sends a message from the client
`ServSend(message)`
 + sends a message from the server
