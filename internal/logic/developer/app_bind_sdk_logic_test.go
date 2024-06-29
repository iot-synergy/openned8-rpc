package developer

import (
	"context"
	"flag"
	"testing"

	"github.com/iot-synergy/openned8-rpc/ent"
	"github.com/iot-synergy/openned8-rpc/internal/config"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestAppBindSdkLogic_AppBindSdk(t *testing.T) {
	type args struct {
		in *openned8.AppBindSdkReq
	}
	var configFile = flag.String("f", "../../../etc/openned8.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	tests := []struct {
		name    string
		l       *AppBindSdkLogic
		args    args
		want    *openned8.AppSdkInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "AppBindSdk",
			l: &AppBindSdkLogic{
				ctx: context.Background(),
				svcCtx: &svc.ServiceContext{
					Config: c,
					DB:     db,
					Redis:  c.RedisConf.MustNewUniversalRedis(),
				},
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &openned8.AppBindSdkReq{
					SdkId: "018f5c03-458c-7451-8624-d70ab751db1f",
					AppId: "P0ROQ77q7lbriWnyZMc7G1XC3wk5MYJO3xH2",
				},
			},
			want:    &openned8.AppSdkInfo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.AppBindSdk(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppBindSdkLogic.AppBindSdk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.l.Logger.Infow("AppBindSdk OK", logx.LogField{Key: "got", Value: got})
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("AppBindSdkLogic.AppBindSdk() = %v, want %v", got, tt.want)
			// }
		})
	}
}
