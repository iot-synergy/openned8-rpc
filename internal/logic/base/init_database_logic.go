package base

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/bsm/redislock"
	"github.com/iot-synergy/openned8-rpc/internal/svc"
	"github.com/iot-synergy/openned8-rpc/types/openned8"
	"github.com/iot-synergy/synergy-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *openned8.Empty) (*openned8.BaseMsg, error) {
	// todo: add your logic here and delete this line
	// If your mysql speed is high, comment the code below.
	// Because the context deadline will reach if the database is too slow
	l.ctx = context.Background()

	// add lock to avoid duplicate initialization
	locker := redislock.New(l.svcCtx.Redis)

	lock, err := locker.Obtain(l.ctx, "INIT:DATABASE:LOCK", 10*time.Minute, nil)
	if errors.Is(err, redislock.ErrNotObtained) {
		logx.Error("last initialization is running")
		return nil, err
	} else if err != nil {
		logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
		return nil, err
	}

	defer lock.Release(l.ctx)

	// initialize table structure
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", err.Error(), 300*time.Second).Err()
		return nil, err
	}

	return &openned8.BaseMsg{}, nil
}
