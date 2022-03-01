// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/micro_one.proto

package micro_one

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

// Api Endpoints for MicroOne service

func NewMicroOneEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MicroOne service

type MicroOneService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (MicroOne_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (MicroOne_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (MicroOne_BidiStreamService, error)
}

type microOneService struct {
	c    client.Client
	name string
}

func NewMicroOneService(name string, c client.Client) MicroOneService {
	return &microOneService{
		c:    c,
		name: name,
	}
}

func (c *microOneService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "MicroOne.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microOneService) ClientStream(ctx context.Context, opts ...client.CallOption) (MicroOne_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "MicroOne.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &microOneServiceClientStream{stream}, nil
}

type MicroOne_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type microOneServiceClientStream struct {
	stream client.Stream
}

func (x *microOneServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *microOneServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *microOneServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microOneServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microOneServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microOneServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *microOneService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (MicroOne_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "MicroOne.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &microOneServiceServerStream{stream}, nil
}

type MicroOne_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type microOneServiceServerStream struct {
	stream client.Stream
}

func (x *microOneServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *microOneServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *microOneServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microOneServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microOneServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microOneServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *microOneService) BidiStream(ctx context.Context, opts ...client.CallOption) (MicroOne_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "MicroOne.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &microOneServiceBidiStream{stream}, nil
}

type MicroOne_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type microOneServiceBidiStream struct {
	stream client.Stream
}

func (x *microOneServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *microOneServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *microOneServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microOneServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microOneServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microOneServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *microOneServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MicroOne service

type MicroOneHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, MicroOne_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, MicroOne_ServerStreamStream) error
	BidiStream(context.Context, MicroOne_BidiStreamStream) error
}

func RegisterMicroOneHandler(s server.Server, hdlr MicroOneHandler, opts ...server.HandlerOption) error {
	type microOne interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type MicroOne struct {
		microOne
	}
	h := &microOneHandler{hdlr}
	return s.Handle(s.NewHandler(&MicroOne{h}, opts...))
}

type microOneHandler struct {
	MicroOneHandler
}

func (h *microOneHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.MicroOneHandler.Call(ctx, in, out)
}

func (h *microOneHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.MicroOneHandler.ClientStream(ctx, &microOneClientStreamStream{stream})
}

type MicroOne_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type microOneClientStreamStream struct {
	stream server.Stream
}

func (x *microOneClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *microOneClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microOneClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microOneClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microOneClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *microOneHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MicroOneHandler.ServerStream(ctx, m, &microOneServerStreamStream{stream})
}

type MicroOne_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type microOneServerStreamStream struct {
	stream server.Stream
}

func (x *microOneServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *microOneServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microOneServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microOneServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microOneServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *microOneHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.MicroOneHandler.BidiStream(ctx, &microOneBidiStreamStream{stream})
}

type MicroOne_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type microOneBidiStreamStream struct {
	stream server.Stream
}

func (x *microOneBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *microOneBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *microOneBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *microOneBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *microOneBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *microOneBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
