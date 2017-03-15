# comm
a server + client connection library

This program was based off the code [here](https://systembash.com/a-simple-go-tcp-server-and-tcp-client/).

Commands:

`connection := Start(port)`
 + starts a server
 
`connection := Connect(ip,port)`
 + connects to a server
 
`message := connection.Read()`
 + listens for a message and reads it
 
`connection.Send(message)`
 + sends a message to the server / client

