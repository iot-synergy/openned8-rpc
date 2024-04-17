package developer

import (
	"context"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
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
	where := l.svcCtx.DB.AppInfo.Query().Where(
		appinfo.AppNameContains(in.AppName),
		appinfo.AppCategoryNameContains(in.AppCategoryName),
		appinfo.UseIndustryNameContains(in.UseIndustryName),
	)
	count, err := where.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data, err := where.Offset(int((in.PageInfo.Page - 1) * in.PageInfo.PageSize)).Limit(int(in.PageInfo.PageSize)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.AppInfo, 0)
	for _, datum := range data {
		result = append(result, &openned8.AppInfo{
			Id:              datum.ID.String(),
			CreatedAt:       datum.CreatedAt.UnixMilli(),
			UpdatedAt:       datum.UpdatedAt.UnixMilli(),
			UserId:          datum.UserID,
			AppName:         datum.AppName,
			Summary:         datum.Summary,
			AppCategory:     datum.AppCategory,
			UseIndustry:     datum.UseIndustry,
			AppCategoryName: datum.AppCategoryName,
			UseIndustryName: datum.UseIndustryName,
			AppKey:          datum.AppKey,
			AppSecret:       datum.AppSecret,
		})
	}
	return &openned8.ApplistInfo{
		Total: uint64(count),
		Data:  result,
	}, nil
}
