package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type SdkListQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSdkListQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SdkListQueryLogic {
	return &SdkListQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SdkListQueryLogic) SdkListQuery(in *openned8.SdkListQueryReq) (*openned8.SdkListResp, error) {
	where := l.svcCtx.DB.SdkInfo.Query().Where()
	count, err := where.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data, err := where.Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.SdkInfo, 0)
	for _, datum := range data {
		result = append(result, &openned8.SdkInfo{
			Id:          datum.ID.String(),
			Name:        datum.Name,
			Avatar:      datum.Avatar,
			Desc:        datum.Desc,
			DownloadUrl: datum.DownloadURL,
		})
	}
	return &openned8.SdkListResp{
		Count: int64(count),
		Data:  result,
	}, nil
}
