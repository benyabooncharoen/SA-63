package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Systemmember holds the schema definition for the Systemmember entity.
type Systemmember struct {
	ent.Schema
}

// Fields of the Systemmember.
func (Systemmember) Fields() []ent.Field {
	return []ent.Field{
		field.String("systemmemberName").
			NotEmpty(),
		field.String("Password").
			NotEmpty(),
	}
}

// Edges of the Systemmember.
func (Systemmember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patients", Patient.Type).StorageKey(edge.Column("Systemmember")),
	}
}