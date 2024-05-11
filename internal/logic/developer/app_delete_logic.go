package developer

import (
	"context"
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

func (l *AppDeleteLogic) AppDelete(in *openned8.AppInfoDeleteReq) (*openned8.BeanMsg, error) {
	err := l.svcCtx.DB.AppInfo.DeleteOneID(in.Id).Where(appinfo.UserIDEQ(in.UserId)).Exec(l.ctx)
	if err != nil {
		return nil, err
	}
	return &openned8.BeanMsg{Msg: "成功"}, nil
}
