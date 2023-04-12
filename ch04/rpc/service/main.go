package main

import (
	"log"
	"net"
	"net/rpc"

	myRpc "ad/ch04/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(myRpc.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
