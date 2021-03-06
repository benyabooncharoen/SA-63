// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/benyabooncharoen/app/ent/patient"
	"github.com/benyabooncharoen/app/ent/rightoftreatment"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RightoftreatmentCreate is the builder for creating a Rightoftreatment entity.
type RightoftreatmentCreate struct {
	config
	mutation *RightoftreatmentMutation
	hooks    []Hook
}

// SetRightoftreatmentName sets the rightoftreatmentName field.
func (rc *RightoftreatmentCreate) SetRightoftreatmentName(s string) *RightoftreatmentCreate {
	rc.mutation.SetRightoftreatmentName(s)
	return rc
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (rc *RightoftreatmentCreate) AddPatientIDs(ids ...int) *RightoftreatmentCreate {
	rc.mutation.AddPatientIDs(ids...)
	return rc
}

// AddPatients adds the patients edges to Patient.
func (rc *RightoftreatmentCreate) AddPatients(p ...*Patient) *RightoftreatmentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPatientIDs(ids...)
}

// Mutation returns the RightoftreatmentMutation object of the builder.
func (rc *RightoftreatmentCreate) Mutation() *RightoftreatmentMutation {
	return rc.mutation
}

// Save creates the Rightoftreatment in the database.
func (rc *RightoftreatmentCreate) Save(ctx context.Context) (*Rightoftreatment, error) {
	if _, ok := rc.mutation.RightoftreatmentName(); !ok {
		return nil, &ValidationError{Name: "rightoftreatmentName", err: errors.New("ent: missing required field \"rightoftreatmentName\"")}
	}
	var (
		err  error
		node *Rightoftreatment
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RightoftreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RightoftreatmentCreate) SaveX(ctx context.Context) *Rightoftreatment {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RightoftreatmentCreate) sqlSave(ctx context.Context) (*Rightoftreatment, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RightoftreatmentCreate) createSpec() (*Rightoftreatment, *sqlgraph.CreateSpec) {
	var (
		r     = &Rightoftreatment{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rightoftreatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rightoftreatment.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.RightoftreatmentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rightoftreatment.FieldRightoftreatmentName,
		})
		r.RightoftreatmentName = value
	}
	if nodes := rc.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rightoftreatment.PatientsTable,
			Columns: []string{rightoftreatment.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}
