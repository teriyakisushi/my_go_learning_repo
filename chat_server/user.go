package main

import (
	"net"
)

type User struct {
	userAddress string
	conn        net.Conn
	chanUser    chan string
}

func newUser(conn net.Conn) *User {
	userAddress := conn.RemoteAddr().String()
	user := &User{
		userAddress: userAddress,
		conn:        conn,
		chanUser:    make(chan string),
	}
	go user.ListenMessage()
	return user
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.chanUser
		user.conn.Write([]byte(msg + "\n"))
	}
}

func (user *User) HandleMessage(server *Server) {
	buf := make([]byte, 4096)
	for {
		n, err := user.conn.Read(buf)
		if n == 0 {
			server.UserOffline(user)
			return
		}
		if err != nil {
			return
		}
		msg := string(buf[:n-1])
		server.Broadcast(user, msg)
	}
}
