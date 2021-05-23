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
	fmt.Fprintln(file, "Name is: ", name)
	fmt.Fprintln(file, "Age is: ", age)
	e := employee{
		Name: name,
		Age:  age,
	}
	fmt.Fprintln(file, e)
	fmt.Fprintln(file, "a", 12, "b", 12.0)
}
