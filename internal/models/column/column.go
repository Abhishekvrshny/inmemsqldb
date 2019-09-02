package column

import (
	"inmemdb/internal/models/constraint"
	"inmemdb/internal/models/field"
)

type Column struct {
	name        string
	typ         field.Type
	constraints []constraint.Type
}

func New(name string, f field.Type, constraints []constraint.Type) *Column {
	return &Column{name, f, constraints}
}

func (c *Column) GetName() string {
	return c.name
}

func (c *Column) GetType() field.Type {
	return c.typ
}

func (c *Column) Validate(value interface{}) error {
	for _, c := range c.constraints {
		err := constraint.Validate(c, value)
		if err != nil {
			return err
		}
	}
	return nil
}
