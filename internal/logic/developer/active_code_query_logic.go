package developer

import (
	"context"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveCodeQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActiveCodeQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveCodeQueryLogic {
	return &ActiveCodeQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActiveCodeQueryLogic) ActiveCodeQuery(in *openned8.ActiveCodeListReq) (*openned8.ActiveCodeListInfo, error) {
	where := l.svcCtx.DB.ActiveCodeInfo.Query().Where(activecodeinfo.UserIDEQ(in.UserId))
	if in.AppId != "" {
		where.Where(activecodeinfo.AppIDEQ(in.AppId))
	}
	if in.AppSdkId != "" {
		where.Where(activecodeinfo.AppSdkIDEQ(uuid.FromStringOrNil(in.AppSdkId)))
	}
	if in.Status != 0 {
		where.Where(activecodeinfo.StatusEQ(uint8(in.Status)))
	}
	if in.ExpireDate != 0 {
		where.Where(activecodeinfo.ExpireDateLTE(time.UnixMilli(in.ExpireDate)))
	}
	count, err := where.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	all, err := where.Offset(int((in.PageInfo.Page - 1) * in.PageInfo.PageSize)).Limit(int(in.PageInfo.PageSize)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.ActiveCodeInfo, 0)
	for _, info := range all {
		result = append(result, &openned8.ActiveCodeInfo{
			Id:             info.ID.String(),
			CreatedAt:      info.CreatedAt.UnixMilli(),
			UpdatedAt:      info.UpdatedAt.UnixMilli(),
			ActiveKey:      info.ActiveKey,
			UserId:         info.UserID,
			AppId:          info.AppID,
			ActiveIP:       info.ActiveIP,
			DeviceSN:       info.DeviceSn,
			DeviceMac:      info.DeviceMAC,
			DeviceIdentity: info.DeviceIdentity,
			ActiveDate:     info.ActiveDate.UnixMilli(),
			ActiveType:     info.ActiveType,
			ActiveFile:     info.ActiveFile,
			Version:        info.Version,
			StartDate:      info.StartDate.UnixMilli(),
			ExpireDate:     info.ExpireDate.UnixMilli(),
			Status:         int64(info.Status),
		})
	}
	return &openned8.ActiveCodeListInfo{
		Total: uint64(count),
		Data:  result,
	}, nil
}
