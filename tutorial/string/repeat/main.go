package main

import (
	"fmt"
	"strings"
)

func main() {
	copy := strings.Repeat("a", 2)
	fmt.Println(copy)

	copy = strings.Repeat("abc", 3)
	fmt.Println(copy)
}
