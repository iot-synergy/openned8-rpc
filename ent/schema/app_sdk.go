package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type AppSdk struct {
	ent.Schema
}

func (AppSdk) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("app", uuid.UUID{}).
			Optional().
			SchemaType(map[string]string{dialect.MySQL: "char(36)"}).
			Comment("app的id").
			Annotations(entsql.WithComments(true)),
		field.UUID("sdk", uuid.UUID{}).Optional().
			SchemaType(map[string]string{dialect.MySQL: "char(36)"}).
			Comment("sdk的id").
			Annotations(entsql.WithComments(true)),
		field.String("sdk_key").NotEmpty().
			SchemaType(map[string]string{dialect.MySQL: "char(32)"}).
			Comment("分配给app的sdk").
			Annotations(entsql.WithComments(true)),
	}
}

func (AppSdk) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

func (AppSdk) Indexes() []ent.Index {
	return []ent.Index{}
}

func (AppSdk) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_sdk"},
	}
}
func (AppSdk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("active_code", ActiveCodeInfo.Type),
		edge.From("app_info", AppInfo.Type).Ref("app_sdk").Field("app").Unique(),
		edge.From("sdk_info", SdkInfo.Type).Ref("app_sdk").Field("sdk").Unique(),
	}
}
