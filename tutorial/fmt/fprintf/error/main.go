package main

import (
	"fmt"
	"os"
)

type employee struct {
	Name string
	Age  int
}

func main() {
	name := "John"

	fmt.Fprintf(os.Stdout, "Name is: %s %d\n", name)
}
