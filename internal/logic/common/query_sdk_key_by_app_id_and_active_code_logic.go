package common

import (
	"context"
	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySdkKeyByAppIdAndActiveCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySdkKeyByAppIdAndActiveCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySdkKeyByAppIdAndActiveCodeLogic {
	return &QuerySdkKeyByAppIdAndActiveCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySdkKeyByAppIdAndActiveCodeLogic) QuerySdkKeyByAppIdAndActiveCode(in *openned8.QuerySdkKeyByAppIdAndActiveCodeReq) (*openned8.BaseString, error) {
	data, err := l.svcCtx.DB.ActiveCodeInfo.Query().Where(activecodeinfo.AppIDEQ(in.AppId), activecodeinfo.ActiveKey(in.ActiveCode)).QueryAppSdk().First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &openned8.BaseString{Data: data.SdkKey}, nil
}
