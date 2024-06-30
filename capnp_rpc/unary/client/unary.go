package main

import (
	"capnpbench/capnp"
	"capnpbench/config"
	"context"
	"net"

	"capnproto.org/go/capnp/v3/rpc"
)

func runUnaryClient() {
	ctx := context.Background()
	conn, err := net.Dial("tcp", config.ClientAddr)
	if err != nil {
		panic(err)
	}
	rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	client := capnp.Unary(rpcConn.Bootstrap(ctx))

	future, release := client.DoSomething(ctx, func(u capnp.Unary_doSomething_Params) error {
		reqMessage, err := u.NewReqMessage()
		if err != nil {
			return err
		}
		reqMessage.SetChunk([]byte("Hello, server!"))
		u.SetReqMessage(reqMessage)
		return nil
	})
	defer release()
	res, err := future.Struct()
	if err != nil {
		panic(err)
	}
	msg, err := res.RespMessage()
	if err != nil {
		panic(err)
	}
	chunk, err := msg.Chunk()
	if err != nil {
		panic(err)
	}
	println(string(chunk))
}

func main() {
	runUnaryClient()
}
