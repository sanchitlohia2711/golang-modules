package main

import (
	"log"
	"os"
)

func main() {
	//Create a folder/directory at a full qualified path
	err := os.Mkdir("/Users/temp", 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Create a folder/directory at a full qualified path
	err = os.Mkdir("temp", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
