package main

import (
	"capnpbench/capnp"
	"capnpbench/config"
	"context"
	"fmt"
	"net"

	"capnproto.org/go/capnp/v3/rpc"
)

type callback struct{}

func (cb *callback) SendMessage(ctx context.Context, s capnp.ServerStreaming_Callback_sendMessage) error {
	msg, err := s.Args().ServerStreamingMsg()
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

func runServerStreamingClient() {
	ctx := context.Background()
	conn, err := net.Dial("tcp", config.ClientAddr)
	if err != nil {
		panic(err)
	}
	callback := capnp.ServerStreaming_Callback_ServerToClient(&callback{})
	rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	client := capnp.ServerStreaming(rpcConn.Bootstrap(ctx))

	future, release := client.DoStreaming(ctx, func(s capnp.ServerStreaming_doStreaming_Params) error {
		s.SetCallback(callback)
		return nil
	})
	defer release()
	_, err = future.Struct()
	if err != nil {
		panic(err)
	}
}

func main() {
	runServerStreamingClient()
}
