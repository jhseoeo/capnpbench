// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
)

type ClientStreamingV2 capnp.Client

// ClientStreamingV2_TypeID is the unique identifier for the type ClientStreamingV2.
const ClientStreamingV2_TypeID = 0x91a1e8bbf5ca27f0

func (c ClientStreamingV2) DoStreaming(ctx context.Context, params func(ClientStreamingV2_doStreaming_Params) error) (ClientStreamingV2_doStreaming_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x91a1e8bbf5ca27f0,
			MethodID:      0,
			InterfaceName: "capnp/schema/clientStreamingv2.capnp:ClientStreamingV2",
			MethodName:    "doStreaming",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(ClientStreamingV2_doStreaming_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return ClientStreamingV2_doStreaming_Results_Future{Future: ans.Future()}, release

}

func (c ClientStreamingV2) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c ClientStreamingV2) String() string {
	return "ClientStreamingV2(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c ClientStreamingV2) AddRef() ClientStreamingV2 {
	return ClientStreamingV2(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c ClientStreamingV2) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c ClientStreamingV2) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c ClientStreamingV2) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (ClientStreamingV2) DecodeFromPtr(p capnp.Ptr) ClientStreamingV2 {
	return ClientStreamingV2(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c ClientStreamingV2) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c ClientStreamingV2) IsSame(other ClientStreamingV2) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c ClientStreamingV2) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c ClientStreamingV2) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A ClientStreamingV2_Server is a ClientStreamingV2 with a local implementation.
type ClientStreamingV2_Server interface {
	DoStreaming(context.Context, ClientStreamingV2_doStreaming) error
}

// ClientStreamingV2_NewServer creates a new Server from an implementation of ClientStreamingV2_Server.
func ClientStreamingV2_NewServer(s ClientStreamingV2_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(ClientStreamingV2_Methods(nil, s), s, c)
}

// ClientStreamingV2_ServerToClient creates a new Client from an implementation of ClientStreamingV2_Server.
// The caller is responsible for calling Release on the returned Client.
func ClientStreamingV2_ServerToClient(s ClientStreamingV2_Server) ClientStreamingV2 {
	return ClientStreamingV2(capnp.NewClient(ClientStreamingV2_NewServer(s)))
}

// ClientStreamingV2_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func ClientStreamingV2_Methods(methods []server.Method, s ClientStreamingV2_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x91a1e8bbf5ca27f0,
			MethodID:      0,
			InterfaceName: "capnp/schema/clientStreamingv2.capnp:ClientStreamingV2",
			MethodName:    "doStreaming",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.DoStreaming(ctx, ClientStreamingV2_doStreaming{call})
		},
	})

	return methods
}

// ClientStreamingV2_doStreaming holds the state for a server call to ClientStreamingV2.doStreaming.
// See server.Call for documentation.
type ClientStreamingV2_doStreaming struct {
	*server.Call
}

// Args returns the call's arguments.
func (c ClientStreamingV2_doStreaming) Args() ClientStreamingV2_doStreaming_Params {
	return ClientStreamingV2_doStreaming_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c ClientStreamingV2_doStreaming) AllocResults() (ClientStreamingV2_doStreaming_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingV2_doStreaming_Results(r), err
}

// ClientStreamingV2_List is a list of ClientStreamingV2.
type ClientStreamingV2_List = capnp.CapList[ClientStreamingV2]

// NewClientStreamingV2_List creates a new list of ClientStreamingV2.
func NewClientStreamingV2_List(s *capnp.Segment, sz int32) (ClientStreamingV2_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[ClientStreamingV2](l), err
}

type ClientStreamingV2_Callback capnp.Client

// ClientStreamingV2_Callback_TypeID is the unique identifier for the type ClientStreamingV2_Callback.
const ClientStreamingV2_Callback_TypeID = 0xf31dd865af9cf51c

