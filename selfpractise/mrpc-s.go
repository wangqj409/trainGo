package main

import (
	"net/rpc"
	"net"
	"log"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	con, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(con)
}
