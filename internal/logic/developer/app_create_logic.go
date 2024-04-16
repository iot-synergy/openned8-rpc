package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppCreateLogic {
	return &AppCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppCreateLogic) AppCreate(in *openned8.AppInfo) (*openned8.AppInfo, error) {
	// todo: add your logic here and delete this line

	return &openned8.AppInfo{}, nil
}
