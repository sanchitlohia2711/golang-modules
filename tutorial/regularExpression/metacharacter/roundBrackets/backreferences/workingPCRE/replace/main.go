package main

import (
	"fmt"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func main() {
	regex := pcre.MustCompile(`(\d)\1+`, 0)

	input := "The number is 91-88888888"

	result := regex.ReplaceAll([]byte(input), []byte("redacted"), 0)
	fmt.Println("result: ", string(result))
}
