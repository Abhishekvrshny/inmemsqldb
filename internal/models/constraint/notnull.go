package constraint

import "errors"

type NotNULL string

func (nn *NotNULL) Validate(value interface{}) error {
	if value == nil {
		return errors.New("NOTNULL constraint failed")
	}
	return nil
}
