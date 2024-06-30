package main

import (
	"capnpbench/capnp"
	"capnpbench/config"
	"context"
	"fmt"
	"net"

	capn "capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
)

// capnp.Unary_Server 인터페이스를 구현하는 구조체
type unaryServer struct{}

func (s *unaryServer) DoSomething(ctx context.Context, call capnp.Unary_doSomething) error {
	reqMsg, err := call.Args().ReqMessage()
	if err != nil {
		fmt.Println(err)
	}
	chunk, err := reqMsg.Chunk()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(chunk))

	res, err := call.AllocResults()
	if err != nil {
		return err
	}
	respMsg, err := res.NewRespMessage()
	if err != nil {
		return err
	}
	respMsg.SetChunk([]byte("Hello, client!"))
	res.SetRespMessage(respMsg)
	return nil
}

func runUnaryServer() error {
	service := &unaryServer{}
	client := capnp.Unary_ServerToClient(service)

	ln, err := net.Listen("tcp", config.ServerAddr)
	if err != nil {
		panic(err)
	}

	// 아 고루틴으로 뺴기 귀찮아
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), &rpc.Options{
			BootstrapClient: capn.Client(client.AddRef()),
		})

		select {
		case <-rpcConn.Done():
			continue
		case <-ctx.Done():
			rpcConn.Close()
			continue
		}
	}
}

func main() {
	runUnaryServer()
}
