package tcp

import (
	"fmt"
	"net"
)

type Server struct {
	IpAddr string
	Port   uint16
	listen net.Listener
	handle Handler
}

func (tcp *Server) Listener(IpAddress string, Port uint16) {
	listen, err := net.Listen("tcp", IpAddress)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	tcp.listen = listen
}

func (tcp *Server) Accept() {
	for {
		conn, err := tcp.listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		tcp.handle.OnProcess(conn)
	}
}
