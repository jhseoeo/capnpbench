package main

import (
	"capnpbench/capnp"
	"capnpbench/config"
	"context"
	"fmt"
	"net"

	"capnproto.org/go/capnp/v3/rpc"
)

func runClientStreamingV2Client() {
	ctx := context.Background()
	conn, err := net.Dial("tcp", config.ClientAddr)
	if err != nil {
		panic(err)
	}
	rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	client := capnp.ClientStreamingV2(rpcConn.Bootstrap(ctx))

	future, release := client.DoStreaming(ctx, func(csv capnp.ClientStreamingV2_doStreaming_Params) error {
		return nil
	})
	defer release()

	for i := 0; i < 10; i++ {
		_, r2 := future.Callback().SendMessage(ctx, func(csv capnp.ClientStreamingV2_Callback_sendMessage_Params) error {
			msg, err := capnp.NewClientStreamingMessageV2(csv.Segment())
			if err != nil {
				return err
			}
			msg.SetChunk([]byte(fmt.Sprintf("Hello, server! %d", i)))
			csv.SetClientStreamingMsg(msg)
			return nil
		})
		r2()
	}
}

func main() {
	runClientStreamingV2Client()
}
