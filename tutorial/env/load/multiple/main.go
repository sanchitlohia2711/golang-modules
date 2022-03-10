package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env", "test.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("STACK")
	fmt.Println(val)

	val = os.Getenv("DATABASE")
	fmt.Println(val)

	val = os.Getenv("TEST")
	fmt.Println(val)
}
