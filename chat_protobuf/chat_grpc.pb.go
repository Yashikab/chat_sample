// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: chat.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloGrpcClient is the client API for HelloGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloGrpcClient interface {
	GreetServer(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetMessage, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (HelloGrpc_SendMessageClient, error)
	GetMessage(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (HelloGrpc_GetMessageClient, error)
}

type helloGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloGrpcClient(cc grpc.ClientConnInterface) HelloGrpcClient {
	return &helloGrpcClient{cc}
}

func (c *helloGrpcClient) GreetServer(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetMessage, error) {
	out := new(GreetMessage)
	err := c.cc.Invoke(ctx, "/chat.HelloGrpc/GreetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloGrpcClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (HelloGrpc_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloGrpc_ServiceDesc.Streams[0], "/chat.HelloGrpc/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloGrpcSendMessageClient{stream}
	return x, nil
}

type HelloGrpc_SendMessageClient interface {
	Send(*SendRequest) error
	CloseAndRecv() (*SendResult, error)
	grpc.ClientStream
}

type helloGrpcSendMessageClient struct {
	grpc.ClientStream
}

func (x *helloGrpcSendMessageClient) Send(m *SendRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloGrpcSendMessageClient) CloseAndRecv() (*SendResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloGrpcClient) GetMessage(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (HelloGrpc_GetMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloGrpc_ServiceDesc.Streams[1], "/chat.HelloGrpc/GetMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloGrpcGetMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloGrpc_GetMessageClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type helloGrpcGetMessageClient struct {
	grpc.ClientStream
}

func (x *helloGrpcGetMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloGrpcServer is the server API for HelloGrpc service.
// All implementations must embed UnimplementedHelloGrpcServer
// for forward compatibility
type HelloGrpcServer interface {
	GreetServer(context.Context, *GreetRequest) (*GreetMessage, error)
	SendMessage(HelloGrpc_SendMessageServer) error
	GetMessage(*MessagesRequest, HelloGrpc_GetMessageServer) error
	mustEmbedUnimplementedHelloGrpcServer()
}

// UnimplementedHelloGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedHelloGrpcServer struct {
}

func (UnimplementedHelloGrpcServer) GreetServer(context.Context, *GreetRequest) (*GreetMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetServer not implemented")
}
func (UnimplementedHelloGrpcServer) SendMessage(HelloGrpc_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedHelloGrpcServer) GetMessage(*MessagesRequest, HelloGrpc_GetMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedHelloGrpcServer) mustEmbedUnimplementedHelloGrpcServer() {}

// UnsafeHelloGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloGrpcServer will
// result in compilation errors.
type UnsafeHelloGrpcServer interface {
	mustEmbedUnimplementedHelloGrpcServer()
}

func RegisterHelloGrpcServer(s grpc.ServiceRegistrar, srv HelloGrpcServer) {
	s.RegisterService(&HelloGrpc_ServiceDesc, srv)
}

func _HelloGrpc_GreetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloGrpcServer).GreetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.HelloGrpc/GreetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloGrpcServer).GreetServer(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloGrpc_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloGrpcServer).SendMessage(&helloGrpcSendMessageServer{stream})
}

type HelloGrpc_SendMessageServer interface {
	SendAndClose(*SendResult) error
	Recv() (*SendRequest, error)
	grpc.ServerStream
}

type helloGrpcSendMessageServer struct {
	grpc.ServerStream
}

func (x *helloGrpcSendMessageServer) SendAndClose(m *SendResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloGrpcSendMessageServer) Recv() (*SendRequest, error) {
	m := new(SendRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloGrpc_GetMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloGrpcServer).GetMessage(m, &helloGrpcGetMessageServer{stream})
}

type HelloGrpc_GetMessageServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type helloGrpcGetMessageServer struct {
	grpc.ServerStream
}

func (x *helloGrpcGetMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// HelloGrpc_ServiceDesc is the grpc.ServiceDesc for HelloGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.HelloGrpc",
	HandlerType: (*HelloGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetServer",
			Handler:    _HelloGrpc_GreetServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _HelloGrpc_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMessage",
			Handler:       _HelloGrpc_GetMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
