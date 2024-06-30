@0x8b7d18a4959e8756;

using Go = import "/go.capnp";
$Go.package("capnp");
$Go.import("rpc");

interface ClientStreaming {
    doStreaming @0 (clientStreamingMsg :ClientStreamingMessage) -> stream;
	done @1 ();
}

struct ClientStreamingMessage {
    chunk @0 :Data;
}
