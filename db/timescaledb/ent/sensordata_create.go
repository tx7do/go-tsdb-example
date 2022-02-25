// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-tsdb-example/db/timescaledb/ent/sensordata"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SensorDataCreate is the builder for creating a SensorData entity.
type SensorDataCreate struct {
	config
	mutation *SensorDataMutation
	hooks    []Hook
}

// SetTime sets the "time" field.
func (sdc *SensorDataCreate) SetTime(t time.Time) *SensorDataCreate {
	sdc.mutation.SetTime(t)
	return sdc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (sdc *SensorDataCreate) SetNillableTime(t *time.Time) *SensorDataCreate {
	if t != nil {
		sdc.SetTime(*t)
	}
	return sdc
}

// SetSensorID sets the "sensor_id" field.
func (sdc *SensorDataCreate) SetSensorID(i int) *SensorDataCreate {
	sdc.mutation.SetSensorID(i)
	return sdc
}

// SetTemperature sets the "temperature" field.
func (sdc *SensorDataCreate) SetTemperature(f float64) *SensorDataCreate {
	sdc.mutation.SetTemperature(f)
	return sdc
}

// SetCPU sets the "cpu" field.
func (sdc *SensorDataCreate) SetCPU(f float64) *SensorDataCreate {
	sdc.mutation.SetCPU(f)
	return sdc
}

// Mutation returns the SensorDataMutation object of the builder.
func (sdc *SensorDataCreate) Mutation() *SensorDataMutation {
	return sdc.mutation
}

// Save creates the SensorData in the database.
func (sdc *SensorDataCreate) Save(ctx context.Context) (*SensorData, error) {
	var (
		err  error
		node *SensorData
	)
	if len(sdc.hooks) == 0 {
		if err = sdc.check(); err != nil {
			return nil, err
		}
		node, err = sdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdc.check(); err != nil {
				return nil, err
			}
			sdc.mutation = mutation
			if node, err = sdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sdc.hooks) - 1; i >= 0; i-- {
			if sdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SensorDataCreate) SaveX(ctx context.Context) *SensorData {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SensorDataCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SensorDataCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SensorDataCreate) check() error {
	if _, ok := sdc.mutation.SensorID(); !ok {
		return &ValidationError{Name: "sensor_id", err: errors.New(`ent: missing required field "SensorData.sensor_id"`)}
	}
	if _, ok := sdc.mutation.Temperature(); !ok {
		return &ValidationError{Name: "temperature", err: errors.New(`ent: missing required field "SensorData.temperature"`)}
	}
	if _, ok := sdc.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "SensorData.cpu"`)}
	}
	return nil
}

func (sdc *SensorDataCreate) sqlSave(ctx context.Context) (*SensorData, error) {
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sdc *SensorDataCreate) createSpec() (*SensorData, *sqlgraph.CreateSpec) {
	var (
		_node = &SensorData{config: sdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sensordata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensordata.FieldID,
			},
		}
	)
	if value, ok := sdc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sensordata.FieldTime,
		})
		_node.Time = &value
	}
	if value, ok := sdc.mutation.SensorID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sensordata.FieldSensorID,
		})
		_node.SensorID = value
	}
	if value, ok := sdc.mutation.Temperature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldTemperature,
		})
		_node.Temperature = value
	}
	if value, ok := sdc.mutation.CPU(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldCPU,
		})
		_node.CPU = value
	}
	return _node, _spec
}

// SensorDataCreateBulk is the builder for creating many SensorData entities in bulk.
type SensorDataCreateBulk struct {
	config
	builders []*SensorDataCreate
}

// Save creates the SensorData entities in the database.
func (sdcb *SensorDataCreateBulk) Save(ctx context.Context) ([]*SensorData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SensorData, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SensorDataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SensorDataCreateBulk) SaveX(ctx context.Context) []*SensorData {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SensorDataCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SensorDataCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
