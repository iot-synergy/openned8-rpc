package developer

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
	"github.com/iot-synergy/openned8-rpc/ent/sdkusage"
	"github.com/iot-synergy/openned8-rpc/internal/common"
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
		err1 := tx.Rollback()
		if err1 != nil {
			logx.Error("rollback fail:" + err1.Error())
			return nil, err1
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
	err = client.SdkUsage.Update().Where(sdkusage.UserIDEQ(first.UserID)).SetUsed(first.Used + int64(in.Quantity)).Exec(ctx)
	if err != nil {
		return 0, "", err, nil
	}

	//查询应用数据
	appId, err := uuid.FromString(in.AppId)
	if err != nil {
		return 0, "", err, nil
	}
	appInfo, err := client.AppInfo.Query().Where(appinfo.IDEQ(appId)).First(ctx)
	if err != nil {
		return 0, "", err, nil
	}
	if appInfo == nil {
		return 0, "", errors.New("appInfo no data"), nil
	}

	//判断app_sdk是否由数据没有添加数据
	sdkId, err := uuid.FromString(in.SdkId)
	if err != nil {
		return 0, "", err, nil
	}
	where := client.AppSdk.Query().Where(appsdk.App(appId), appsdk.Sdk(sdkId))
	appSdkCount, err := where.Count(ctx)
	if err != nil {
		return 0, "", err, nil
	}
	var appSdk *ent.AppSdk
	if appSdkCount == 0 {
		appSdk, err = client.AppSdk.Create().SetApp(appId).SetSdk(sdkId).SetSdkKey(common.RandomString(32)).Save(ctx)
		if err != nil {
			return 0, "", err, nil
		}
	} else {
		appSdk, err = where.First(ctx)
		if err != nil {
			return 0, "", err, nil
		}
	}

	data := make([]*openned8.ActiveCodeInfo, 0)
	//创建激活码
	if err != nil {
		return 0, "", err, nil
	}
	for i := 0; i < int(in.GetQuantity()); i++ {
		save, err := client.ActiveCodeInfo.Create().
			SetActiveKey(common.RandomString(16)).
			SetUserID(in.UserId).
			SetAppID(appInfo.ID.String()).
			SetActiveIP("").
			SetDeviceSn("").
			SetDeviceMAC("").
			SetDeviceIdentity("").
			SetActiveType(0).
			SetActiveFile("").
			SetActiveFile("").
			SetVersion("v0.0.0").
			SetStartDate(time.Now()).
			SetExpireDate(time.Now().AddDate(0, 0, 3)).
			SetStatus(1).
			SetAppSdkID(appSdk.ID).
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
			AppSdkId:       save.AppSkdID.String(),
		})
	}
	return 0, "成功", nil, data
}
