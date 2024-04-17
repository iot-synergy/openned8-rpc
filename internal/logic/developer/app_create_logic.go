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
	save, err := l.svcCtx.DB.AppInfo.Create().SetUserID(in.UserId).SetAppName(in.AppName).SetSummary(in.Summary).
		SetAppCategory(in.AppCategory).SetUseIndustry(in.UseIndustry).
		SetAppCategoryName(in.AppCategoryName).SetUseIndustryName(in.UseIndustryName).
		SetAppKey(in.AppKey).SetAppSecret(in.AppSecret).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &openned8.AppInfo{
		Id:              save.ID.String(),
		CreatedAt:       save.CreatedAt.UnixMilli(),
		UpdatedAt:       save.UpdatedAt.UnixMilli(),
		UserId:          save.UserID,
		AppName:         save.AppName,
		Summary:         save.Summary,
		AppCategory:     save.AppCategory,
		UseIndustry:     save.UseIndustry,
		AppCategoryName: save.AppCategoryName,
		UseIndustryName: save.UseIndustryName,
		AppKey:          save.AppKey,
		AppSecret:       save.AppSecret,
	}, nil
}
