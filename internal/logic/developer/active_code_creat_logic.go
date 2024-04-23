package developer

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent"
	"github.com/iot-synergy/openned8-rpc/ent/sdkusage"
	"time"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveCodeCreatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActiveCodeCreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveCodeCreatLogic {
	return &ActiveCodeCreatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActiveCodeCreatLogic) ActiveCodeCreat(in *openned8.ActiveCodeCreatReq) (*openned8.ActiveCodeResp, error) {
	if in.Quantity <= 0 {
		return nil, errors.New("生成激活码的数量必须大于零")
	}
	tx, err := l.svcCtx.DB.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	client := tx.Client()

	code, msg, err, data := creatActiveCode(l.ctx, in, client)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			logx.Error("rollback fail:" + err.Error())
		}
		return nil, err
	} else {
		err := tx.Commit()
		if err != nil {
			logx.Error("commit fail:" + err.Error())
			return nil, err
		}
	}

	return &openned8.ActiveCodeResp{
		Code: code,
		Msg:  msg,
		Data: data,
	}, nil
}

func creatActiveCode(ctx context.Context, in *openned8.ActiveCodeCreatReq, client *ent.Client) (
	int32, string, error, []*openned8.ActiveCodeInfo) {
	//减少可用sdk数量
	first, err := client.SdkUsage.Query().Where(sdkusage.UserIDEQ(in.UserId)).First(ctx)
	if err != nil {
		return 0, "", err, nil
	}
	if first.All < first.Used+int64(in.Quantity) {
		return -2, "可用sdk数量不足", nil, nil
	}
	first.Used += int64(in.Quantity)
	err = client.SdkUsage.UpdateOne(first).Exec(ctx)
	if err != nil {
		return 0, "", err, nil
	}

	data := make([]*openned8.ActiveCodeInfo, 0)
	//创建激活码
	for i := 0; i < int(in.GetQuantity()); i++ {
		save, err := client.ActiveCodeInfo.Create().
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			SetActiveKey(uuid.UUID{}.String()).
			SetUserID(in.UserId).
			SetVersion("v0.0.0").
			SetStartDate(time.Now()).
			SetExpireDate(time.Now().AddDate(0, 0, 3)).
			SetActiveFile("").
			SetStatus(1).
			Save(ctx)
		if err != nil {
			return 0, "", err, nil
		}
		data = append(data, &openned8.ActiveCodeInfo{
			Id:             save.ID.String(),
			CreatedAt:      save.CreatedAt.UnixMilli(),
			UpdatedAt:      save.UpdatedAt.UnixMilli(),
			ActiveKey:      save.ActiveKey,
			UserId:         save.UserID,
			AppId:          save.AppID,
			ActiveIP:       save.ActiveIP,
			DeviceSN:       save.DeviceSn,
			DeviceMac:      save.DeviceMAC,
			DeviceIdentity: save.DeviceIdentity,
			ActiveDate:     save.ActiveDate.UnixMilli(),
			ActiveType:     save.ActiveType,
			ActiveFile:     save.ActiveFile,
			Version:        save.Version,
			StartDate:      save.StartDate.UnixMilli(),
			ExpireDate:     save.ExpireDate.UnixMilli(),
			Status:         int64(save.Status),
		})
	}
	return 0, "成功", nil, data
}
