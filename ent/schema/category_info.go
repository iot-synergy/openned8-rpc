package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type CategoryInfo struct {
	ent.Schema
}

func (CategoryInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("名字").
			Annotations(entsql.WithComments(true)),
	}
}

func (CategoryInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (CategoryInfo) Indexes() []ent.Index {
	return []ent.Index{}
}

func (CategoryInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "category_info"},
	}
}
