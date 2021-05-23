// package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator"
// )

// var validatee *validator.Validate

// type employee struct {
// 	Name string `validate:"required"`
// }

// func main() {
// 	e := employee{}
// 	validate(e)
// }

// func validate(e employee) error {
// 	//v := validator.Validate{}
// 	validate = validatee.New()
// 	err := validate.Struct(e)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// 	fmt.Println("s")
// }
