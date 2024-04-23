package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndustryQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIndustryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndustryQueryLogic {
	return &IndustryQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IndustryQueryLogic) IndustryQuery(in *openned8.Empty) (*openned8.IndustrylistResp, error) {
	all, err := l.svcCtx.DB.IndustryInfo.Query().All(l.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.IndustryInfo, 0)
	for _, info := range all {
		result = append(result, &openned8.IndustryInfo{
			Id:   int64(info.ID),
			Name: info.Name,
		})
	}
	return &openned8.IndustrylistResp{
		Data: result,
	}, nil
}
