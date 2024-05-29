package form

import (
	"bufio"
	"errors"
	"fmt"
)

type TextField struct {
	Label      string
	Value      string
	ValueType  FieldTypes
	Validators map[ValidatorTypes]Validator
	reader     *bufio.Reader
}

func (f *Form) AddTextField(id, label string) (*TextField, error) {
	_, ok := f.Inputs[id]
	if ok {
		return nil, errors.New("cannot have two form inputs of the same id")
	}
	field := &TextField{
		Label:      label,
		ValueType:  FieldTypes_Text,
		Validators: make(map[ValidatorTypes]Validator),
		reader:     f.reader,
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

func (t *TextField) Validate(preRender bool) error {
	for validatorType, validator := range t.Validators {
		switch validatorType {
		case ValidatorTypes_Gt:
			intValue, ok := validator.Value.(int)
			if !ok {
				return fmt.Errorf("invalid value type for field [%v]", t.Label)
			}
			if len(t.Value) <= intValue && !preRender {
				return errors.New(validator.CustomError)
			}

		case ValidatorTypes_Lt:
			intValue, ok := validator.Value.(int)
			if !ok {
				return fmt.Errorf("invalid value type for field [%v]", t.Label)
			}
			if len(t.Value) >= intValue && !preRender {
				return errors.New(validator.CustomError)
			}

		case ValidatorTypes_Eq:
			intValue, ok := validator.Value.(int)
			if !ok {
				return fmt.Errorf("invalid value type for field [%v]", t.Label)
			}
			if len(t.Value) != intValue && !preRender {
				return errors.New(validator.CustomError)
			}
		}
	}
	return nil
}

func (t *TextField) Render() error {
	err := t.Validate(true)
	if err != nil {
		return fmt.Errorf("form-definition error: %v", err)
	}
	fmt.Print(t.Label, ": ")
	val, err := t.reader.ReadString('\n')
	if err != nil {
		return err
	}
	err = t.Validate(false)
	if err != nil {
		return fmt.Errorf("user-input error: %v", err)
	}
	t.Value = val
	return nil
}
