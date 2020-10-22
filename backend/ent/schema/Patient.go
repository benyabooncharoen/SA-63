package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("hn").Unique(),
		field.String("patientName"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("patients").Unique(),
		edge.From("rightoftreatment", Rightoftreatment.Type).Ref("patients").Unique(),
		edge.From("systemmember", Systemmember.Type).Ref("patients").Unique(),
	}
}