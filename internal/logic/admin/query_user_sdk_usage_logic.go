package admin

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserSdkUsageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserSdkUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserSdkUsageLogic {
	return &QueryUserSdkUsageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserSdkUsageLogic) QueryUserSdkUsage(in *openned8.UserSdkUsageQueryReq) (*openned8.SdkUsage, error) {
	// todo: add your logic here and delete this line

	return &openned8.SdkUsage{}, nil
}
