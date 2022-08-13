package schema_test

import (
	"testing"

	"github.com/boundedinfinity/go-schemer/schema"
	o "github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
)

func createIntSchema() schema.IntegerSchema {
	return schema.IntegerSchema{
		SharedSchema: createSharedSchema(),
		Max:          o.Some(10),
		Min:          o.Some(1),
	}
}

func assertIntegerSchema(t *testing.T, expected, actual schema.IntegerSchema) {
	assertSharedSchema(t, expected.SharedSchema, actual.SharedSchema)

	assert.Equal(t, expected.Max.Defined(), actual.Max.Defined())
	assert.Equal(t, expected.Max.Get(), actual.Max.Get())

	assert.Equal(t, expected.Min.Defined(), actual.Min.Defined())
	assert.Equal(t, expected.Min.Get(), actual.Min.Get())
}

func Test_Integer_Merge_Min(t *testing.T) {
	input := createIntSchema()
	input.Min = o.Some(2)
	actual := createIntSchema()

	assert.Nil(t, actual.Merge(&input))

	expected := createIntSchema()
	expected.Min = o.Some(2)

	assertIntegerSchema(t, expected, actual)
}

func Test_Integer_Merge_Max(t *testing.T) {
	input := createIntSchema()
	input.Max = o.Some(10)
	actual := createIntSchema()

	assert.Nil(t, actual.Merge(&input))

	expected := createIntSchema()
	expected.Max = o.Some(10)

	assertIntegerSchema(t, expected, actual)
}

func Test_Integer_Copy(t *testing.T) {
	input := createIntSchema()

	actual := *input.Copy().(*schema.IntegerSchema)
	expected := createIntSchema()

	assertIntegerSchema(t, expected, actual)
}
