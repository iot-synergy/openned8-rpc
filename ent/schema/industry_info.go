package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type IndustryInfo struct {
	ent.Schema
}

func (IndustryInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("名字").
			Annotations(entsql.WithComments(true)),
	}
}

func (IndustryInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (IndustryInfo) Indexes() []ent.Index {
	return []ent.Index{}
}

func (IndustryInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "industry_info"},
	}
}
