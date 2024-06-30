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

type clientStreamingV2Callback struct{}

func (cs *clientStreamingV2Callback) SendMessage(ctx context.Context, call capnp.ClientStreamingV2_Callback_sendMessage) error {
	msg, err := call.Args().ClientStreamingMsg()
	if err != nil {
		return err
	}
	chunk, err := msg.Chunk()
	if err != nil {
		return err
	}
	fmt.Println(string(chunk))
	return nil
}

// capnp.ClientStreaming_Server 인터페이스를 구현하는 구조체
type clientStreamingV2 struct{}

func (cs *clientStreamingV2) DoStreaming(ctx context.Context, call capnp.ClientStreamingV2_doStreaming) error {
	res, err := call.AllocResults()
	if err != nil {
		return err
	}
	res.SetCallback(capnp.ClientStreamingV2_Callback_ServerToClient(&clientStreamingV2Callback{}))
	return nil
}

func runClientStreamingV2Server() error {
	service := &clientStreamingV2{}
	client := capnp.ClientStreamingV2_ServerToClient(service)

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
	runClientStreamingV2Server()
}
