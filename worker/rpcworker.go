package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	//should have comment
	sertype int
}

type Ret struct {
	back string
}

type Arith int

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	args := &Args{1}
	var reply bool
	err = client.Call("arith.Check", args, &reply)
	if err != nil {
		log.Fatal("Check error", err)
	}
	fmt.Printf("Check: type %d", args.sertype, reply)

	args = &Args{0}
	var getRet Ret
	err = client.Call("arith.Serve", args, &getRet)
	if err != nil {
		log.Fatal("Serve error:", err)
	}
	fmt.Printf("Serve %s", getRet.back)

}
