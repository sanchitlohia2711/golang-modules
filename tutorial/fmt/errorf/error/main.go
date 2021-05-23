package main

import "fmt"

func main() {
	name := "John"

	err := fmt.Errorf("Employee not found with name: %s and age %d", name)
	fmt.Println(err)
}
