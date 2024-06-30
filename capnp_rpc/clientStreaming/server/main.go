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

// capnp.ClientStreaming_Server 인터페이스를 구현하는 구조체
type clientStreaming struct{}

func (cs *clientStreaming) DoStreaming(ctx context.Context, call capnp.ClientStreaming_doStreaming) error {
	msg, err := call.Args().ClientStreamingMsg()
	if err != nil {
		return err
	}
	chunk, err := msg.Chunk()
	if err != nil {
		return err
	}
	fmt.Println(string(chunk))

	_, err = call.AllocResults()
	if err != nil {
		return err
	}
	return nil
}

func (ss *clientStreaming) Done(ctx context.Context, call capnp.ClientStreaming_done) error {
	return nil
}

func runClientStreamingServer() error {
	service := &clientStreaming{}
	client := capnp.ClientStreaming_ServerToClient(service)

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
	runClientStreamingServer()
}
