package main

import "fmt"

type class struct {
	className string
	students  []student
}

type student struct {
	name   string
	rollNo int
	city   string
}

func main() {
	goerge := student{"Goerge", 35, "Newyork"}
	john := student{"Goerge", 25, "London"}

	students := []student{goerge, john}
	class := class{"firstA", students}

	fmt.Printf("class is %v\n", class)
}
