package developer

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/categoryinfo"
	"github.com/iot-synergy/openned8-rpc/ent/industryinfo"
	"time"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppUpdateLogic {
	return &AppUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppUpdateLogic) AppUpdate(in *openned8.AppInfoUpdateReq) (*openned8.AppInfo, error) {
	id, err := uuid.FromString(in.Id)
	if err != nil {
		return nil, err
	}
	appinfo, err := l.svcCtx.DB.AppInfo.Query().Where(appinfo.ID(id)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	if appinfo == nil || appinfo.ID.IsNil() {
		return nil, nil
	}
	update := l.svcCtx.DB.AppInfo.UpdateOneID(id).SetUpdatedAt(time.Now())
	if in.UserId != appinfo.UserID {
		return nil, errors.New("不是当前用户的数据")
	}
	if in.AppName != "" && in.AppName != appinfo.AppName {
		update.SetAppName(in.AppName)
	}
	if in.Summary != "" && in.Summary != appinfo.Summary {
		update.SetSummary(in.Summary)
	}
	if in.AppCategory != appinfo.AppCategory {
		update.SetAppCategory(in.AppCategory)
		category, err := l.svcCtx.DB.CategoryInfo.Query().Where(categoryinfo.IDEQ(uint64(in.AppCategory))).First(l.ctx)
		if err != nil {
			return nil, err
		}
		update.SetAppCategoryName(category.Name)
	}
	if in.UseIndustry != appinfo.UseIndustry {
		update.SetUseIndustry(in.UseIndustry)
		industry, err := l.svcCtx.DB.IndustryInfo.Query().Where(industryinfo.IDEQ(uint64(in.UseIndustry))).First(l.ctx)
		if err != nil {
			return nil, err
		}
		update.SetUseIndustryName(industry.Name)
	}
	save, err := update.Save(l.ctx)
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
