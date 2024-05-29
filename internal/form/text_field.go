package form

import "errors"

type TextField struct {
	Label      string
	Value      string
	Validators map[ValidatorTypes]Validator
}

func (f *Form) AddTextField(id, label string) (*TextField, error) {
	_, ok := f.Inputs[id]
	if ok {
		return nil, errors.New("cannot have two form inputs of the same id")
	}
	field := &TextField{
		Label:      label,
		Validators: make(map[ValidatorTypes]Validator),
	}
	f.Inputs[id] = field
	return field, nil
}

func (t *TextField) AddValidator(errMsg string, validatorType ValidatorTypes, value interface{}) error {
	_, ok := t.Validators[validatorType]
	if ok {
		return errors.New("cannot have two validators of the same type")
	}

	t.Validators[validatorType] = Validator{
		CustomError: errMsg,
		Value:       value,
	}
	return nil
}

func (t *TextField) Validate() error {
	for _, validator := range t.Validators {
		if validatorType, ok := validator.Value.(ValidatorTypes); ok {
			switch validatorType {
			case ValidatorTypes_Gt:
				intValue, ok := validator.Value.(int)
				if !ok {
					return errors.New("invalid value type for text field validator")
				}
				if len(t.Value) < intValue {
					return errors.New(validator.CustomError)
				}

			case ValidatorTypes_Lt:
				intValue, ok := validator.Value.(int)
				if !ok {
					return errors.New("invalid value type for text field validator")
				}
				if len(t.Value) > intValue {
					return errors.New(validator.CustomError)
				}
			}
		}

	}
	return nil
}
