package form

import (
	"bufio"
	"errors"
	"os"
)

type Form struct {
	Inputs map[string]*Input
	reader *bufio.Reader
}

type Input struct {
	Label      string
	Value      string
	Validators string
}

func NewForm() Form {
	reader := bufio.NewReader(os.Stdin)

	return Form{
		Inputs: make(map[string]*Input),
		reader: reader,
	}
}

func (f *Form) AddInput(id, label string) error {
	_, ok := f.Inputs[id]
	if ok {
		return errors.New("cannot have two form inputs of the same id")
	}
	f.Inputs[id] = &Input{
		Label: label,
	}
	return nil
}
