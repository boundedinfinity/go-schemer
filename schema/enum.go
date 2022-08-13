package schema

import (
	"github.com/boundedinfinity/go-schemer/enum_format"
	o "github.com/boundedinfinity/optioner"
)

type EnumSchema struct {
	SharedSchema
	Format o.Option[enum_format.EnumFormat] `json:"format" yaml:"format"`
}
