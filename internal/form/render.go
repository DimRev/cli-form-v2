package form

import "fmt"

func (f *Form) Render() error {
	for name, input := range f.Inputs {
		fmt.Printf("%s: ", input.Label)
		text, err := f.reader.ReadString('\n')
		if err != nil {
			return err
		}
		f.Inputs[name].Value = text
	}
	return nil
}
