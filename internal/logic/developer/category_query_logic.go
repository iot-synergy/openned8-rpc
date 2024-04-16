package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryQueryLogic {
	return &CategoryQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryQueryLogic) CategoryQuery(in *openned8.Empty) (*openned8.CategorylistResp, error) {
	// todo: add your logic here and delete this line

	return &openned8.CategorylistResp{}, nil
}
