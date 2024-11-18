package main

import (
	"fmt"
	"net"
	"sync"
)

/*上线服务器需要的IP & 端口*/
type Server struct {
	IP         string
	PORT       int
	online_map map[string]*User // <--- 服务器上线用户的映射
	map_lock   sync.RWMutex     // <--- 服务器上线用户的映射锁
	public_msg chan string      // <--- 服务器广播消息的管道
}

/*服务器启动方法*/
func (this *Server) Start() {
	listener, err := net.Listen(
		"tcp4",
		fmt.Sprintf("%s:%d", this.IP, this.PORT),
	)
	if err != nil {
		fmt.Println("Launch Failed!")
		return
	}
	defer listener.Close() // <---执行完后要服务器占用的监听地址

	// 启动服务器广播消息监听
	go this.ListenMessage()

	// 等待Client连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connect failed! Retrying..")
			continue
		}

		go this.Handler(conn) // <---连接成功后交给Handler处理事务

	}
}

/*服务器广播消息*/
func (this *Server) Broadcast(user *User, message string) {
	sendMsg := "[" + user.userAddress + "] :" + message
	fmt.Println(sendMsg)
	this.public_msg <- sendMsg
}

/*服务器的public msg管道监听是否有来自服务器广播消息*/
func (this *Server) ListenMessage() {
	for {
		msg := <-this.public_msg // <--- 不断监听服务器广播消息
		this.map_lock.Lock()     // <--- 上锁
		for _, user := range this.online_map {
			user.chanUser <- msg // <--- 遍历服务器上线用户的映射，然后给他们的channel发送消息
		}
		this.map_lock.Unlock() // <--- 发送完消息后解锁
	}
}

/*
*用来处理连接成功后的事务
*接收的参数应该是连接成功后目标客户端的conn
 */
func (this *Server) Handler(conn net.Conn) {

	_, err := conn.Write([]byte("Howdy today!")) // <--- Client Message
	if err != nil {
		fmt.Println("Write failed!")
		return
	}

	user := newUser(conn)       // <--- 创建一个新用户
	this.UserOnline(user)       // <--- 用户上线
	go user.HandleMessage(this) // <--- 处理用户消息

	select {} //<---- 阻塞
}

/*用户上线*/
func (this *Server) UserOnline(user *User) {
	this.map_lock.Lock()
	this.online_map[user.userAddress] = user
	this.map_lock.Unlock()
	this.Broadcast(user, "Join in !")
}

/*用户下线*/
func (this *Server) UserOffline(user *User) {
	this.map_lock.Lock()
	delete(this.online_map, user.userAddress)
	this.map_lock.Unlock()
	this.Broadcast(user, "Offline !")
}

/*创建服务器方法，返回一个创建好的服务器实例*/
func newServer(server_ip string, server_port int) *Server {
	server := &Server{
		IP:         server_ip,
		PORT:       server_port,
		online_map: make(map[string]*User),
		map_lock:   sync.RWMutex{},
		public_msg: make(chan string),
	}

	return server // <--- 返回服务器实例
}
