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

func (user *User) Send2User(msg string) {
	user.chanUser <- msg
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
		// 处理用户消息,不广播
		if msg == "who" && len(msg) == 3 {
			user.conn.Write([]byte("online user list:\n"))
			server.map_lock.RLock() // 上锁
			for _, u := range server.online_map {
				onlineMsg := u.userAddress + " Online\n"
				// 谁查询的就给谁发
				user.Send2User(onlineMsg)
			}
			server.map_lock.RUnlock() // 解锁
		} else if len(msg) > 7 && msg[:7] == "rename|" {
			newName := msg[7:]
			if _, ok := server.online_map[newName]; ok {
				user.Send2User("User name already exists!")
			} else {
				server.map_lock.Lock()
				delete(server.online_map, user.userAddress)
				user.userAddress = newName
				server.online_map[newName] = user
				server.map_lock.Unlock()
				user.Send2User("Rename successful!")
			}
		} else {
			server.Broadcast(user, msg)
		}
	}
}
