package main

import (
	"fmt"
	"log"
)

type employee struct {
	Name string
	Age  int
}

func main() {
	name := "John"
	age := 21
	fmt.Println("Name is: ", name)
	fmt.Println("Age is: ", age)
	e := employee{
		Name: name,
		Age:  age,
	}
	fmt.Println(e)
	fmt.Println("a", 12, "b", 12.0)

	bytesPrinted, err := fmt.Println("Name is: ", name)
	if err != nil {
		log.Fatalln("Error occured", err)
	}
	fmt.Println(bytesPrinted)
}
