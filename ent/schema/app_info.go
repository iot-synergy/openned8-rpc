package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type AppInfo struct {
	ent.Schema
}

func (AppInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("用户id").
			Annotations(entsql.WithComments(true)),
		field.String("app_name").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("app名字").
			Annotations(entsql.WithComments(true)),
		field.String("summary").
			SchemaType(map[string]string{dialect.MySQL: "varchar(1024)"}).
			Comment("摘要").
			Annotations(entsql.WithComments(true)),
		field.Int64("app_category").
			Comment("种类").
			Annotations(entsql.WithComments(true)),
		field.Int64("use_industry").
			Comment("使用领域").
			Annotations(entsql.WithComments(true)),
		field.String("app_category_name").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("种类名字").
			Annotations(entsql.WithComments(true)),
		field.String("use_industry_name").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("使用领域名字").
			Annotations(entsql.WithComments(true)),
		field.String("app_key").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("钥匙").
			Annotations(entsql.WithComments(true)),
		field.String("app_secret").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("秘密").
			Annotations(entsql.WithComments(true)),
	}
}

func (AppInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

func (AppInfo) Indexes() []ent.Index {
	return []ent.Index{}
}

func (AppInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_info"},
	}
}

func (AppInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app_sdk", AppSdk.Type),
	}
}
