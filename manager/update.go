package manager

import (
	"github.com/boundedinfinity/go-schemer/schema"
	"github.com/boundedinfinity/go-schemer/schema_type"
)

func (m *SchemaManager) Update(s schema.Schema) error {
	if s.GetId().Empty() {
		return schema.ErrMissingId
	}

	if s.GetType().Empty() {
		return schema.ErrMissingType
	}

	if _, ok := m.Schemas[s.GetId().Get()]; !ok {
		m.Schemas[s.GetId().Get()] = s
	}

	if m.Unified == nil {
		m.Unified = s.Copy()
	} else if m.Unified.GetType().Get() != s.GetType().Get() {
		return schema.ErrCantCombineTypef(
			"%v and %v", m.Unified.GetType(), s.GetType(),
		)
	} else {
		switch s.GetType().Get() {
		case schema_type.String:

		default:
			return schema.ErrNotImplementedf("$type %v", s.GetId().Get())
		}
	}

	return nil
}
