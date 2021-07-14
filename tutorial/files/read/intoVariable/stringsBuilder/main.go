package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatalf("Error while opening file. Err: %s", err)
	}
	defer file.Close()

	fileString := new(strings.Builder)
	io.Copy(fileString, file)

	fmt.Print(fileString)
}
