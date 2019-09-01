package field

import "errors"

type FieldType string

const (
	INT    FieldType = "INT"
	STRING FieldType = "STRING"
)

type Field interface {
	validate() error
}

func Validate(f FieldType, value interface{}) error {
	switch f {
	case "INT":
		if value != nil {
			fld := Int(value.(int))
			return fld.validate()
		}
		return nil
	case "STRING":
		if value != nil {
			fld := String(value.(string))
			return fld.validate()
		}
		return nil
	default:
		return errors.New("unknown type")
	}
}
