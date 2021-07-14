package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-examples/articles/tutorial/initFunction/simpleImport/redis"
)

func main() {
	err := redis.SetKey("a", "b", time.Minute*1)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	var value string
	err = redis.GetKey("a", &value)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	fmt.Println(value)

}
