package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveCodeQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActiveCodeQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveCodeQueryLogic {
	return &ActiveCodeQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActiveCodeQueryLogic) ActiveCodeQuery(in *openned8.ActiveCodeListReq) (*openned8.ActiveCodeListInfo, error) {
	// todo: add your logic here and delete this line

	return &openned8.ActiveCodeListInfo{}, nil
}
