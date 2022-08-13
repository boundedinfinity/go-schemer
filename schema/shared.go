package schema

import (
	"fmt"

	"github.com/boundedinfinity/go-schemer/schema_type"
	o "github.com/boundedinfinity/optioner"
)

type Schema interface {
	GetId() o.Option[string]
	GetType() o.Option[schema_type.SchemaType]
	Copy() Schema
	Merge(Schema) error
}

type SharedSchema struct {
	Id          o.Option[string]                 `json:"$id" yaml:"$id"`
	Type        o.Option[schema_type.SchemaType] `json:"$type" yaml:"$type"`
	Ref         o.Option[string]                 `json:"$ref" yaml:"$ref"`
	Name        o.Option[string]                 `json:"name" yaml:"name"`
	Required    o.Option[bool]                   `json:"required" yaml:"required"`
	Title       o.Option[string]                 `json:"title" yaml:"title"`
	Description o.Option[string]                 `json:"description" yaml:"description"`
}

func (s SharedSchema) GetId() o.Option[string] {
	return s.Id
}

func (s SharedSchema) GetType() o.Option[schema_type.SchemaType] {
	return s.Type
}

func (s SharedSchema) copy() SharedSchema {
	return SharedSchema{
		Id:          s.Id,
		Type:        s.Type,
		Name:        s.Name,
		Ref:         s.Ref,
		Required:    s.Required,
		Title:       s.Title,
		Description: s.Description,
	}
}

func (t *SharedSchema) merge(other SharedSchema) error {
	if err := errIfDiff(t.Id, other.Id, ErrCantMergeType); err != nil {
		return err
	}

	if err := errIfDiff(t.Type, other.Type, ErrCantMergeType); err != nil {
		return err
	}

	t.Name = t.Name.OrLast(other.Name)
	t.Title = t.Title.OrLast(other.Title)
	t.Required = t.Required.OrLast(other.Required)

	if other.Description.Defined() && other.Description.Get() != t.Description.Get() {
		d := other.Description.Get()

		if t.Description.Defined() && t.Description.Get() != "" {
			d += fmt.Sprintf("\n%v", t.Description.Get())
		}

		t.Description = o.Some(d)
	}

	return nil
}
