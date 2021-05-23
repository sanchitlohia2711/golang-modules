package main

import (
	"fmt"

	"sample.com/learn/model"
)

func main() {
	p := &model.Person{
		Name: "John",
	}
	fmt.Println(p.Name)
}
