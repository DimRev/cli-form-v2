package form

import (
	"bufio"
	"errors"
	"os"
)

type Form struct {
	Inputs map[string]*Field
	reader *bufio.Reader
}

type Field struct {
	Label      string
	Value      string
	Type       FieldTypes
	Validators map[ValidatorTypes]struct {
		CustomError string
		Value       any
	}
}

type FieldTypes string

const (
	FieldTypes_Text  FieldTypes = "text"
	FieldTypes_Int   FieldTypes = "int"
	FieldTypes_Float FieldTypes = "float"
	FieldTypes_Bool  FieldTypes = "bool"
)

type ValidatorTypes string

const (
	ValidatorTypes_Lt ValidatorTypes = "LT"
	ValidatorTypes_Gt ValidatorTypes = "GT"
	ValidatorTypes_Eq ValidatorTypes = "EQ"
)

func NewForm() Form {
	reader := bufio.NewReader(os.Stdin)

	return Form{
		Inputs: make(map[string]*Field),
		reader: reader,
	}
}

func (f *Form) AddInput(id, label string, valueType FieldTypes) (*Field, error) {
	_, ok := f.Inputs[id]
	if ok {
		return &Field{}, errors.New("cannot have two form inputs of the same id")
	}
	field := Field{
		Label: label,
		Type:  valueType,
		Validators: make(map[ValidatorTypes]struct {
			CustomError string
			Value       any
		}),
	}
	f.Inputs[id] = &field
	return &field, nil
}

func (f *Field) AddValidator(errMsg string, validatorType ValidatorTypes, value any) error {
	_, ok := f.Validators[validatorType]
	if ok {
		return errors.New("cannot have two validators of the same type")
	}
	f.Validators[validatorType] = struct {
		CustomError string
		Value       any
	}{
		CustomError: errMsg,
		Value:       value,
	}
	return nil
}
