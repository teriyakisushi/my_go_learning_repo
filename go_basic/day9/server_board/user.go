package main

import "net"

type User struct {
	userAddress string      // <--- 用户上线的地址
	userConn    net.Conn    // <--- 用户连接的套接字
	chanUser    chan string // <--- 用户通信的管道
}

/*
 * 创建用户方法
 * @param userConn 用户连接的套接字
 * @return *User 返回一个创建好的用户实例
 */
func newUser(userConn net.Conn) *User {
	new_user := &User{
		userAddress: userConn.RemoteAddr().String(),
		userConn:    userConn,
		chanUser:    make(chan string),
	}

	go new_user.userListener() // <--- 启动用户监听器

	return new_user
}

// 监听服务器的消息
func (this *User) userListener() {
	for {
		msg := <-this.chanUser
		this.userConn.Write([]byte(msg + "\n"))
	}
}
