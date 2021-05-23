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

	fmt.Fprint(os.Stdout, "Name is:", name, "\n")
	fmt.Fprint(os.Stdout, "Age is:", age, "\n")

	e := employee{
		Name: name,
		Age:  age,
	}

	fmt.Fprint(os.Stdout, e, "\n")

	fmt.Fprint(os.Stdout, "a", 12, "b", 12.0, "\n")

	fmt.Fprint(os.Stdout, 12, 12.0, "\n")

	bytesPrinted, err := fmt.Fprint(os.Stdout, "Name is: ", name, "\n")
	if err != nil {
		log.Fatalln("Error occured", err)
	}
	fmt.Println(bytesPrinted)
}
