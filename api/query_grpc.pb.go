// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: query.proto

package api

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

// QueryListenerClient is the client API for QueryListener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryListenerClient interface {
	PostCredit(ctx context.Context, in *CreditInfo, opts ...grpc.CallOption) (*Empty, error)
	PostReserve(ctx context.Context, in *CashFlow, opts ...grpc.CallOption) (*Empty, error)
	PostCancelReserve(ctx context.Context, in *CashFlow, opts ...grpc.CallOption) (*Empty, error)
	PostWriteOff(ctx context.Context, in *CashFlow, opts ...grpc.CallOption) (*Empty, error)
	PostGetBalance(ctx context.Context, in *User, opts ...grpc.CallOption) (*Balance, error)
	PostGetReport(ctx context.Context, in *User, opts ...grpc.CallOption) (*Report, error)
}

type queryListenerClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryListenerClient(cc grpc.ClientConnInterface) QueryListenerClient {
	return &queryListenerClient{cc}
}

func (c *queryListenerClient) PostCredit(ctx context.Context, in *CreditInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/query.QueryListener/PostCredit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryListenerClient) PostReserve(ctx context.Context, in *CashFlow, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/query.QueryListener/PostReserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryListenerClient) PostCancelReserve(ctx context.Context, in *CashFlow, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/query.QueryListener/PostCancelReserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryListenerClient) PostWriteOff(ctx context.Context, in *CashFlow, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/query.QueryListener/PostWriteOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryListenerClient) PostGetBalance(ctx context.Context, in *User, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/query.QueryListener/PostGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryListenerClient) PostGetReport(ctx context.Context, in *User, opts ...grpc.CallOption) (*Report, error) {
	out := new(Report)
	err := c.cc.Invoke(ctx, "/query.QueryListener/PostGetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryListenerServer is the server API for QueryListener service.
// All implementations must embed UnimplementedQueryListenerServer
// for forward compatibility
type QueryListenerServer interface {
	PostCredit(context.Context, *CreditInfo) (*Empty, error)
	PostReserve(context.Context, *CashFlow) (*Empty, error)
	PostCancelReserve(context.Context, *CashFlow) (*Empty, error)
	PostWriteOff(context.Context, *CashFlow) (*Empty, error)
	PostGetBalance(context.Context, *User) (*Balance, error)
	PostGetReport(context.Context, *User) (*Report, error)
	mustEmbedUnimplementedQueryListenerServer()
}

// UnimplementedQueryListenerServer must be embedded to have forward compatible implementations.
type UnimplementedQueryListenerServer struct {
}

func (UnimplementedQueryListenerServer) PostCredit(context.Context, *CreditInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCredit not implemented")
}
func (UnimplementedQueryListenerServer) PostReserve(context.Context, *CashFlow) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostReserve not implemented")
}
func (UnimplementedQueryListenerServer) PostCancelReserve(context.Context, *CashFlow) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCancelReserve not implemented")
}
func (UnimplementedQueryListenerServer) PostWriteOff(context.Context, *CashFlow) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostWriteOff not implemented")
}
func (UnimplementedQueryListenerServer) PostGetBalance(context.Context, *User) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGetBalance not implemented")
}
func (UnimplementedQueryListenerServer) PostGetReport(context.Context, *User) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGetReport not implemented")
}
func (UnimplementedQueryListenerServer) mustEmbedUnimplementedQueryListenerServer() {}

// UnsafeQueryListenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryListenerServer will
// result in compilation errors.
type UnsafeQueryListenerServer interface {
	mustEmbedUnimplementedQueryListenerServer()
}

func RegisterQueryListenerServer(s grpc.ServiceRegistrar, srv QueryListenerServer) {
	s.RegisterService(&QueryListener_ServiceDesc, srv)
}

func _QueryListener_PostCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryListenerServer).PostCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryListener/PostCredit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryListenerServer).PostCredit(ctx, req.(*CreditInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryListener_PostReserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashFlow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryListenerServer).PostReserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryListener/PostReserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryListenerServer).PostReserve(ctx, req.(*CashFlow))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryListener_PostCancelReserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashFlow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryListenerServer).PostCancelReserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryListener/PostCancelReserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryListenerServer).PostCancelReserve(ctx, req.(*CashFlow))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryListener_PostWriteOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashFlow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryListenerServer).PostWriteOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryListener/PostWriteOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryListenerServer).PostWriteOff(ctx, req.(*CashFlow))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryListener_PostGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryListenerServer).PostGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryListener/PostGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryListenerServer).PostGetBalance(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryListener_PostGetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryListenerServer).PostGetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryListener/PostGetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryListenerServer).PostGetReport(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryListener_ServiceDesc is the grpc.ServiceDesc for QueryListener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryListener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "query.QueryListener",
	HandlerType: (*QueryListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostCredit",
			Handler:    _QueryListener_PostCredit_Handler,
		},
		{
			MethodName: "PostReserve",
			Handler:    _QueryListener_PostReserve_Handler,
		},
		{
			MethodName: "PostCancelReserve",
			Handler:    _QueryListener_PostCancelReserve_Handler,
		},
		{
			MethodName: "PostWriteOff",
			Handler:    _QueryListener_PostWriteOff_Handler,
		},
		{
			MethodName: "PostGetBalance",
			Handler:    _QueryListener_PostGetBalance_Handler,
		},
		{
			MethodName: "PostGetReport",
			Handler:    _QueryListener_PostGetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}
