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
	Validate() error
}

type Validator struct {
	CustomError string
	Value       interface{}
}

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
