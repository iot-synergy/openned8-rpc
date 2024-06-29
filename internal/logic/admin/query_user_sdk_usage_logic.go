package admin

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/ent"
	"github.com/iot-synergy/openned8-rpc/ent/sdkusage"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserSdkUsageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserSdkUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserSdkUsageLogic {
	return &QueryUserSdkUsageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserSdkUsageLogic) QueryUserSdkUsage(in *openned8.UserSdkUsageQueryReq) (*openned8.SdkUsage, error) {
	tx, err := l.svcCtx.DB.Tx(l.ctx)
	client := tx.Client()
	if err != nil {
		return nil, err
	}
	data, err := QueryUserSdk(client, in, l.ctx)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			l.Logger.Error("rollback fail:" + err.Error())
		}
		return nil, err
	} else {
		err := tx.Commit()
		if err != nil {
			l.Logger.Error("commit fail:" + err.Error())
			return nil, err
		}
	}
	return &openned8.SdkUsage{
		UserId: data.UserID,
		All:    data.All,
		Used:   data.Used,
	}, nil
}

func QueryUserSdk(client *ent.Client, in *openned8.UserSdkUsageQueryReq, ctx context.Context) (*ent.SdkUsage, error) {
	count, err := client.SdkUsage.Query().Where(sdkusage.UserIDEQ(in.UserId)).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		save, err := client.SdkUsage.Create().
			SetUserID(in.UserId).
			SetAll(100).
			SetUsed(0).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &ent.SdkUsage{
			ID:        save.ID,
			CreatedAt: save.CreatedAt,
			UpdatedAt: save.UpdatedAt,
			UserID:    save.UserID,
			All:       save.All,
			Used:      save.Used,
		}, nil
	}
	first, err := client.SdkUsage.Query().Where(sdkusage.UserIDEQ(in.UserId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &ent.SdkUsage{
		ID:        first.ID,
		CreatedAt: first.CreatedAt,
		UpdatedAt: first.UpdatedAt,
		UserID:    first.UserID,
		All:       first.All,
		Used:      first.Used,
	}, nil
}
