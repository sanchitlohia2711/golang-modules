package main

import (
	"fmt"
)

type employee struct {
	Name string
	Age  int
}

func main() {
	sampleErr := "database connection issue"

	err := fmt.Errorf("Err is: %s", sampleErr)
	fmt.Println(err)

	port := 8080

	err = fmt.Errorf("Err is: %s to port %d", sampleErr, port)
	fmt.Println(err)

	e := employee{
		Name: "John",
	}

	err = fmt.Errorf("Employee not found. Details %v", e)
	fmt.Println(err)
	err = fmt.Errorf("Employee not found. Details %+v", e)
	fmt.Println(err)
	err = fmt.Errorf("Employee not found. Details %#v", e)
	fmt.Println(err)
}
