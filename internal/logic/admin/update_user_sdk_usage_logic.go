package admin

import (
	"context"
	"github.com/iot-synergy/openned8-rpc/ent/sdkusage"
	"time"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserSdkUsageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserSdkUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserSdkUsageLogic {
	return &UpdateUserSdkUsageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserSdkUsageLogic) UpdateUserSdkUsage(in *openned8.UserSdkUsageUpdateReq) (*openned8.SdkUsage, error) {
	_, err := l.svcCtx.DB.SdkUsage.Update().Where(sdkusage.UserIDEQ(in.UserId)).
		SetUpdatedAt(time.Now()).SetAll(in.All).Save(l.ctx)
	if err != nil {
		return nil, err
	}
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
