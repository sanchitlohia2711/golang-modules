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

	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	name := "John"
	age := 21

	fmt.Fprint(file, "Name is:", name, "\n")
	fmt.Fprint(file, "Age is:", age, "\n")

	e := employee{
		Name: name,
		Age:  age,
	}

	fmt.Fprint(file, e, "\n")

	fmt.Fprint(file, "a", 12, "b", 12.0, "\n")

	fmt.Fprint(file, 12, 12.0, "\n")
}
