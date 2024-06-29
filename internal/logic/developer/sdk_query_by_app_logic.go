package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
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

func (l *SdkQueryByAppLogic) SdkQueryByApp(in *openned8.SdkQueryByAppReq) (*openned8.AppSdkListResp, error) {
	where := l.svcCtx.DB.AppInfo.Query().Where(appinfo.UserID(in.UserId), appinfo.ID(in.AppId)).
		QueryAppSdk().WithSdkInfo().WithAppInfo()
	count, err := where.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data, err := where.Order(appsdk.ByCreatedAt()).Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.AppSdkInfo, 0)
	for _, datum := range data {
		result = append(result, &openned8.AppSdkInfo{
			Id:          datum.ID.String(),
			AppId:       datum.Edges.AppInfo.ID,
			AppName:     datum.Edges.AppInfo.AppName,
			SdkId:       datum.Edges.SdkInfo.ID,
			SdkName:     datum.Edges.SdkInfo.Name,
			SdkKey:      datum.SdkKey,
			SdkAvatar:   datum.Edges.SdkInfo.Avatar,
			SdkDesc:     datum.Edges.SdkInfo.Desc,
			DownloadUrl: datum.Edges.SdkInfo.DownloadURL,
		})
	}
	return &openned8.AppSdkListResp{
		Count: int64(count),
		Data:  result,
	}, nil
}
