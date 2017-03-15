package main

import (
        "comm/multicomm"
)
var Clients []net.conn
func handle_client(client) {
        if CLOSED == true {
                client.send(Packet{Type:PacketExit,Contents:"The server is not accepting further connections."})
        }
        Clients = append(Clients, client)
        
