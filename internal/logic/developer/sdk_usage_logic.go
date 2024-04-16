package developer

import (
	"context"

	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"

	"github.com/zeromicro/go-zero/core/logx"
)

type SdkUsageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSdkUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SdkUsageLogic {
	return &SdkUsageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SdkUsageLogic) SdkUsage(in *openned8.Empty) (*openned8.SdkUsage, error) {
	// todo: add your logic here and delete this line

	return &openned8.SdkUsage{}, nil
}
