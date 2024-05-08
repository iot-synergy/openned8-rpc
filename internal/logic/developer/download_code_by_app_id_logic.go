package developer

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCodeByAppIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDownloadCodeByAppIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCodeByAppIdLogic {
	return &DownloadCodeByAppIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DownloadCodeByAppIdLogic) DownloadCodeByAppId(in *openned8.DownloadCodeByAppIdReq) (*openned8.DownloadCodeByAppIdResp, error) {
	appId, err := uuid.FromString(in.AppId)
	if err != nil {
		return nil, err
	}
	r, err := regexp.Compile(`^[A-Za-z0-9]+$`)
	if err != nil {
		return nil, err
	}
	if !r.Match([]byte(in.UserId)) {
		return nil, errors.New("wrong userId")
	}
	sql := "SELECT t1.app_key,t2.sdk_key,t3.active_key FROM `app_info` t1 " +
		"INNER JOIN `app_sdk` t2 ON t1.id = t2.app " +
		"INNER JOIN `active_code_info` t3 ON t2.id = t3.app_sdk_id " +
		"WHERE t1.id = " + appId.String() + " AND t1.user_id = " + in.UserId
	data, err := l.svcCtx.DB.QueryContext(l.ctx, sql)
	if err != nil {
		return nil, err
	}
	result := make([]*openned8.DownloadCode, 0)
	for data.Next() {
		downloadCode := new(openned8.DownloadCode)
		data.Scan(&downloadCode.AppKey, &downloadCode.SdkKey, &downloadCode.Code)
		result = append(result, downloadCode)
	}

	return &openned8.DownloadCodeByAppIdResp{Data: result}, nil
}
