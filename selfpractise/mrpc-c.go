package main

import (
	"net/rpc"
	"log"
)

func main() {
	client, err := rpc.Dial("tcp", ":12345")
	if err != nil {
		log.Fatal("dialing", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "Lilei", &reply)
	if err != nil {
		log.Fatal(err)
	}

	println(reply)
}
