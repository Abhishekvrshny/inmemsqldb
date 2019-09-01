package constraint

import "errors"

type ConstraintType string

const (
	NOTNULL ConstraintType = "NOTNULL"
)

type Constraint interface {
	Validate(interface{}) error
}

func Validate(cType ConstraintType, value interface{}) error {
	switch cType {
	case "NOTNULL":
		return new(NotNULL).Validate(value)
	default:
		return errors.New("unknown constraint type")
	}
}
