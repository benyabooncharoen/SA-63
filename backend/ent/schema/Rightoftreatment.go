package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Rightoftreatment holds the schema definition for the Rightoftreatment entity.
type Rightoftreatment struct {
	ent.Schema
}

// Fields of the Rightoftreatment.
func (Rightoftreatment) Fields() []ent.Field {
	return []ent.Field{
		field.String("rightoftreatmentName"),
	}
}

// Edges of the Rightoftreatment.
func (Rightoftreatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patients", Patient.Type).StorageKey(edge.Column("Rightoftreatment")),
	}
}