package schema

import (
	o "github.com/boundedinfinity/optioner"
)

type IntegerSchema struct {
	SharedSchema
	Max o.Option[int] `json:"max" yaml:"max"`
	Min o.Option[int] `json:"min" yaml:"min"`
}

func (t IntegerSchema) Copy() Schema {
	return &IntegerSchema{
		SharedSchema: t.SharedSchema.copy(),
		Max:          t.Max,
		Min:          t.Min,
	}
}

func (t *IntegerSchema) Merge(other Schema) error {
	ot, ok := other.(*IntegerSchema)

	if !ok {
		return ErrWrongTypef("not IntegerSchema")
	}

	t.merge(ot.SharedSchema)
	t.Max = t.Max.OrLast(ot.Max)
	t.Min = t.Min.OrLast(ot.Min)

	return nil
}
