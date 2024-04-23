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
	all, err := l.svcCtx.DB.CategoryInfo.Query().All(l.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.CategoryInfo, 0)
	for _, info := range all {
		result = append(result, &openned8.CategoryInfo{
			Id:   int64(info.ID),
			Name: info.Name,
		})
	}
	return &openned8.CategorylistResp{
		Data: result,
	}, nil
}