func (c ClientStreamingV2_Callback) SendMessage(ctx context.Context, params func(ClientStreamingV2_Callback_sendMessage_Params) error) (ClientStreamingV2_Callback_sendMessage_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xf31dd865af9cf51c,
			MethodID:      0,
			InterfaceName: "capnp/schema/clientStreamingv2.capnp:ClientStreamingV2.Callback",
			MethodName:    "sendMessage",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(ClientStreamingV2_Callback_sendMessage_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return ClientStreamingV2_Callback_sendMessage_Results_Future{Future: ans.Future()}, release

}

func (c ClientStreamingV2_Callback) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c ClientStreamingV2_Callback) String() string {
	return "ClientStreamingV2_Callback(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c ClientStreamingV2_Callback) AddRef() ClientStreamingV2_Callback {
	return ClientStreamingV2_Callback(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c ClientStreamingV2_Callback) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c ClientStreamingV2_Callback) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c ClientStreamingV2_Callback) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (ClientStreamingV2_Callback) DecodeFromPtr(p capnp.Ptr) ClientStreamingV2_Callback {
	return ClientStreamingV2_Callback(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c ClientStreamingV2_Callback) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c ClientStreamingV2_Callback) IsSame(other ClientStreamingV2_Callback) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c ClientStreamingV2_Callback) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c ClientStreamingV2_Callback) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A ClientStreamingV2_Callback_Server is a ClientStreamingV2_Callback with a local implementation.
type ClientStreamingV2_Callback_Server interface {
	SendMessage(context.Context, ClientStreamingV2_Callback_sendMessage) error
}

// ClientStreamingV2_Callback_NewServer creates a new Server from an implementation of ClientStreamingV2_Callback_Server.
func ClientStreamingV2_Callback_NewServer(s ClientStreamingV2_Callback_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(ClientStreamingV2_Callback_Methods(nil, s), s, c)
}

// ClientStreamingV2_Callback_ServerToClient creates a new Client from an implementation of ClientStreamingV2_Callback_Server.
// The caller is responsible for calling Release on the returned Client.
func ClientStreamingV2_Callback_ServerToClient(s ClientStreamingV2_Callback_Server) ClientStreamingV2_Callback {
	return ClientStreamingV2_Callback(capnp.NewClient(ClientStreamingV2_Callback_NewServer(s)))
}

// ClientStreamingV2_Callback_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func ClientStreamingV2_Callback_Methods(methods []server.Method, s ClientStreamingV2_Callback_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf31dd865af9cf51c,
			MethodID:      0,
			InterfaceName: "capnp/schema/clientStreamingv2.capnp:ClientStreamingV2.Callback",
			MethodName:    "sendMessage",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.SendMessage(ctx, ClientStreamingV2_Callback_sendMessage{call})
		},
	})

	return methods
}

// ClientStreamingV2_Callback_sendMessage holds the state for a server call to ClientStreamingV2_Callback.sendMessage.
// See server.Call for documentation.
type ClientStreamingV2_Callback_sendMessage struct {
	*server.Call
}

// Args returns the call's arguments.
func (c ClientStreamingV2_Callback_sendMessage) Args() ClientStreamingV2_Callback_sendMessage_Params {
	return ClientStreamingV2_Callback_sendMessage_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c ClientStreamingV2_Callback_sendMessage) AllocResults() (ClientStreamingV2_Callback_sendMessage_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ClientStreamingV2_Callback_sendMessage_Results(r), err
}

// ClientStreamingV2_Callback_List is a list of ClientStreamingV2_Callback.
type ClientStreamingV2_Callback_List = capnp.CapList[ClientStreamingV2_Callback]

// NewClientStreamingV2_Callback_List creates a new list of ClientStreamingV2_Callback.
func NewClientStreamingV2_Callback_List(s *capnp.Segment, sz int32) (ClientStreamingV2_Callback_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[ClientStreamingV2_Callback](l), err
}

type ClientStreamingV2_Callback_sendMessage_Params capnp.Struct

// ClientStreamingV2_Callback_sendMessage_Params_TypeID is the unique identifier for the type ClientStreamingV2_Callback_sendMessage_Params.
const ClientStreamingV2_Callback_sendMessage_Params_TypeID = 0x8b4a4f5159acdd73

func NewClientStreamingV2_Callback_sendMessage_Params(s *capnp.Segment) (ClientStreamingV2_Callback_sendMessage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingV2_Callback_sendMessage_Params(st), err
}

func NewRootClientStreamingV2_Callback_sendMessage_Params(s *capnp.Segment) (ClientStreamingV2_Callback_sendMessage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingV2_Callback_sendMessage_Params(st), err
}

func ReadRootClientStreamingV2_Callback_sendMessage_Params(msg *capnp.Message) (ClientStreamingV2_Callback_sendMessage_Params, error) {
	root, err := msg.Root()
	return ClientStreamingV2_Callback_sendMessage_Params(root.Struct()), err
}

func (s ClientStreamingV2_Callback_sendMessage_Params) String() string {
	str, _ := text.Marshal(0x8b4a4f5159acdd73, capnp.Struct(s))
	return str
}

