package constraint

import "errors"

type Type string

const (
	NOTNULL Type = "NOTNULL"
)

type Constraint interface {
	Validate(interface{}) error
}

func Validate(cType Type, value interface{}) error {
	switch cType {
	case "NOTNULL":
		return new(NotNULL).Validate(value)
	default:
		return errors.New("unknown constraint type")
	}
}
