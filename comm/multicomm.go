package multicomm

const (
        PacketId iota
        PacketName
        PacketMessage
        PacketExit
)
type Packet struct {
        From int 
        Type int
        Contents interface{}
        To int
}
        
