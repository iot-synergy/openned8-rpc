package admin

import (
	"context"
	"github.com/iot-synergy/openned8-rpc/ent/sdkusage"

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
	data, err := l.svcCtx.DB.SdkUsage.Query().Where(sdkusage.UserIDEQ(in.UserId)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &openned8.SdkUsage{
		UserId: data.UserID,
		All:    data.All,
		Used:   data.Used,
	}, nil
}
