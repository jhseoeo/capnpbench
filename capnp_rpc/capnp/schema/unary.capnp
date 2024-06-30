@0xd4f5c30d909da4d9;

using Go = import "/go.capnp";
$Go.package("capnp");
$Go.import("rpc");

interface Unary {
    doSomething @0 (reqMessage :ReqMessage) -> (respMessage :RespMessage);
}

struct ReqMessage {
    chunk @0 :Data;
}

struct RespMessage {
    chunk @0 :Data;
}