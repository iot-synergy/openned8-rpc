// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: openned8.proto

package openned8

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

const (
	Openned8_AppCreate_FullMethodName          = "/openned8.openned8/appCreate"
	Openned8_AppUpdate_FullMethodName          = "/openned8.openned8/appUpdate"
	Openned8_AppDelete_FullMethodName          = "/openned8.openned8/appDelete"
	Openned8_AppQuery_FullMethodName           = "/openned8.openned8/appQuery"
	Openned8_CategoryQuery_FullMethodName      = "/openned8.openned8/categoryQuery"
	Openned8_IndustryQuery_FullMethodName      = "/openned8.openned8/industryQuery"
	Openned8_ActiveCodeQuery_FullMethodName    = "/openned8.openned8/activeCodeQuery"
	Openned8_ActiveCodeCreat_FullMethodName    = "/openned8.openned8/activeCodeCreat"
	Openned8_QueryUserSdkUsage_FullMethodName  = "/openned8.openned8/queryUserSdkUsage"
	Openned8_UpdateUserSdkUsage_FullMethodName = "/openned8.openned8/updateUserSdkUsage"
)

// Openned8Client is the client API for Openned8 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Openned8Client interface {
	// group: developer
	AppCreate(ctx context.Context, in *AppInfoCreateReq, opts ...grpc.CallOption) (*AppInfo, error)
	// group: developer
	AppUpdate(ctx context.Context, in *AppInfoUpdateReq, opts ...grpc.CallOption) (*AppInfo, error)
	// group: developer
	AppDelete(ctx context.Context, in *IdString, opts ...grpc.CallOption) (*BeanMsg, error)
	// group: developer
	AppQuery(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*ApplistInfo, error)
	// group: developer
	CategoryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategorylistResp, error)
	// group: developer
	IndustryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IndustrylistResp, error)
	// group: developer
	ActiveCodeQuery(ctx context.Context, in *ActiveCodeListReq, opts ...grpc.CallOption) (*ActiveCodeListInfo, error)
	// group: developer
	ActiveCodeCreat(ctx context.Context, in *ActiveCodeCreatReq, opts ...grpc.CallOption) (*ActiveCodeResp, error)
	// group: admin
	QueryUserSdkUsage(ctx context.Context, in *UserSdkUsageQueryReq, opts ...grpc.CallOption) (*SdkUsage, error)
	// group: admin
	UpdateUserSdkUsage(ctx context.Context, in *UserSdkUsageUpdateReq, opts ...grpc.CallOption) (*SdkUsage, error)
}

type openned8Client struct {
	cc grpc.ClientConnInterface
}

func NewOpenned8Client(cc grpc.ClientConnInterface) Openned8Client {
	return &openned8Client{cc}
}

