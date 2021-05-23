package main

import (
	"fmt"
	"log"
	"os"
)

type employee struct {
	Name string
	Age  int
}

func main() {
	name := "John"
	age := 21

	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(file, "Name is: %s\n", name)
	fmt.Fprintf(file, "Age is: %d\n", age)

	fmt.Fprintf(file, "Name: %s Age: %d\n", name, age)

	e := employee{
		Name: name,
		Age:  age,
	}

	fmt.Fprintf(file, "Employee is %v\n", e)
	fmt.Fprintf(file, "Employee is %+v\n", e)
	fmt.Fprintf(file, "Employee is %#v\n", e)
}
