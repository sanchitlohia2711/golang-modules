package main

import "fmt"

type employee struct {
	Name string
	Age  int
}

func main() {
	name := "John"

	fmt.Printf("Name is: %s %d\n", name)
}
