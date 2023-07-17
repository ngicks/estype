package fielddatatype

import (
	"fmt"
)

type InvalidTypeError struct {
	// name of Go type
	Type         string
	SupposedToBe []any
	InputValue   []byte
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf(
		"invalid type error: input is unacceptable to type %s."+
			" expected to be: one of %+v."+
			" actual: %s.",
		e.Type,
		e.SupposedToBe,
		e.InputValue,
	)
}
