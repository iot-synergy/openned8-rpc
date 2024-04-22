// Code generated by goctl. DO NOT EDIT.
// Source: openned8.proto

package server

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/logic/admin"
	"github.com/iot-synergy/openned8-rpc/internal/logic/developer"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"
)

type Openned8Server struct {
	svcCtx *svc.ServiceContext
	openned8.UnimplementedOpenned8Server
}

func NewOpenned8Server(svcCtx *svc.ServiceContext) *Openned8Server {
	return &Openned8Server{
		svcCtx: svcCtx,
	}
}

func (s *Openned8Server) AppCreate(ctx context.Context, in *openned8.AppInfoCreateReq) (*openned8.AppInfo, error) {
	l := developer.NewAppCreateLogic(ctx, s.svcCtx)
	return l.AppCreate(in)
}

func (s *Openned8Server) AppUpdate(ctx context.Context, in *openned8.AppInfoUpdateReq) (*openned8.AppInfo, error) {
	l := developer.NewAppUpdateLogic(ctx, s.svcCtx)
	return l.AppUpdate(in)
}

func (s *Openned8Server) AppDelete(ctx context.Context, in *openned8.IdString) (*openned8.BeanMsg, error) {
	l := developer.NewAppDeleteLogic(ctx, s.svcCtx)
	return l.AppDelete(in)
}

func (s *Openned8Server) AppQuery(ctx context.Context, in *openned8.AppListReq) (*openned8.ApplistInfo, error) {
	l := developer.NewAppQueryLogic(ctx, s.svcCtx)
	return l.AppQuery(in)
}

func (s *Openned8Server) CategoryQuery(ctx context.Context, in *openned8.Empty) (*openned8.CategorylistResp, error) {
	l := developer.NewCategoryQueryLogic(ctx, s.svcCtx)
	return l.CategoryQuery(in)
}

func (s *Openned8Server) IndustryQuery(ctx context.Context, in *openned8.Empty) (*openned8.IndustrylistResp, error) {
	l := developer.NewIndustryQueryLogic(ctx, s.svcCtx)
	return l.IndustryQuery(in)
}

func (s *Openned8Server) ActiveCodeQuery(ctx context.Context, in *openned8.ActiveCodeListReq) (*openned8.ActiveCodeListInfo, error) {
	l := developer.NewActiveCodeQueryLogic(ctx, s.svcCtx)
	return l.ActiveCodeQuery(in)
}

func (s *Openned8Server) ActiveCodeCreat(ctx context.Context, in *openned8.ActiveCodeCreatReq) (*openned8.ActiveCodeResp, error) {
	l := developer.NewActiveCodeCreatLogic(ctx, s.svcCtx)
	return l.ActiveCodeCreat(in)
}

func (s *Openned8Server) QueryUserSdkUsage(ctx context.Context, in *openned8.UserSdkUsageQueryReq) (*openned8.SdkUsage, error) {
	l := admin.NewQueryUserSdkUsageLogic(ctx, s.svcCtx)
	return l.QueryUserSdkUsage(in)
}

func (s *Openned8Server) UpdateUserSdkUsage(ctx context.Context, in *openned8.UserSdkUsageUpdateReq) (*openned8.SdkUsage, error) {
	l := admin.NewUpdateUserSdkUsageLogic(ctx, s.svcCtx)
	return l.UpdateUserSdkUsage(in)
}
