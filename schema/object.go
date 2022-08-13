package schema

import o "github.com/boundedinfinity/optioner"

type ObjectShema struct {
	SharedSchema
	Properties o.Option[map[string]Schema] `json:"properties" yaml:"properties"`
}
