package admin

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &openned8.SdkUsage{}, nil
}
