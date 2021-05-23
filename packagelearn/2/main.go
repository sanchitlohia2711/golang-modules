package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(subtract(2, 1))
}

func add(a, b int) int {
	return a + b
}
