@0xbeef480ffc0abb72;

using Go = import "/go.capnp";
$Go.package("capnp");
$Go.import("rpc");

interface ClientStreamingV2 {
    doStreaming @0 () -> (callback :Callback);

    interface Callback {
        sendMessage @0 (clientStreamingMsg :ClientStreamingMessageV2) -> ();
    }
}

struct ClientStreamingMessageV2 {
    chunk @0 :Data;
}