func (s ClientStreamingV2_Callback_sendMessage_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ClientStreamingV2_Callback_sendMessage_Params) DecodeFromPtr(p capnp.Ptr) ClientStreamingV2_Callback_sendMessage_Params {
	return ClientStreamingV2_Callback_sendMessage_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ClientStreamingV2_Callback_sendMessage_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ClientStreamingV2_Callback_sendMessage_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ClientStreamingV2_Callback_sendMessage_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ClientStreamingV2_Callback_sendMessage_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s ClientStreamingV2_Callback_sendMessage_Params) ClientStreamingMsg() (ClientStreamingMessageV2, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return ClientStreamingMessageV2(p.Struct()), err
}

func (s ClientStreamingV2_Callback_sendMessage_Params) HasClientStreamingMsg() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s ClientStreamingV2_Callback_sendMessage_Params) SetClientStreamingMsg(v ClientStreamingMessageV2) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewClientStreamingMsg sets the clientStreamingMsg field to a newly
// allocated ClientStreamingMessageV2 struct, preferring placement in s's segment.
func (s ClientStreamingV2_Callback_sendMessage_Params) NewClientStreamingMsg() (ClientStreamingMessageV2, error) {
	ss, err := NewClientStreamingMessageV2(capnp.Struct(s).Segment())
	if err != nil {
		return ClientStreamingMessageV2{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

// ClientStreamingV2_Callback_sendMessage_Params_List is a list of ClientStreamingV2_Callback_sendMessage_Params.
type ClientStreamingV2_Callback_sendMessage_Params_List = capnp.StructList[ClientStreamingV2_Callback_sendMessage_Params]

// NewClientStreamingV2_Callback_sendMessage_Params creates a new list of ClientStreamingV2_Callback_sendMessage_Params.
func NewClientStreamingV2_Callback_sendMessage_Params_List(s *capnp.Segment, sz int32) (ClientStreamingV2_Callback_sendMessage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[ClientStreamingV2_Callback_sendMessage_Params](l), err
}

// ClientStreamingV2_Callback_sendMessage_Params_Future is a wrapper for a ClientStreamingV2_Callback_sendMessage_Params promised by a client call.
type ClientStreamingV2_Callback_sendMessage_Params_Future struct{ *capnp.Future }

func (f ClientStreamingV2_Callback_sendMessage_Params_Future) Struct() (ClientStreamingV2_Callback_sendMessage_Params, error) {
	p, err := f.Future.Ptr()
	return ClientStreamingV2_Callback_sendMessage_Params(p.Struct()), err
}
func (p ClientStreamingV2_Callback_sendMessage_Params_Future) ClientStreamingMsg() ClientStreamingMessageV2_Future {
	return ClientStreamingMessageV2_Future{Future: p.Future.Field(0, nil)}
}

type ClientStreamingV2_Callback_sendMessage_Results capnp.Struct

// ClientStreamingV2_Callback_sendMessage_Results_TypeID is the unique identifier for the type ClientStreamingV2_Callback_sendMessage_Results.
const ClientStreamingV2_Callback_sendMessage_Results_TypeID = 0xb1cb1ea514047f2b

func NewClientStreamingV2_Callback_sendMessage_Results(s *capnp.Segment) (ClientStreamingV2_Callback_sendMessage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ClientStreamingV2_Callback_sendMessage_Results(st), err
}

func NewRootClientStreamingV2_Callback_sendMessage_Results(s *capnp.Segment) (ClientStreamingV2_Callback_sendMessage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ClientStreamingV2_Callback_sendMessage_Results(st), err
}

func ReadRootClientStreamingV2_Callback_sendMessage_Results(msg *capnp.Message) (ClientStreamingV2_Callback_sendMessage_Results, error) {
	root, err := msg.Root()
	return ClientStreamingV2_Callback_sendMessage_Results(root.Struct()), err
}

func (s ClientStreamingV2_Callback_sendMessage_Results) String() string {
	str, _ := text.Marshal(0xb1cb1ea514047f2b, capnp.Struct(s))
	return str
}

func (s ClientStreamingV2_Callback_sendMessage_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ClientStreamingV2_Callback_sendMessage_Results) DecodeFromPtr(p capnp.Ptr) ClientStreamingV2_Callback_sendMessage_Results {
	return ClientStreamingV2_Callback_sendMessage_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ClientStreamingV2_Callback_sendMessage_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ClientStreamingV2_Callback_sendMessage_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ClientStreamingV2_Callback_sendMessage_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ClientStreamingV2_Callback_sendMessage_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// ClientStreamingV2_Callback_sendMessage_Results_List is a list of ClientStreamingV2_Callback_sendMessage_Results.
type ClientStreamingV2_Callback_sendMessage_Results_List = capnp.StructList[ClientStreamingV2_Callback_sendMessage_Results]

// NewClientStreamingV2_Callback_sendMessage_Results creates a new list of ClientStreamingV2_Callback_sendMessage_Results.
func NewClientStreamingV2_Callback_sendMessage_Results_List(s *capnp.Segment, sz int32) (ClientStreamingV2_Callback_sendMessage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[ClientStreamingV2_Callback_sendMessage_Results](l), err
}

// ClientStreamingV2_Callback_sendMessage_Results_Future is a wrapper for a ClientStreamingV2_Callback_sendMessage_Results promised by a client call.
type ClientStreamingV2_Callback_sendMessage_Results_Future struct{ *capnp.Future }

func (f ClientStreamingV2_Callback_sendMessage_Results_Future) Struct() (ClientStreamingV2_Callback_sendMessage_Results, error) {
	p, err := f.Future.Ptr()
	return ClientStreamingV2_Callback_sendMessage_Results(p.Struct()), err
}

type ClientStreamingV2_doStreaming_Params capnp.Struct

// ClientStreamingV2_doStreaming_Params_TypeID is the unique identifier for the type ClientStreamingV2_doStreaming_Params.
const ClientStreamingV2_doStreaming_Params_TypeID = 0x855454c9aa1a7626

func NewClientStreamingV2_doStreaming_Params(s *capnp.Segment) (ClientStreamingV2_doStreaming_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ClientStreamingV2_doStreaming_Params(st), err
}

func NewRootClientStreamingV2_doStreaming_Params(s *capnp.Segment) (ClientStreamingV2_doStreaming_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ClientStreamingV2_doStreaming_Params(st), err
}

func ReadRootClientStreamingV2_doStreaming_Params(msg *capnp.Message) (ClientStreamingV2_doStreaming_Params, error) {
	root, err := msg.Root()
	return ClientStreamingV2_doStreaming_Params(root.Struct()), err
}

func (s ClientStreamingV2_doStreaming_Params) String() string {
	str, _ := text.Marshal(0x855454c9aa1a7626, capnp.Struct(s))
	return str
}

func (s ClientStreamingV2_doStreaming_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ClientStreamingV2_doStreaming_Params) DecodeFromPtr(p capnp.Ptr) ClientStreamingV2_doStreaming_Params {
	return ClientStreamingV2_doStreaming_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ClientStreamingV2_doStreaming_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ClientStreamingV2_doStreaming_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ClientStreamingV2_doStreaming_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ClientStreamingV2_doStreaming_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// ClientStreamingV2_doStreaming_Params_List is a list of ClientStreamingV2_doStreaming_Params.
type ClientStreamingV2_doStreaming_Params_List = capnp.StructList[ClientStreamingV2_doStreaming_Params]

// NewClientStreamingV2_doStreaming_Params creates a new list of ClientStreamingV2_doStreaming_Params.
func NewClientStreamingV2_doStreaming_Params_List(s *capnp.Segment, sz int32) (ClientStreamingV2_doStreaming_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[ClientStreamingV2_doStreaming_Params](l), err
}

// ClientStreamingV2_doStreaming_Params_Future is a wrapper for a ClientStreamingV2_doStreaming_Params promised by a client call.
type ClientStreamingV2_doStreaming_Params_Future struct{ *capnp.Future }

func (f ClientStreamingV2_doStreaming_Params_Future) Struct() (ClientStreamingV2_doStreaming_Params, error) {
	p, err := f.Future.Ptr()
	return ClientStreamingV2_doStreaming_Params(p.Struct()), err
}

type ClientStreamingV2_doStreaming_Results capnp.Struct

// ClientStreamingV2_doStreaming_Results_TypeID is the unique identifier for the type ClientStreamingV2_doStreaming_Results.
const ClientStreamingV2_doStreaming_Results_TypeID = 0x95be816cb8e28cb0

func NewClientStreamingV2_doStreaming_Results(s *capnp.Segment) (ClientStreamingV2_doStreaming_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingV2_doStreaming_Results(st), err
}

func NewRootClientStreamingV2_doStreaming_Results(s *capnp.Segment) (ClientStreamingV2_doStreaming_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingV2_doStreaming_Results(st), err
}

func ReadRootClientStreamingV2_doStreaming_Results(msg *capnp.Message) (ClientStreamingV2_doStreaming_Results, error) {
	root, err := msg.Root()
	return ClientStreamingV2_doStreaming_Results(root.Struct()), err
}

func (s ClientStreamingV2_doStreaming_Results) String() string {
	str, _ := text.Marshal(0x95be816cb8e28cb0, capnp.Struct(s))
	return str
}

func (s ClientStreamingV2_doStreaming_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ClientStreamingV2_doStreaming_Results) DecodeFromPtr(p capnp.Ptr) ClientStreamingV2_doStreaming_Results {
	return ClientStreamingV2_doStreaming_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ClientStreamingV2_doStreaming_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ClientStreamingV2_doStreaming_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ClientStreamingV2_doStreaming_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ClientStreamingV2_doStreaming_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s ClientStreamingV2_doStreaming_Results) Callback() ClientStreamingV2_Callback {
	p, _ := capnp.Struct(s).Ptr(0)
	return ClientStreamingV2_Callback(p.Interface().Client())
}

func (s ClientStreamingV2_doStreaming_Results) HasCallback() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s ClientStreamingV2_doStreaming_Results) SetCallback(v ClientStreamingV2_Callback) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// ClientStreamingV2_doStreaming_Results_List is a list of ClientStreamingV2_doStreaming_Results.
type ClientStreamingV2_doStreaming_Results_List = capnp.StructList[ClientStreamingV2_doStreaming_Results]

// NewClientStreamingV2_doStreaming_Results creates a new list of ClientStreamingV2_doStreaming_Results.
func NewClientStreamingV2_doStreaming_Results_List(s *capnp.Segment, sz int32) (ClientStreamingV2_doStreaming_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[ClientStreamingV2_doStreaming_Results](l), err
}

// ClientStreamingV2_doStreaming_Results_Future is a wrapper for a ClientStreamingV2_doStreaming_Results promised by a client call.
type ClientStreamingV2_doStreaming_Results_Future struct{ *capnp.Future }

func (f ClientStreamingV2_doStreaming_Results_Future) Struct() (ClientStreamingV2_doStreaming_Results, error) {
	p, err := f.Future.Ptr()
	return ClientStreamingV2_doStreaming_Results(p.Struct()), err
}
func (p ClientStreamingV2_doStreaming_Results_Future) Callback() ClientStreamingV2_Callback {
	return ClientStreamingV2_Callback(p.Future.Field(0, nil).Client())
}

type ClientStreamingMessageV2 capnp.Struct

// ClientStreamingMessageV2_TypeID is the unique identifier for the type ClientStreamingMessageV2.
const ClientStreamingMessageV2_TypeID = 0xc68b54a0995dcce0

func NewClientStreamingMessageV2(s *capnp.Segment) (ClientStreamingMessageV2, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingMessageV2(st), err
}

func NewRootClientStreamingMessageV2(s *capnp.Segment) (ClientStreamingMessageV2, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ClientStreamingMessageV2(st), err
}

func ReadRootClientStreamingMessageV2(msg *capnp.Message) (ClientStreamingMessageV2, error) {
	root, err := msg.Root()
	return ClientStreamingMessageV2(root.Struct()), err
}

func (s ClientStreamingMessageV2) String() string {
	str, _ := text.Marshal(0xc68b54a0995dcce0, capnp.Struct(s))
	return str
}

func (s ClientStreamingMessageV2) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ClientStreamingMessageV2) DecodeFromPtr(p capnp.Ptr) ClientStreamingMessageV2 {
	return ClientStreamingMessageV2(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ClientStreamingMessageV2) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ClientStreamingMessageV2) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ClientStreamingMessageV2) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ClientStreamingMessageV2) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s ClientStreamingMessageV2) Chunk() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s ClientStreamingMessageV2) HasChunk() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s ClientStreamingMessageV2) SetChunk(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

// ClientStreamingMessageV2_List is a list of ClientStreamingMessageV2.
type ClientStreamingMessageV2_List = capnp.StructList[ClientStreamingMessageV2]

// NewClientStreamingMessageV2 creates a new list of ClientStreamingMessageV2.
func NewClientStreamingMessageV2_List(s *capnp.Segment, sz int32) (ClientStreamingMessageV2_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[ClientStreamingMessageV2](l), err
}

// ClientStreamingMessageV2_Future is a wrapper for a ClientStreamingMessageV2 promised by a client call.
type ClientStreamingMessageV2_Future struct{ *capnp.Future }

func (f ClientStreamingMessageV2_Future) Struct() (ClientStreamingMessageV2, error) {
	p, err := f.Future.Ptr()
	return ClientStreamingMessageV2(p.Struct()), err
}
