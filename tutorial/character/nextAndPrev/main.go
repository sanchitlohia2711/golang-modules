package main

import "fmt"

func main() {
	var curr byte
	curr = 'b'

	fmt.Println(string(curr + 1))
	fmt.Println(string(curr - 1))
}
