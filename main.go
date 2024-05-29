package main

import (
	"fmt"
	"log"

	"github.com/DimRev/cli-from-v2/internal/form"
)

func main() {
	f := form.NewForm()
	i1, err := f.AddTextField("name", "Name")
	if err != nil {
		log.Fatalf("%v", err)
	}
	i1.AddValidator("too short", form.ValidatorTypes_Gt, "3")
	err = f.Render()
	if err != nil {
		log.Fatalf("%v", err)
	}
	for key, input := range f.Inputs {
		fmt.Printf("%v: %+v\n", key, input)
	}
}
