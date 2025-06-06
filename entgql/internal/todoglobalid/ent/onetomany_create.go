// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entgql/internal/todoglobalid/ent/onetomany"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OneToManyCreate is the builder for creating a OneToMany entity.
type OneToManyCreate struct {
	config
	mutation *OneToManyMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (otmc *OneToManyCreate) SetName(s string) *OneToManyCreate {
	otmc.mutation.SetName(s)
	return otmc
}

// SetField2 sets the "field2" field.
func (otmc *OneToManyCreate) SetField2(s string) *OneToManyCreate {
	otmc.mutation.SetField2(s)
	return otmc
}

// SetNillableField2 sets the "field2" field if the given value is not nil.
func (otmc *OneToManyCreate) SetNillableField2(s *string) *OneToManyCreate {
	if s != nil {
		otmc.SetField2(*s)
	}
	return otmc
}

// SetParentID sets the "parent_id" field.
func (otmc *OneToManyCreate) SetParentID(i int) *OneToManyCreate {
	otmc.mutation.SetParentID(i)
	return otmc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (otmc *OneToManyCreate) SetNillableParentID(i *int) *OneToManyCreate {
	if i != nil {
		otmc.SetParentID(*i)
	}
	return otmc
}

// SetParent sets the "parent" edge to the OneToMany entity.
func (otmc *OneToManyCreate) SetParent(o *OneToMany) *OneToManyCreate {
	return otmc.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the OneToMany entity by IDs.
func (otmc *OneToManyCreate) AddChildIDs(ids ...int) *OneToManyCreate {
	otmc.mutation.AddChildIDs(ids...)
	return otmc
}

// AddChildren adds the "children" edges to the OneToMany entity.
func (otmc *OneToManyCreate) AddChildren(o ...*OneToMany) *OneToManyCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otmc.AddChildIDs(ids...)
}

// Mutation returns the OneToManyMutation object of the builder.
func (otmc *OneToManyCreate) Mutation() *OneToManyMutation {
	return otmc.mutation
}

// Save creates the OneToMany in the database.
func (otmc *OneToManyCreate) Save(ctx context.Context) (*OneToMany, error) {
	return withHooks(ctx, otmc.sqlSave, otmc.mutation, otmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (otmc *OneToManyCreate) SaveX(ctx context.Context) *OneToMany {
	v, err := otmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otmc *OneToManyCreate) Exec(ctx context.Context) error {
	_, err := otmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otmc *OneToManyCreate) ExecX(ctx context.Context) {
	if err := otmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otmc *OneToManyCreate) check() error {
	if _, ok := otmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OneToMany.name"`)}
	}
	if v, ok := otmc.mutation.Name(); ok {
		if err := onetomany.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OneToMany.name": %w`, err)}
		}
	}
	return nil
}

func (otmc *OneToManyCreate) sqlSave(ctx context.Context) (*OneToMany, error) {
	if err := otmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := otmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, otmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	otmc.mutation.id = &_node.ID
	otmc.mutation.done = true
	return _node, nil
}

func (otmc *OneToManyCreate) createSpec() (*OneToMany, *sqlgraph.CreateSpec) {
	var (
		_node = &OneToMany{config: otmc.config}
		_spec = sqlgraph.NewCreateSpec(onetomany.Table, sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt))
	)
	if value, ok := otmc.mutation.Name(); ok {
		_spec.SetField(onetomany.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := otmc.mutation.Field2(); ok {
		_spec.SetField(onetomany.FieldField2, field.TypeString, value)
		_node.Field2 = value
	}
	if nodes := otmc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onetomany.ParentTable,
			Columns: []string{onetomany.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := otmc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OneToManyCreateBulk is the builder for creating many OneToMany entities in bulk.
type OneToManyCreateBulk struct {
	config
	err      error
	builders []*OneToManyCreate
}

// Save creates the OneToMany entities in the database.
func (otmcb *OneToManyCreateBulk) Save(ctx context.Context) ([]*OneToMany, error) {
	if otmcb.err != nil {
		return nil, otmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(otmcb.builders))
	nodes := make([]*OneToMany, len(otmcb.builders))
	mutators := make([]Mutator, len(otmcb.builders))
	for i := range otmcb.builders {
		func(i int, root context.Context) {
			builder := otmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OneToManyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, otmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, otmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, otmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (otmcb *OneToManyCreateBulk) SaveX(ctx context.Context) []*OneToMany {
	v, err := otmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otmcb *OneToManyCreateBulk) Exec(ctx context.Context) error {
	_, err := otmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otmcb *OneToManyCreateBulk) ExecX(ctx context.Context) {
	if err := otmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
