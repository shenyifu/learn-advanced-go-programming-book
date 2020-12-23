package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listen, err := net.Listen("tcp", "1234")
	if err != nil {
		log.Fatal("Listen Tcp error", err)
	}

	conn, err := listen.Accept()
	if err != nil {
		log.Fatal("accept error", err)
	}
	rpc.ServeConn(conn)
}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "hello :" + request
	return nil
}
