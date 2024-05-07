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

type SdkInfo struct {
	ent.Schema
}

func (SdkInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("名字").
			Annotations(entsql.WithComments(true)),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Comment("头像").
			Annotations(entsql.WithComments(true)),
		field.Int64("desc").
			Comment("排序").
			Annotations(entsql.WithComments(true)),
		field.String("download_url").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Comment("下载地址").
			Annotations(entsql.WithComments(true)),
	}
}

func (SdkInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (SdkInfo) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SdkInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sdk_info"},
	}
}

func (SdkInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app_sdk", AppSdk.Type),
	}
}
