package main

func main() {
	server := newServer("127.0.0.1", 8888)
	server.Start()
}
