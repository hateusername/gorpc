package main

import (
	"log"
	"net"
	"net/http"
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

func (t *Arith) Check(args *Args, reply *bool) error {
	*reply = true
	return nil
}

func (t *Arith) Serve(args *Args, reply *Ret) error {
	reply.Retback = string("服务类型为" + string(args.Sertype))
	return nil
}

func main() {
	ari := new(Arith)
	rpc.Register(ari)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error", e)
	}

	http.Serve(l, nil)
}
