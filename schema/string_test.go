package schema_test

import (
	"testing"

	"github.com/boundedinfinity/go-schemer/schema"
	o "github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
)

func createStringSchema() schema.StringSchema {
	return schema.StringSchema{
		SharedSchema: createSharedSchema(),
		Max:          o.Some(10),
		Min:          o.Some(1),
	}
}

func assertStringSchema(t *testing.T, expected, actual schema.StringSchema) {
	assertSharedSchema(t, expected.SharedSchema, actual.SharedSchema)

	assert.Equal(t, expected.Max.Defined(), actual.Max.Defined())
	assert.Equal(t, expected.Max.Get(), actual.Max.Get())

	assert.Equal(t, expected.Min.Defined(), actual.Min.Defined())
	assert.Equal(t, expected.Min.Get(), actual.Min.Get())
}

func Test_String_Merge_Min(t *testing.T) {
	input := createStringSchema()
	input.Min = o.Some(2)
	actual := createStringSchema()

	assert.Nil(t, actual.Merge(&input))

	expected := createStringSchema()
	expected.Min = o.Some(2)

	assertStringSchema(t, expected, actual)
}

func Test_String_Merge_Max(t *testing.T) {
	input := createStringSchema()
	input.Max = o.Some(10)
	actual := createStringSchema()

	assert.Nil(t, actual.Merge(&input))

	expected := createStringSchema()
	expected.Max = o.Some(10)

	assertStringSchema(t, expected, actual)
}

func Test_String_Copy(t *testing.T) {
	input := createStringSchema()

	actual := *input.Copy().(*schema.StringSchema)
	expected := createStringSchema()

	assertStringSchema(t, expected, actual)
}
