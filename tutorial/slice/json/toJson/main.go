package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := make([]string, 2)

	a[0] = "John"
	a[1] = "Sam"

	j, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}
}
