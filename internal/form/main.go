package form

import (
	"bufio"
	"os"
)

type Form struct {
	Inputs map[string]Field
	reader *bufio.Reader
}

type Field interface {
	AddValidator(errMsg string, validatorType ValidatorTypes, value interface{}) error
	Validate(preRender bool) error
	Render() error
}

type Validator struct {
	CustomError string
	Value       interface{}
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

func NewForm() *Form {
	reader := bufio.NewReader(os.Stdin)
	return &Form{
		Inputs: make(map[string]Field),
		reader: reader,
	}
}

func (f *Form) Render() error {
	for _, input := range f.Inputs {
		err := input.Render()
		if err != nil {
			return err
		}
	}
	return nil
}
