package developer

import (
	"context"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type SdkQueryByAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSdkQueryByAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SdkQueryByAppLogic {
	return &SdkQueryByAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SdkQueryByAppLogic) SdkQueryByApp(in *openned8.SdkQueryByAppReq) (*openned8.SdkListResp, error) {
	appId, err := uuid.FromString(in.AppId)
	if err != nil {
		return nil, err
	}
	where := l.svcCtx.DB.AppInfo.Query().Where(appinfo.UserID(in.UserId), appinfo.ID(appId)).
		QueryAppSdk().QuerySdkInfo().Where()
	count, err := where.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data, err := where.Order(sdkinfo.ByDesc()).Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).All(l.ctx)
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
