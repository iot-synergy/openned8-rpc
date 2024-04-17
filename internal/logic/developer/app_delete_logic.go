package developer

import (
	"context"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppDeleteLogic {
	return &AppDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppDeleteLogic) AppDelete(in *openned8.AppInfo) (*openned8.BeanMsg, error) {
	fromString, err := uuid.FromString(in.UserId)
	if err != nil {
		return nil, err
	}
	l.svcCtx.DB.AppInfo.Delete().Where(appinfo.IDEQ(fromString))

	return &openned8.BeanMsg{Msg: "成功"}, nil
}
