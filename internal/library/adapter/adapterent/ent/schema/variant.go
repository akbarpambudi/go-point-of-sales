package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Variant struct {
	ent.Schema
}

// Fields of the Category.
func (Variant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("code"),
		field.String("name"),
		field.Float("price"),
	}
}

// Edges of the Category.
func (Variant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("variants").
			Unique(),
	}
}
