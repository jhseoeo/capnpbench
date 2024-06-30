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

type bidirectionalStreamingClientCallback struct{}

func (bs *bidirectionalStreamingClientCallback) SendMessage(ctx context.Context, call capnp.BidirectionalStreaming_ClientCallback_sendMessage) error {
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

type bidirectionalStreaming struct{}

func (b *bidirectionalStreaming) DoServerStreaming(ctx context.Context, call capnp.BidirectionalStreaming_doServerStreaming) error {
	serverCallback := call.Args().ServerCallback()
	for i := 0; i < 10; i++ {
		_, release := serverCallback.SendMessage(ctx, func(bs capnp.BidirectionalStreaming_ServerCallback_sendMessage_Params) error {
			msg, err := bs.NewStreamingMsg()
			if err != nil {
				return err
			}
			msg.SetChunk([]byte(fmt.Sprintf("Hello, client! %d", i)))
			bs.SetStreamingMsg(msg)
			return nil
		})
		release()
	}
	return nil
}

func (b *bidirectionalStreaming) DoClientStreaming(ctx context.Context, call capnp.BidirectionalStreaming_doClientStreaming) error {
	res, err := call.AllocResults()
	if err != nil {
		return err
	}
	res.SetClientCallback(capnp.BidirectionalStreaming_ClientCallback_ServerToClient(&bidirectionalStreamingClientCallback{}))
	return nil
}

func runBidirectionalStreamingServer() error {
	service := &bidirectionalStreaming{}
	client := capnp.BidirectionalStreaming_ServerToClient(service)

	ln, err := net.Listen("tcp", config.ServerAddr)
	if err != nil {
		panic(err)
	}

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
	runBidirectionalStreamingServer()
}
