package main

import (
	"fmt"
	"net"
)

// Construct a Server
type Server struct {
	IP   string
	Port int
}

func newServer(IP string, port int) *Server {
	server := &Server{
		IP,
		port,
	}

	return server
}

// 用来处理连接成功后的事务,需要的参数为连接成功后的套接字
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("Connect Successful! ")
	ToClientMsg := "Howdy Today ^_^"
	_, err := conn.Write([]byte(ToClientMsg))
	if err != nil {
		fmt.Println("Write error!")
		return
	}
}

// Start Server
func (this *Server) Launch() {
	// Socket Listen，开启响应端口的服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("Server Start Failed!")
		return
	}
	defer listener.Close() // <--- 结束后要关闭listener

	for {
		conn, err := listener.Accept() //<--- 在for内等待端口连接
		if err != nil {
			fmt.Println("Connect failed! ")
			continue // <---继续
		}

		go this.Handler(conn) // <---创建一个goroutine，用来处理连接事务
	}
}
