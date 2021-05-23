package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type employee struct {
	Name string `validate:"required" valid:"required"`
}

func main() {
	e := employee{}
	err := validate1(e)
	if err != nil {
		fmt.Printf("Error Validate1: %s\n", err)
	}

	err = validate2(e)
	if err != nil {
		fmt.Printf("Error Validate2: %s\n", err)
	}
}

func validate1(e employee) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func validate2(e employee) error {
	_, err := govalidator.ValidateStruct(e)
	if err != nil {
		return err
	}
	return nil

}

