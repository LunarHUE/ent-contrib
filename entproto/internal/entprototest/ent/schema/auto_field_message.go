package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AutoFieldMessage holds the schema definition for testing AutoField numbering.
type AutoFieldMessage struct {
	ent.Schema
}

func (AutoFieldMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entproto.AutoField()),
		field.Int32("age").Annotations(entproto.AutoField()),
	}
}

func (AutoFieldMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{entproto.Message()}
}
