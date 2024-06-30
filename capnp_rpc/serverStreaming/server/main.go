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

type serverStreaming struct{}

func (ss *serverStreaming) DoStreaming(ctx context.Context, call capnp.ServerStreaming_doStreaming) error {
	callback := call.Args().Callback()
	for i := 0; i < 10; i++ {
		callback.SendMessage(ctx, func(ss capnp.ServerStreaming_Callback_sendMessage_Params) error {
			msg, err := ss.NewServerStreamingMsg()
			if err != nil {
				return err
			}
			msg.SetChunk([]byte(fmt.Sprintf("Hello, client! %d", i)))
			ss.SetServerStreamingMsg(msg)
			return nil
		})
	}
	return nil
}

func runServerStreamingServer() error {
	service := &serverStreaming{}
	client := capnp.ServerStreaming_ServerToClient(service)

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
	runServerStreamingServer()
}
