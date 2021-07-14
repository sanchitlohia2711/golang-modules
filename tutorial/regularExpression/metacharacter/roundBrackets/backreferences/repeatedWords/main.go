package main

import (
	"fmt"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func main() {
	regex := pcre.MustCompile(`\b(\w+)\s\1\b`, 0)

	input := "The name name is John John"

	result := regex.ReplaceAll([]byte(input), []byte(`$1`), 0)
	fmt.Println("result: ", string(result))
}
