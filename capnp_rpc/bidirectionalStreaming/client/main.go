package main

import (
	"capnpbench/capnp"
	"capnpbench/config"
	"context"
	"fmt"
	"net"

	"capnproto.org/go/capnp/v3/rpc"
)

type bidirectionalStreamingServerCallback struct{}

func (bs *bidirectionalStreamingServerCallback) SendMessage(ctx context.Context, call capnp.BidirectionalStreaming_ServerCallback_sendMessage) error {
	msg, err := call.Args().StreamingMsg()
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

func runBidirectionalStreamingClient() error {
	ctx := context.Background()
	conn, err := net.Dial("tcp", config.ClientAddr)
	if err != nil {
		panic(err)
	}
	serverCallback := capnp.BidirectionalStreaming_ServerCallback_ServerToClient(&bidirectionalStreamingServerCallback{})
	rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	client := capnp.BidirectionalStreaming(rpcConn.Bootstrap(ctx))

	_, r := client.DoServerStreaming(ctx, func(bs capnp.BidirectionalStreaming_doServerStreaming_Params) error {
		bs.SetServerCallback(serverCallback)
		return nil
	})
	r()

	f, r := client.DoClientStreaming(ctx, func(bs capnp.BidirectionalStreaming_doClientStreaming_Params) error {
		return nil
	})
	defer r()

	// go func() {
	for i := 0; i < 10; i++ {
		_, r := f.ClientCallback().SendMessage(ctx, func(bs capnp.BidirectionalStreaming_ClientCallback_sendMessage_Params) error {
			msg, err := bs.NewStreamingMsg()
			if err != nil {
				return err
			}
			msg.SetChunk([]byte(fmt.Sprintf("Hello, server! %d", i)))
			bs.SetStreamingMsg(msg)
			return nil
		})
		r()
	}
	// }()

	return nil
}

func main() {
	runBidirectionalStreamingClient()
}
