// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: mysqlctl.proto

package mysqlctl

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

// MysqlCtlClient is the client API for MysqlCtl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MysqlCtlClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	RunMysqlUpgrade(ctx context.Context, in *RunMysqlUpgradeRequest, opts ...grpc.CallOption) (*RunMysqlUpgradeResponse, error)
	ReinitConfig(ctx context.Context, in *ReinitConfigRequest, opts ...grpc.CallOption) (*ReinitConfigResponse, error)
	RefreshConfig(ctx context.Context, in *RefreshConfigRequest, opts ...grpc.CallOption) (*RefreshConfigResponse, error)
}

type mysqlCtlClient struct {
	cc grpc.ClientConnInterface
}

func NewMysqlCtlClient(cc grpc.ClientConnInterface) MysqlCtlClient {
	return &mysqlCtlClient{cc}
}

func (c *mysqlCtlClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/mysqlctl.MysqlCtl/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/mysqlctl.MysqlCtl/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) RunMysqlUpgrade(ctx context.Context, in *RunMysqlUpgradeRequest, opts ...grpc.CallOption) (*RunMysqlUpgradeResponse, error) {
	out := new(RunMysqlUpgradeResponse)
	err := c.cc.Invoke(ctx, "/mysqlctl.MysqlCtl/RunMysqlUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) ReinitConfig(ctx context.Context, in *ReinitConfigRequest, opts ...grpc.CallOption) (*ReinitConfigResponse, error) {
	out := new(ReinitConfigResponse)
	err := c.cc.Invoke(ctx, "/mysqlctl.MysqlCtl/ReinitConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) RefreshConfig(ctx context.Context, in *RefreshConfigRequest, opts ...grpc.CallOption) (*RefreshConfigResponse, error) {
	out := new(RefreshConfigResponse)
	err := c.cc.Invoke(ctx, "/mysqlctl.MysqlCtl/RefreshConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MysqlCtlServer is the server API for MysqlCtl service.
// All implementations must embed UnimplementedMysqlCtlServer
// for forward compatibility
type MysqlCtlServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	RunMysqlUpgrade(context.Context, *RunMysqlUpgradeRequest) (*RunMysqlUpgradeResponse, error)
	ReinitConfig(context.Context, *ReinitConfigRequest) (*ReinitConfigResponse, error)
	RefreshConfig(context.Context, *RefreshConfigRequest) (*RefreshConfigResponse, error)
	mustEmbedUnimplementedMysqlCtlServer()
}

// UnimplementedMysqlCtlServer must be embedded to have forward compatible implementations.
type UnimplementedMysqlCtlServer struct {
}

func (UnimplementedMysqlCtlServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedMysqlCtlServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedMysqlCtlServer) RunMysqlUpgrade(context.Context, *RunMysqlUpgradeRequest) (*RunMysqlUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunMysqlUpgrade not implemented")
}
func (UnimplementedMysqlCtlServer) ReinitConfig(context.Context, *ReinitConfigRequest) (*ReinitConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReinitConfig not implemented")
}
func (UnimplementedMysqlCtlServer) RefreshConfig(context.Context, *RefreshConfigRequest) (*RefreshConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshConfig not implemented")
}
func (UnimplementedMysqlCtlServer) mustEmbedUnimplementedMysqlCtlServer() {}

// UnsafeMysqlCtlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MysqlCtlServer will
// result in compilation errors.
type UnsafeMysqlCtlServer interface {
	mustEmbedUnimplementedMysqlCtlServer()
}

func RegisterMysqlCtlServer(s grpc.ServiceRegistrar, srv MysqlCtlServer) {
	s.RegisterService(&MysqlCtl_ServiceDesc, srv)
}

func _MysqlCtl_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_RunMysqlUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunMysqlUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).RunMysqlUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/RunMysqlUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).RunMysqlUpgrade(ctx, req.(*RunMysqlUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_ReinitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReinitConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).ReinitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/ReinitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).ReinitConfig(ctx, req.(*ReinitConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_RefreshConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).RefreshConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/RefreshConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).RefreshConfig(ctx, req.(*RefreshConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MysqlCtl_ServiceDesc is the grpc.ServiceDesc for MysqlCtl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MysqlCtl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mysqlctl.MysqlCtl",
	HandlerType: (*MysqlCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _MysqlCtl_Start_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _MysqlCtl_Shutdown_Handler,
		},
		{
			MethodName: "RunMysqlUpgrade",
			Handler:    _MysqlCtl_RunMysqlUpgrade_Handler,
		},
		{
			MethodName: "ReinitConfig",
			Handler:    _MysqlCtl_ReinitConfig_Handler,
		},
		{
			MethodName: "RefreshConfig",
			Handler:    _MysqlCtl_RefreshConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mysqlctl.proto",
}
