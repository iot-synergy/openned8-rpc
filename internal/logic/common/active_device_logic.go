package common

import (
	"context"
	"errors"
	"github.com/iot-synergy/openned8-rpc/ent"
	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActiveDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveDeviceLogic {
	return &ActiveDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActiveDeviceLogic) ActiveDevice(in *openned8.ActiveDeviceReq) (*openned8.ActiveDeviceResp, error) {
	tx, err := l.svcCtx.DB.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	client := tx.Client()
	data, err := activeDevice(l.ctx, client, in)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			l.Logger.Error("rollback fail:" + err1.Error())
		}
		return nil, err1
	} else {
		err := tx.Commit()
		if err != nil {
			l.Logger.Error("commit fail:" + err.Error())
			return nil, err
		}
	}
	return data, nil
}

func activeDevice(ctx context.Context, client *ent.Client, in *openned8.ActiveDeviceReq) (*openned8.ActiveDeviceResp, error) {
	activeCode, err := client.AppSdk.Query().Where(appsdk.AppEQ(in.AppId)).QueryActiveCode().Where(activecodeinfo.ActiveKeyEQ(in.ActiveCode)).First(ctx)
	if err != nil {
		return nil, err
	}
	if activeCode == nil {
		return nil, errors.New("data is empty")
	}
	if activeCode.DeviceSn != "" && activeCode.DeviceSn != in.DeviceSn {
		return nil, errors.New("The activation code cannot activate two devices")
	}
	activeCode.Status = 2
	activeCode.Imei = in.Imei
	activeCode.DeviceSn = in.DeviceSn
	client.ActiveCodeInfo.UpdateOne(activeCode)
	return &openned8.ActiveDeviceResp{}, nil
}
