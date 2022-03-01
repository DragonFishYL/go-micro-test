// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/micro_two.proto

package micro_two

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MicroTwo service

func NewMicroTwoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MicroTwo service

type MicroTwoService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (MicroTwo_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (MicroTwo_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (MicroTwo_BidiStreamService, error)
}

type microTwoService struct {
	c    client.Client
	name string
}

func NewMicroTwoService(name string, c client.Client) MicroTwoService {
	return &microTwoService{
		c:    c,
		name: name,
	}
}

func (c *microTwoService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "MicroTwo.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microTwoService) ClientStream(ctx context.Context, opts ...client.CallOption) (MicroTwo_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "MicroTwo.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &microTwoServiceClientStream{stream}, nil
}

type MicroTwo_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type microTwoServiceClientStream struct {
	stream client.Stream
}

func (x *microTwoServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *microTwoServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *microTwoServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microTwoServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microTwoServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microTwoServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *microTwoService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (MicroTwo_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "MicroTwo.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &microTwoServiceServerStream{stream}, nil
}

type MicroTwo_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type microTwoServiceServerStream struct {
	stream client.Stream
}

func (x *microTwoServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *microTwoServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *microTwoServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microTwoServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microTwoServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microTwoServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *microTwoService) BidiStream(ctx context.Context, opts ...client.CallOption) (MicroTwo_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "MicroTwo.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &microTwoServiceBidiStream{stream}, nil
}

type MicroTwo_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type microTwoServiceBidiStream struct {
	stream client.Stream
}

func (x *microTwoServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *microTwoServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *microTwoServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microTwoServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microTwoServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microTwoServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *microTwoServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MicroTwo service

type MicroTwoHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, MicroTwo_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, MicroTwo_ServerStreamStream) error
	BidiStream(context.Context, MicroTwo_BidiStreamStream) error
}

func RegisterMicroTwoHandler(s server.Server, hdlr MicroTwoHandler, opts ...server.HandlerOption) error {
	type microTwo interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type MicroTwo struct {
		microTwo
	}
	h := &microTwoHandler{hdlr}
	return s.Handle(s.NewHandler(&MicroTwo{h}, opts...))
}

type microTwoHandler struct {
	MicroTwoHandler
}

func (h *microTwoHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.MicroTwoHandler.Call(ctx, in, out)
}

func (h *microTwoHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.MicroTwoHandler.ClientStream(ctx, &microTwoClientStreamStream{stream})
}

type MicroTwo_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type microTwoClientStreamStream struct {
	stream server.Stream
}

func (x *microTwoClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *microTwoClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microTwoClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microTwoClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microTwoClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *microTwoHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MicroTwoHandler.ServerStream(ctx, m, &microTwoServerStreamStream{stream})
}

type MicroTwo_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type microTwoServerStreamStream struct {
	stream server.Stream
}

func (x *microTwoServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *microTwoServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microTwoServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microTwoServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microTwoServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *microTwoHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.MicroTwoHandler.BidiStream(ctx, &microTwoBidiStreamStream{stream})
}

type MicroTwo_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type microTwoBidiStreamStream struct {
	stream server.Stream
}

func (x *microTwoBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *microTwoBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microTwoBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microTwoBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microTwoBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *microTwoBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}