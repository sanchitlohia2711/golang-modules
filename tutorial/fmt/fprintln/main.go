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
	fmt.Fprintln(os.Stdout, "Name is: ", name)
	fmt.Fprintln(os.Stdout, "Age is: ", age)
	e := employee{
		Name: name,
		Age:  age,
	}
	fmt.Fprintln(os.Stdout, e)
	fmt.Fprintln(os.Stdout, "a", 12, "b", 12.0)

	bytesPrinted, err := fmt.Fprintln(os.Stdout, "Name is: ", name)
	if err != nil {
		log.Fatalln("Error occured", err)
	}
	fmt.Println(bytesPrinted)
}
