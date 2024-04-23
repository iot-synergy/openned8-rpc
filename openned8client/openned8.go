// Code generated by goctl. DO NOT EDIT.
// Source: openned8.proto

package openned8client

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActiveCodeCreatReq    = openned8.ActiveCodeCreatReq
	ActiveCodeInfo        = openned8.ActiveCodeInfo
	ActiveCodeListInfo    = openned8.ActiveCodeListInfo
	ActiveCodeListReq     = openned8.ActiveCodeListReq
	ActiveCodeResp        = openned8.ActiveCodeResp
	AppInfo               = openned8.AppInfo
	AppInfoCreateReq      = openned8.AppInfoCreateReq
	AppInfoDeleteReq      = openned8.AppInfoDeleteReq
	AppInfoUpdateReq      = openned8.AppInfoUpdateReq
	AppListReq            = openned8.AppListReq
	ApplistInfo           = openned8.ApplistInfo
	BeanMsg               = openned8.BeanMsg
	CategoryInfo          = openned8.CategoryInfo
	CategorylistResp      = openned8.CategorylistResp
	Empty                 = openned8.Empty
	IdString              = openned8.IdString
	IndustryInfo          = openned8.IndustryInfo
	IndustrylistResp      = openned8.IndustrylistResp
	PageInfo              = openned8.PageInfo
	SdkUsage              = openned8.SdkUsage
	UserSdkUsageQueryReq  = openned8.UserSdkUsageQueryReq
	UserSdkUsageUpdateReq = openned8.UserSdkUsageUpdateReq

	Openned8 interface {
		AppCreate(ctx context.Context, in *AppInfoCreateReq, opts ...grpc.CallOption) (*AppInfo, error)
		AppUpdate(ctx context.Context, in *AppInfoUpdateReq, opts ...grpc.CallOption) (*AppInfo, error)
		AppDelete(ctx context.Context, in *AppInfoDeleteReq, opts ...grpc.CallOption) (*BeanMsg, error)
		AppQuery(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*ApplistInfo, error)
		CategoryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategorylistResp, error)
		IndustryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IndustrylistResp, error)
		ActiveCodeQuery(ctx context.Context, in *ActiveCodeListReq, opts ...grpc.CallOption) (*ActiveCodeListInfo, error)
		ActiveCodeCreat(ctx context.Context, in *ActiveCodeCreatReq, opts ...grpc.CallOption) (*ActiveCodeResp, error)
		QueryUserSdkUsage(ctx context.Context, in *UserSdkUsageQueryReq, opts ...grpc.CallOption) (*SdkUsage, error)
		UpdateUserSdkUsage(ctx context.Context, in *UserSdkUsageUpdateReq, opts ...grpc.CallOption) (*SdkUsage, error)
	}

	defaultOpenned8 struct {
		cli zrpc.Client
	}
)

func NewOpenned8(cli zrpc.Client) Openned8 {
	return &defaultOpenned8{
		cli: cli,
	}
}

func (m *defaultOpenned8) AppCreate(ctx context.Context, in *AppInfoCreateReq, opts ...grpc.CallOption) (*AppInfo, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.AppCreate(ctx, in, opts...)
}

func (m *defaultOpenned8) AppUpdate(ctx context.Context, in *AppInfoUpdateReq, opts ...grpc.CallOption) (*AppInfo, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.AppUpdate(ctx, in, opts...)
}

func (m *defaultOpenned8) AppDelete(ctx context.Context, in *AppInfoDeleteReq, opts ...grpc.CallOption) (*BeanMsg, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.AppDelete(ctx, in, opts...)
}

func (m *defaultOpenned8) AppQuery(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*ApplistInfo, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.AppQuery(ctx, in, opts...)
}

func (m *defaultOpenned8) CategoryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategorylistResp, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.CategoryQuery(ctx, in, opts...)
}

func (m *defaultOpenned8) IndustryQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IndustrylistResp, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.IndustryQuery(ctx, in, opts...)
}

func (m *defaultOpenned8) ActiveCodeQuery(ctx context.Context, in *ActiveCodeListReq, opts ...grpc.CallOption) (*ActiveCodeListInfo, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.ActiveCodeQuery(ctx, in, opts...)
}

func (m *defaultOpenned8) ActiveCodeCreat(ctx context.Context, in *ActiveCodeCreatReq, opts ...grpc.CallOption) (*ActiveCodeResp, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.ActiveCodeCreat(ctx, in, opts...)
}

func (m *defaultOpenned8) QueryUserSdkUsage(ctx context.Context, in *UserSdkUsageQueryReq, opts ...grpc.CallOption) (*SdkUsage, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.QueryUserSdkUsage(ctx, in, opts...)
}

func (m *defaultOpenned8) UpdateUserSdkUsage(ctx context.Context, in *UserSdkUsageUpdateReq, opts ...grpc.CallOption) (*SdkUsage, error) {
	client := openned8.NewOpenned8Client(m.cli.Conn())
	return client.UpdateUserSdkUsage(ctx, in, opts...)
}
