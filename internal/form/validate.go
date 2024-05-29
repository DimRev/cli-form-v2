package form

import (
	"errors"
	"fmt"
)

func (f *Field) validateField() error {
	for validatorKey, validatorValue := range f.Validators {
		switch validatorKey {
		case ValidatorTypes_Eq:
			if f.Value != validatorValue.Value {
				if validatorValue.CustomError != "" {
					return errors.New(validatorValue.CustomError)
				}
				return fmt.Errorf("error in %v, type %v", f.Label, validatorKey)
			}
		case ValidatorTypes_Gt:
			if f.Value != validatorValue.Value {
				if validatorValue.CustomError != "" {
					return errors.New(validatorValue.CustomError)
				}
				return fmt.Errorf("error in %v, type %v", f.Label, validatorKey)
			}
		case ValidatorTypes_Lt:
			if f.Value != validatorValue.Value {
				if validatorValue.CustomError != "" {
					return errors.New(validatorValue.CustomError)
				}
				return fmt.Errorf("error in %v, type %v", f.Label, validatorKey)
			}
		default:
		}
	}
	return nil
}
