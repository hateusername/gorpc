package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	//should have comment
	Sertype int
}

type Ret struct {
	Retback string
}

type Arith int

func main() {
	fmt.Println("这里是0")

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	fmt.Println("这里是1")

	args := &Args{1}
	var reply bool
	err = client.Call("arith.Check", args, &reply)
	if err != nil {
		log.Fatal("Check error", err)
	}
	fmt.Printf("Check: type %d", args.Sertype, reply)

	args = &Args{0}
	var getRet Ret
	err = client.Call("arith.Serve", args, &getRet)
	if err != nil {
		log.Fatal("Serve error:", err)
	}
	fmt.Printf("Serve %s", getRet.Retback)

}
