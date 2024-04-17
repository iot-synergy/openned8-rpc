package developer

import (
	"context"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
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

func (l *AppUpdateLogic) AppUpdate(in *openned8.AppInfo) (*openned8.AppInfo, error) {
	fromString, err := uuid.FromString(in.UserId)
	if err != nil {
		return nil, err
	}
	save, err := l.svcCtx.DB.AppInfo.UpdateOne(&ent.AppInfo{
		ID:              uuid.UUID{},
		UpdatedAt:       time.Now(),
		UserID:          in.UserId,
		AppName:         in.AppName,
		Summary:         in.Summary,
		AppCategory:     in.AppCategory,
		UseIndustry:     in.UseIndustry,
		AppCategoryName: in.AppCategoryName,
		UseIndustryName: in.UseIndustryName,
		AppKey:          in.AppKey,
		AppSecret:       in.AppSecret,
	}).Where(appinfo.IDEQ(fromString)).Save(l.ctx)
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
