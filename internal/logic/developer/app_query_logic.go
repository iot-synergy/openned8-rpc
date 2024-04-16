package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppQueryLogic {
	return &AppQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppQueryLogic) AppQuery(in *openned8.AppListReq) (*openned8.ApplistInfo, error) {
	// todo: add your logic here and delete this line

	return &openned8.ApplistInfo{}, nil
}
