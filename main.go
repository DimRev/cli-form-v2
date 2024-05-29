package main

import (
	"fmt"

	"github.com/DimRev/cli-from-v2/internal/form"
)

func main() {
	f := form.NewForm()
	f.AddInput("first name", "first name")
	f.AddInput("last name", "last name")
	f.AddInput("age", "age")
	f.Render()

	for key, input := range f.Inputs {
		fmt.Printf("%v: %+v\n", key, input)
	}
}
