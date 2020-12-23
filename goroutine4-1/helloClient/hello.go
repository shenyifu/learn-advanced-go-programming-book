package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialog", err)
	}

	var replay string

	err = client.Call("HelloService.Hello", "hello", &replay)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(replay)

}
