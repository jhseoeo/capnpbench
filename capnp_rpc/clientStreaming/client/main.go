package main

import (
	"capnpbench/capnp"
	"capnpbench/config"
	"context"
	"fmt"
	"net"

	"capnproto.org/go/capnp/v3/rpc"
)

func runClientStreamingClient() {
	ctx := context.Background()
	conn, err := net.Dial("tcp", config.ClientAddr)
	if err != nil {
		panic(err)
	}
	rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	client := capnp.ClientStreaming(rpcConn.Bootstrap(ctx))

	for i := 0; i < 10; i++ {
		err := client.DoStreaming(ctx, func(s capnp.ClientStreaming_doStreaming_Params) error {
			msg, err := s.NewClientStreamingMsg()
			if err != nil {
				return err
			}
			msg.SetChunk([]byte(fmt.Sprintf("Hello, server! %d", i)))
			s.SetClientStreamingMsg(msg)
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	future, release := client.Done(ctx, nil)
	defer release()
	_, err = future.Struct()
	if err != nil {
		panic(err)
	}
	if err := client.WaitStreaming(); err != nil {
		panic(err)
	}
}

func main() {
	runClientStreamingClient()
}
