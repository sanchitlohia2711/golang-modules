package main

import "fmt"

type employee struct {
	name string
	age  int
}

func main() {
	employees := make([]employee, 3)

	employees[0] = employee{name: "John", age: 21}
	employees[1] = employee{name: "Simon", age: 25}
	employees[2] = employee{name: "David", age: 18}

	for _, num := range employees {
		fmt.Println(num)
	}
}
