@0x8836f289aabc01d9;

using Go = import "/go.capnp";
$Go.package("capnp");
$Go.import("rpc");

interface BidirectionalStreaming {
  # doStreaming @0 (serverCallback :ServerCallback) -> (clientCallback :ClientCallback);
  # ㅋㅋ 절대 안되지롱
  doServerStreaming @0 (serverCallback :ServerCallback) -> ();
  doClientStreaming @1 () -> (clientCallback :ClientCallback);


  interface ServerCallback {
    sendMessage @0 (streamingMsg :BidirectionalStreamingMessage) -> ();
  }

  interface ClientCallback {
    sendMessage @0 (streamingMsg :BidirectionalStreamingMessage) -> ();
  }
}

struct BidirectionalStreamingMessage {
  chunk @0 :Data;
}