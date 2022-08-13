package manager

import "github.com/boundedinfinity/go-schemer/schema"

type SchemaManager struct {
	Schemas map[string]schema.Schema
	Unified schema.Schema
}

func New() *SchemaManager {
	return &SchemaManager{
		Schemas: make(map[string]schema.Schema),
	}
}
