package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type SdkUsage struct {
	ent.Schema
}

func (SdkUsage) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Unique().
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("用户id").
			Annotations(entsql.WithComments(true)),
		field.Int64("all").
			Default(0).
			Comment("最大使用量").
			Annotations(entsql.WithComments(true)),
		field.Int64("used").
			Default(0).
			Comment("已使用量").
			Annotations(entsql.WithComments(true)),
	}
}

func (SdkUsage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

func (SdkUsage) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SdkUsage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sdk_usage"},
	}
}
