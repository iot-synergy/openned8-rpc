package schema

import (
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
	"github.com/gofrs/uuid/v5"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type ActiveCodeInfo struct {
	ent.Schema
}

func (ActiveCodeInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("active_key").
			SchemaType(map[string]string{dialect.MySQL: "char(16)"}).
			Unique().
			Comment("激活码").
			Annotations(entsql.WithComments(true)),
		field.String("user_id").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("用户id").
			Annotations(entsql.WithComments(true)),
		field.String("app_id").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("appid").
			Annotations(entsql.WithComments(true)),
		field.String("active_ip").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("激活的ip").
			Annotations(entsql.WithComments(true)),
		field.String("device_sn").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("设备的sn码").
			Annotations(entsql.WithComments(true)),
		field.String("device_mac").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("设备的mac").
			Annotations(entsql.WithComments(true)),
		field.String("device_identity").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("设备的身份").
			Annotations(entsql.WithComments(true)),
		field.Time("active_date").
			Comment("激活时间").
			Default(time.Now()).
			Annotations(entsql.WithComments(true)),
		field.Int64("active_type").
			Default(0).
			Comment("激活类型").
			Annotations(entsql.WithComments(true)),
		field.String("active_file").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("激活的文件").
			Annotations(entsql.WithComments(true)),
		field.String("version").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("版本").
			Annotations(entsql.WithComments(true)),
		field.Time("start_date").
			Comment("开始时间").
			Default(time.Now()).
			Annotations(entsql.WithComments(true)),
		field.Time("expire_date").
			Comment("结束时间").
			Annotations(entsql.WithComments(true)),
		field.UUID("app_sdk_id", uuid.UUID{}).
			Optional().
			Comment("关联app_key").
			Annotations(entsql.WithComments(true)),
	}
}

func (ActiveCodeInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (ActiveCodeInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("active_key").Unique(),
	}
}

func (ActiveCodeInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "active_code_info"},
	}
}

func (ActiveCodeInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app_sdk", AppSdk.Type).Ref("active_code").Field("app_sdk_id").Unique(),
	}
}
