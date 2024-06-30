@0xcd6137573ae9a2cb;

using Go = import "/go.capnp";
$Go.package("capnp");
$Go.import("rpc");

interface ServerStreaming {
  # doStreaming @0 (stream) -> serverStreamingMsg :ServerStreamingMessage;
  # ㅋㅋ 절대 안되지롱

  doStreaming @0 (callback :Callback) -> ();

  interface Callback {
    sendMessage @0 (serverStreamingMsg :ServerStreamingMessage) -> ();
  }
}

struct ServerStreamingMessage {
  chunk @0 :Data;
}
