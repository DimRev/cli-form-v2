package main

import (
	"fmt"
	"log"

	"github.com/DimRev/cli-from-v2/internal/form"
)

func main() {
	f := form.NewForm()
	firstName, err := f.AddInput("first name", "first name", form.FieldTypes_Text)
	if err != nil {
		log.Fatal("Failed to create a field: ", err)
	}
	firstName.AddValidator("Name is not Dima", form.ValidatorTypes_Eq, "Dima")

	lastName, err := f.AddInput("last name", "last name", form.FieldTypes_Text)
	if err != nil {
		log.Fatal("Failed to create a field: ", err)
	}
	lastName.AddValidator("Last name is not Rev", form.ValidatorTypes_Eq, "Rev")

	age, err := f.AddInput("age", "age", form.FieldTypes_Int)
	if err != nil {
		log.Fatal("Failed to create a field: ", err)
	}
	age.AddValidator("Age is less then 18", form.ValidatorTypes_Gt, 19)

	f.Render()

	for key, input := range f.Inputs {
		fmt.Printf("%v: %+v\n", key, input)
	}
}
