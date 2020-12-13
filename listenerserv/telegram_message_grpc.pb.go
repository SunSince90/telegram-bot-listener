// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package listenerserv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TelegramListenerClient is the client API for TelegramListener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelegramListenerClient interface {
	// Obtains the feature at a given position.
	SendMessage(ctx context.Context, in *TelegramMessage, opts ...grpc.CallOption) (*TelegramResponse, error)
}

type telegramListenerClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegramListenerClient(cc grpc.ClientConnInterface) TelegramListenerClient {
	return &telegramListenerClient{cc}
}

func (c *telegramListenerClient) SendMessage(ctx context.Context, in *TelegramMessage, opts ...grpc.CallOption) (*TelegramResponse, error) {
	out := new(TelegramResponse)
	err := c.cc.Invoke(ctx, "/telegramlistener.TelegramListener/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramListenerServer is the server API for TelegramListener service.
// All implementations must embed UnimplementedTelegramListenerServer
// for forward compatibility
type TelegramListenerServer interface {
	// Obtains the feature at a given position.
	SendMessage(context.Context, *TelegramMessage) (*TelegramResponse, error)
	mustEmbedUnimplementedTelegramListenerServer()
}

// UnimplementedTelegramListenerServer must be embedded to have forward compatible implementations.
type UnimplementedTelegramListenerServer struct {
}

func (UnimplementedTelegramListenerServer) SendMessage(context.Context, *TelegramMessage) (*TelegramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedTelegramListenerServer) mustEmbedUnimplementedTelegramListenerServer() {}

// UnsafeTelegramListenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegramListenerServer will
// result in compilation errors.
type UnsafeTelegramListenerServer interface {
	mustEmbedUnimplementedTelegramListenerServer()
}

func RegisterTelegramListenerServer(s grpc.ServiceRegistrar, srv TelegramListenerServer) {
	s.RegisterService(&_TelegramListener_serviceDesc, srv)
}

func _TelegramListener_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelegramMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramListenerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telegramlistener.TelegramListener/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramListenerServer).SendMessage(ctx, req.(*TelegramMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TelegramListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telegramlistener.TelegramListener",
	HandlerType: (*TelegramListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _TelegramListener_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telegram_message.proto",
}