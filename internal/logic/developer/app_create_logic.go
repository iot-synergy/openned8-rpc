package developer

import (
	"context"
	"github.com/google/uuid"
	"github.com/iot-synergy/openned8-rpc/ent/categoryinfo"
	"github.com/iot-synergy/openned8-rpc/ent/industryinfo"
	"regexp"

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

func (l *AppCreateLogic) AppCreate(in *openned8.AppInfoCreateReq) (*openned8.AppInfo, error) {
	category, err := l.svcCtx.DB.CategoryInfo.Query().Where(categoryinfo.IDEQ(uint64(in.AppCategory))).First(l.ctx)
	if err != nil {
		return nil, err
	}
	industry, err := l.svcCtx.DB.IndustryInfo.Query().Where(industryinfo.IDEQ(uint64(in.UseIndustry))).First(l.ctx)
	if err != nil {
		return nil, err
	}
	appKey, _ := uuid.NewRandom()
	save, err := l.svcCtx.DB.AppInfo.Create().
		SetUserID(in.UserId).
		SetAppName(in.AppName).
		SetSummary(in.Summary).
		SetAppCategory(in.AppCategory).
		SetUseIndustry(in.UseIndustry).
		SetAppCategoryName(category.Name).
		SetUseIndustryName(industry.Name).
		SetAppKey(regexp.MustCompile("-").ReplaceAllLiteralString(appKey.String(), "")).
		Save(l.ctx)
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
