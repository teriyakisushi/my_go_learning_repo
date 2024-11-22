package main

import (
	"fmt"
	"net"
)

/*客户端需要的IP & 端口 & Conn*/
type Client struct {
	ServerIP   string
	ServerPort int
	Name       string // <--- 客户端的名字
	Conn       net.Conn
}

/*客户端连接服务器*/
func NewClient(ServerIP string, ServerPort int) *Client {
	client := &Client{
		ServerIP:   ServerIP,
		ServerPort: ServerPort,
		Name:       "Anonymous",
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerIP, ServerPort))
	if err != nil {
		fmt.Println("Dial failed:", err)
		return nil
	} else {
		client.Conn = conn
		fmt.Println("Dial success!")
	}
	return client
}

func main() {
	client := NewClient(
		"127.0.0.1",
		8888,
	)
	if client == nil {
		fmt.Println("Connect to server failed!")
		return
	}
	fmt.Println("Connect to server success!")
	select {} // <--- 用来阻塞
}
