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

	fmt.Fprintf(os.Stdout, "Name is: %s\n", name)
	fmt.Fprintf(os.Stdout, "Age is: %d\n", age)

	fmt.Fprintf(os.Stdout, "Name: %s Age: %d\n", name, age)

	e := employee{
		Name: name,
		Age:  age,
	}

	fmt.Fprintf(os.Stdout, "Employee is %v\n", e)
	fmt.Fprintf(os.Stdout, "Employee is %+v\n", e)
	fmt.Fprintf(os.Stdout, "Employee is %#v\n", e)

	bytesPrinted, err := fmt.Fprintf(os.Stdout, "Name is: %s\n", name)
	if err != nil {
		log.Fatalln("Error occured", err)
	}
	fmt.Println(bytesPrinted)
}
