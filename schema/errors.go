package schema

import (
	"errors"
	"fmt"

	o "github.com/boundedinfinity/optioner"
)

var (
	ErrMissingId        = errors.New("missing schema $id")
	ErrMissingType      = errors.New("missing schema $type")
	ErrNotImplemented   = errors.New("not implemented")
	ErrNotImplementedf  = errPrepend(ErrNotImplemented)
	ErrCantCombineType  = errors.New("can't combine types")
	ErrCantCombineTypef = errPrepend(ErrCantCombineType)
	ErrCantMergeType    = errors.New("can't merge type")
	ErrCantMergeTypef   = errPrepend(ErrCantMergeType)
	ErrWrongType        = errors.New("wrong type")
	ErrWrongTypef       = errPrepend(ErrWrongType)
)

func errPrepend(err error) func(format string, a ...any) error {
	return func(format string, a ...any) error {
		message := fmt.Sprintf(format, a)
		return fmt.Errorf("%v : %w", message, err)
	}
}

func errAppend(err error) func(format string, a ...any) error {
	return func(format string, a ...any) error {
		message := fmt.Sprintf(format, a)
		return fmt.Errorf("%w : %v", err, message)
	}
}

func errIfDiff[T comparable](f, s o.Option[T], err error) error {
	if f.Defined() && s.Defined() && f.Get() != s.Get() {
		return err
	}

	return nil
}
