package schema_test

import (
	"testing"

	"github.com/boundedinfinity/go-schemer/schema"
	"github.com/boundedinfinity/go-schemer/schema_type"
	o "github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
)

func createSharedSchema() schema.SharedSchema {
	return schema.SharedSchema{
		Id:          o.Some("some-id"),
		Type:        o.Some(schema_type.Integer),
		Ref:         o.Some("some-ref"),
		Required:    o.Some(false),
		Name:        o.Some("some-name"),
		Title:       o.Some("some-title"),
		Description: o.Some("some-description"),
	}
}

func assertSharedSchema(t *testing.T, expected, actual schema.SharedSchema) {
	assert.Equal(t, expected.Id.Defined(), actual.Id.Defined(), "id dosen't match")
	assert.Equal(t, expected.Id.Get(), actual.Id.Get(), "id dosen't match")

	assert.Equal(t, expected.Type.Defined(), actual.Type.Defined(), "type dosen't match")
	assert.Equal(t, expected.Type.Get(), actual.Type.Get(), "type dosen't match")

	assert.Equal(t, expected.Name.Defined(), actual.Name.Defined(), "name dosen't match")
	assert.Equal(t, expected.Name.Get(), actual.Name.Get(), "name dosen't match")

	assert.Equal(t, expected.Ref.Defined(), actual.Ref.Defined(), "ref dosen't match")
	assert.Equal(t, expected.Ref.Get(), actual.Ref.Get(), "ref dosen't match")

	assert.Equal(t, expected.Required.Defined(), actual.Required.Defined())
	assert.Equal(t, expected.Required.Get(), actual.Required.Get())

	assert.Equal(t, expected.Title.Defined(), actual.Title.Defined())
	assert.Equal(t, expected.Title.Get(), actual.Title.Get())

	assert.Equal(t, expected.Description.Defined(), actual.Description.Defined())
	assert.Equal(t, expected.Description.Get(), actual.Description.Get())
}

func Test_Shared_Merge_Description(t *testing.T) {
	input := createSharedSchema()
	input.Description = o.None[string]()
	actual := createSharedSchema()
	assert.Nil(t, actual.MergeT(input))
	expected := createSharedSchema()
	expected.Description = o.Some("some-description")
	assertSharedSchema(t, expected, actual)

	input = createSharedSchema()
	input.Description = o.Some("another-description")
	actual = createSharedSchema()
	assert.Nil(t, actual.MergeT(input))
	expected.Description = o.Some("another-description\nsome-description")
	assertSharedSchema(t, expected, actual)

	input = createSharedSchema()
	input.Description = o.Some("another-description")
	actual = createSharedSchema()
	actual.Description = o.None[string]()
	assert.Nil(t, actual.MergeT(input))
	expected.Description = o.Some("another-description")
	assertSharedSchema(t, expected, actual)
}

func Test_Shared_Merge_Title(t *testing.T) {
	input := createSharedSchema()
	input.Title = o.Some("another-title")
	actual := createSharedSchema()

	assert.Nil(t, actual.MergeT(input))

	expected := createSharedSchema()
	expected.Title = o.Some("another-title")

	assertSharedSchema(t, expected, actual)
}

func Test_Shared_Merge_Name(t *testing.T) {
	input := createSharedSchema()
	input.Name = o.Some("another-name")
	actual := createSharedSchema()

	assert.Nil(t, actual.MergeT(input))

	expected := createSharedSchema()
	expected.Name = o.Some("another-name")

	assertSharedSchema(t, expected, actual)
}

func Test_Shared_Merge_Id(t *testing.T) {
	input := createSharedSchema()
	input.Id = o.Some("another-id")
	actual := createSharedSchema()
	assert.NotNil(t, actual.MergeT(input))
}

func Test_Shared_Merge_Type(t *testing.T) {
	input := createSharedSchema()
	input.Type = o.Some(schema_type.String)
	actual := createSharedSchema()
	assert.NotNil(t, actual.MergeT(input))
}

func Test_Shared_Copy(t *testing.T) {
	input := createSharedSchema()

	actual := input.CopyT()
	expected := createSharedSchema()

	assertSharedSchema(t, expected, actual)
}
