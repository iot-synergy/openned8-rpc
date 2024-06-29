package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
	"github.com/iot-synergy/openned8-rpc/internal/common"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppBindSdkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppBindSdkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppBindSdkLogic {
	return &AppBindSdkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppBindSdkLogic) AppBindSdk(in *openned8.AppBindSdkReq) (*openned8.AppSdkInfo, error) {
	tx, err := l.svcCtx.DB.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	client := tx.Client()

	appInfo, err := client.AppInfo.Query().Where(appinfo.IDEQ(in.AppId)).First(l.ctx)
	if err != nil {
		return nil, err
	}

	sdkInfo, err := client.SdkInfo.Query().Where(sdkinfo.IDEQ(in.SdkId)).First(l.ctx)
	if err != nil {
		return nil, err
	}

	appSdk := client.AppSdk.Create().SetApp(in.AppId).SetSdk(in.SdkId).
		SetSdkKey(common.RandomString(32)).
		SaveX(l.ctx)

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &openned8.AppSdkInfo{
		Id:          appSdk.ID.String(),
		AppId:       appSdk.App,
		AppName:     appInfo.AppName,
		SdkId:       sdkInfo.ID,
		SdkName:     sdkInfo.Name,
		SdkKey:      appSdk.SdkKey,
		SdkAvatar:   sdkInfo.Avatar,
		SdkDesc:     sdkInfo.Desc,
		DownloadUrl: sdkInfo.DownloadURL,
	}, nil
}