func (c *openned8Client) AppCreate(ctx context.Context, in *AppInfoCreateReq, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, Openned8_AppCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) AppUpdate(ctx context.Context, in *AppInfoUpdateReq, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, Openned8_AppUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) AppDelete(ctx context.Context, in *IdString, opts ...grpc.CallOption) (*BeanMsg, error) {
	out := new(BeanMsg)
	err := c.cc.Invoke(ctx, Openned8_AppDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) AppQuery(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*ApplistInfo, error) {
	out := new(ApplistInfo)
	err := c.cc.Invoke(ctx, Openned8_AppQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) CategoryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategorylistResp, error) {
	out := new(CategorylistResp)
	err := c.cc.Invoke(ctx, Openned8_CategoryQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) IndustryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IndustrylistResp, error) {
	out := new(IndustrylistResp)
	err := c.cc.Invoke(ctx, Openned8_IndustryQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) ActiveCodeQuery(ctx context.Context, in *ActiveCodeListReq, opts ...grpc.CallOption) (*ActiveCodeListInfo, error) {
	out := new(ActiveCodeListInfo)
	err := c.cc.Invoke(ctx, Openned8_ActiveCodeQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) ActiveCodeCreat(ctx context.Context, in *ActiveCodeCreatReq, opts ...grpc.CallOption) (*ActiveCodeResp, error) {
	out := new(ActiveCodeResp)
	err := c.cc.Invoke(ctx, Openned8_ActiveCodeCreat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) QueryUserSdkUsage(ctx context.Context, in *UserSdkUsageQueryReq, opts ...grpc.CallOption) (*SdkUsage, error) {
	out := new(SdkUsage)
	err := c.cc.Invoke(ctx, Openned8_QueryUserSdkUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openned8Client) UpdateUserSdkUsage(ctx context.Context, in *UserSdkUsageUpdateReq, opts ...grpc.CallOption) (*SdkUsage, error) {
	out := new(SdkUsage)
	err := c.cc.Invoke(ctx, Openned8_UpdateUserSdkUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Openned8Server is the server API for Openned8 service.
// All implementations must embed UnimplementedOpenned8Server
// for forward compatibility
type Openned8Server interface {
	// group: developer
	AppCreate(context.Context, *AppInfoCreateReq) (*AppInfo, error)
	// group: developer
	AppUpdate(context.Context, *AppInfoUpdateReq) (*AppInfo, error)
	// group: developer
	AppDelete(context.Context, *IdString) (*BeanMsg, error)
	// group: developer
	AppQuery(context.Context, *AppListReq) (*ApplistInfo, error)
	// group: developer
	CategoryQuery(context.Context, *Empty) (*CategorylistResp, error)
	// group: developer
	IndustryQuery(context.Context, *Empty) (*IndustrylistResp, error)
	// group: developer
	ActiveCodeQuery(context.Context, *ActiveCodeListReq) (*ActiveCodeListInfo, error)
	// group: developer
	ActiveCodeCreat(context.Context, *ActiveCodeCreatReq) (*ActiveCodeResp, error)
	// group: admin
	QueryUserSdkUsage(context.Context, *UserSdkUsageQueryReq) (*SdkUsage, error)
	// group: admin
	UpdateUserSdkUsage(context.Context, *UserSdkUsageUpdateReq) (*SdkUsage, error)
	mustEmbedUnimplementedOpenned8Server()
}

// UnimplementedOpenned8Server must be embedded to have forward compatible implementations.
type UnimplementedOpenned8Server struct {
}

func (UnimplementedOpenned8Server) AppCreate(context.Context, *AppInfoCreateReq) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppCreate not implemented")
}
func (UnimplementedOpenned8Server) AppUpdate(context.Context, *AppInfoUpdateReq) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdate not implemented")
}
func (UnimplementedOpenned8Server) AppDelete(context.Context, *IdString) (*BeanMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDelete not implemented")
}
func (UnimplementedOpenned8Server) AppQuery(context.Context, *AppListReq) (*ApplistInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppQuery not implemented")
}
func (UnimplementedOpenned8Server) CategoryQuery(context.Context, *Empty) (*CategorylistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryQuery not implemented")
}
func (UnimplementedOpenned8Server) IndustryQuery(context.Context, *Empty) (*IndustrylistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndustryQuery not implemented")
}
func (UnimplementedOpenned8Server) ActiveCodeQuery(context.Context, *ActiveCodeListReq) (*ActiveCodeListInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveCodeQuery not implemented")
}
func (UnimplementedOpenned8Server) ActiveCodeCreat(context.Context, *ActiveCodeCreatReq) (*ActiveCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveCodeCreat not implemented")
}
func (UnimplementedOpenned8Server) QueryUserSdkUsage(context.Context, *UserSdkUsageQueryReq) (*SdkUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserSdkUsage not implemented")
}
func (UnimplementedOpenned8Server) UpdateUserSdkUsage(context.Context, *UserSdkUsageUpdateReq) (*SdkUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSdkUsage not implemented")
}
func (UnimplementedOpenned8Server) mustEmbedUnimplementedOpenned8Server() {}

// UnsafeOpenned8Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Openned8Server will
// result in compilation errors.
type UnsafeOpenned8Server interface {
	mustEmbedUnimplementedOpenned8Server()
}

func RegisterOpenned8Server(s grpc.ServiceRegistrar, srv Openned8Server) {
	s.RegisterService(&Openned8_ServiceDesc, srv)
}

func _Openned8_AppCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfoCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).AppCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_AppCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).AppCreate(ctx, req.(*AppInfoCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_AppUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfoUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).AppUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_AppUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).AppUpdate(ctx, req.(*AppInfoUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_AppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).AppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_AppDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).AppDelete(ctx, req.(*IdString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_AppQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).AppQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_AppQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).AppQuery(ctx, req.(*AppListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_CategoryQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).CategoryQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_CategoryQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).CategoryQuery(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_IndustryQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).IndustryQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_IndustryQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).IndustryQuery(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_ActiveCodeQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveCodeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).ActiveCodeQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_ActiveCodeQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).ActiveCodeQuery(ctx, req.(*ActiveCodeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_ActiveCodeCreat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveCodeCreatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).ActiveCodeCreat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_ActiveCodeCreat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).ActiveCodeCreat(ctx, req.(*ActiveCodeCreatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_QueryUserSdkUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSdkUsageQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).QueryUserSdkUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_QueryUserSdkUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).QueryUserSdkUsage(ctx, req.(*UserSdkUsageQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openned8_UpdateUserSdkUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSdkUsageUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Openned8Server).UpdateUserSdkUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Openned8_UpdateUserSdkUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Openned8Server).UpdateUserSdkUsage(ctx, req.(*UserSdkUsageUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Openned8_ServiceDesc is the grpc.ServiceDesc for Openned8 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Openned8_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openned8.openned8",
	HandlerType: (*Openned8Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "appCreate",
			Handler:    _Openned8_AppCreate_Handler,
		},
		{
			MethodName: "appUpdate",
			Handler:    _Openned8_AppUpdate_Handler,
		},
		{
			MethodName: "appDelete",
			Handler:    _Openned8_AppDelete_Handler,
		},
		{
			MethodName: "appQuery",
			Handler:    _Openned8_AppQuery_Handler,
		},
		{
			MethodName: "categoryQuery",
			Handler:    _Openned8_CategoryQuery_Handler,
		},
		{
			MethodName: "industryQuery",
			Handler:    _Openned8_IndustryQuery_Handler,
		},
		{
			MethodName: "activeCodeQuery",
			Handler:    _Openned8_ActiveCodeQuery_Handler,
		},
		{
			MethodName: "activeCodeCreat",
			Handler:    _Openned8_ActiveCodeCreat_Handler,
		},
		{
			MethodName: "queryUserSdkUsage",
			Handler:    _Openned8_QueryUserSdkUsage_Handler,
		},
		{
			MethodName: "updateUserSdkUsage",
			Handler:    _Openned8_UpdateUserSdkUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openned8.proto",
}
