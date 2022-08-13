package manager_test

import (
	"testing"

	"github.com/boundedinfinity/go-schemer/manager"
	"github.com/stretchr/testify/assert"
)

func createManager() *manager.SchemaManager {
	m := manager.New()
	return m
}

func Test_String(t *testing.T) {
	m := createManager()

	// m.Update(schema.StringSchema{
	// 	SharedSchema: schema.SharedSchema{
	// 		Id:          o.Some("https://www.boundedinfinity.com/schema/string-1"),
	// 		Name:        o.None[string](),
	// 		Type:        o.Some(schema_type.String),
	// 		Ref:         o.None[string](),
	// 		Required:    o.None[bool](),
	// 		Title:       o.None[string](),
	// 		Description: o.None[string](),
	// 	},
	// 	Min: o.Some(1),
	// 	Max: o.Some(10),
	// })

	assert.True(t, m.Unified.GetId().Defined())
}
