package main

import "net"

func main() {

	listener, err := net.Listen("http", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	conn.Write([]byte("hello\n"))
}
