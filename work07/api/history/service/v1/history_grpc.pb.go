// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/history/service/v1/history.proto

package v1

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

// HistoryClient is the client API for History service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryClient interface {
	SaveHistory(ctx context.Context, in *SaveHistoryRequest, opts ...grpc.CallOption) (*SaveHistoryReply, error)
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryReply, error)
}

type historyClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryClient(cc grpc.ClientConnInterface) HistoryClient {
	return &historyClient{cc}
}

func (c *historyClient) SaveHistory(ctx context.Context, in *SaveHistoryRequest, opts ...grpc.CallOption) (*SaveHistoryReply, error) {
	out := new(SaveHistoryReply)
	err := c.cc.Invoke(ctx, "/api.history.service.v1.History/SaveHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryReply, error) {
	out := new(GetHistoryReply)
	err := c.cc.Invoke(ctx, "/api.history.service.v1.History/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServer is the server API for History service.
// All implementations must embed UnimplementedHistoryServer
// for forward compatibility
type HistoryServer interface {
	SaveHistory(context.Context, *SaveHistoryRequest) (*SaveHistoryReply, error)
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryReply, error)
	mustEmbedUnimplementedHistoryServer()
}

// UnimplementedHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryServer struct {
}

func (UnimplementedHistoryServer) SaveHistory(context.Context, *SaveHistoryRequest) (*SaveHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveHistory not implemented")
}
func (UnimplementedHistoryServer) GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedHistoryServer) mustEmbedUnimplementedHistoryServer() {}

// UnsafeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServer will
// result in compilation errors.
type UnsafeHistoryServer interface {
	mustEmbedUnimplementedHistoryServer()
}

func RegisterHistoryServer(s grpc.ServiceRegistrar, srv HistoryServer) {
	s.RegisterService(&History_ServiceDesc, srv)
}

func _History_SaveHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).SaveHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.history.service.v1.History/SaveHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).SaveHistory(ctx, req.(*SaveHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.history.service.v1.History/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).GetHistory(ctx, req.(*GetHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// History_ServiceDesc is the grpc.ServiceDesc for History service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var History_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.history.service.v1.History",
	HandlerType: (*HistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveHistory",
			Handler:    _History_SaveHistory_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _History_GetHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/history/service/v1/history.proto",
}
