package schema

import (
	o "github.com/boundedinfinity/optioner"
)

type StringSchema struct {
	SharedSchema
	Max o.Option[int] `json:"max" yaml:"max"`
	Min o.Option[int] `json:"min" yaml:"min"`
}

func (t StringSchema) Copy() Schema {
	return &StringSchema{
		SharedSchema: t.SharedSchema.copy(),
		Max:          t.Max,
		Min:          t.Min,
	}
}

func (t *StringSchema) Merge(other Schema) error {
	ot, ok := other.(*StringSchema)

	if !ok {
		return ErrWrongTypef("not StringSchema")
	}

	if err := t.SharedSchema.merge(ot.SharedSchema); err != nil {
		return err
	}

	t.Max = ot.Max.OrLast(ot.Max)
	t.Min = ot.Max.OrLast(ot.Min)

	return nil
}
