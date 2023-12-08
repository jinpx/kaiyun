package tcp

import (
	"fmt"
	"net"
)

type Client struct {
	IpAddr string
	Port   uint16
	conn   net.Conn
}

func (tcp *Client) ConnectTo(IpAddress string, Port uint16) {

	// conn, err := net.Dial("tcp", "127.0.0.1:20000")
	conn, err := net.Dial("tcp", tcp.IpAddr)
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	tcp.conn = conn
}

func (tcp *Client) Reader(sendData []byte) {
	for {
		buf := [512]byte{}
		n, err := tcp.conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func (tcp *Client) Write(sendData []byte) {
	_, err := tcp.conn.Write(sendData) // 发送数据
	if err != nil {
		return
	}
}
