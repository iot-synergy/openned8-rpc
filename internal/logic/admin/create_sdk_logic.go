package admin

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSdkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSdkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSdkLogic {
	return &CreateSdkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSdkLogic) CreateSdk(in *openned8.SdkInfoCreateReq) (*openned8.SdkInfo, error) {
	data, err := l.svcCtx.DB.SdkInfo.Create().SetName(in.Name).SetAvatar(in.Avatar).SetDesc(in.Desc).SetDownloadURL(in.DownloadUrl).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &openned8.SdkInfo{
		Id:          data.ID.String(),
		Name:        data.Name,
		Avatar:      data.Avatar,
		Desc:        data.Desc,
		DownloadUrl: data.DownloadURL,
	}, nil
}
