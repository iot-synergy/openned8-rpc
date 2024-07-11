package common

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAppSdkByActiveCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAppSdkByActiveCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAppSdkByActiveCodeLogic {
	return &QueryAppSdkByActiveCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAppSdkByActiveCodeLogic) QueryAppSdkByActiveCode(in *openned8.QuerySdkKeyByAppIdAndActiveCodeReq) (*openned8.AppSdkInfo, error) {
	data, err := l.svcCtx.DB.ActiveCodeInfo.Query().Where(activecodeinfo.AppKey(in.AppId), activecodeinfo.ActiveKey(in.ActiveCode)).QueryAppSdk().First(l.ctx)
	if err != nil {
		return nil, err
	}
	// return &openned8.BaseString{Data: data.SdkKey}, nil
	return &openned8.AppSdkInfo{
		Id:          data.ID.String(),
		AppId:       data.App,
		AppName:     "",
		SdkId:       data.Sdk,
		SdkName:     "",
		SdkKey:      data.SdkKey,
		SdkAvatar:   "",
		SdkDesc:     "",
		DownloadUrl: "",
	}, nil
}
