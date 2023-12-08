package tcp

import "net"

type Handler interface {
	OnProcess(conn net.Conn)
	DoPacket()
}
