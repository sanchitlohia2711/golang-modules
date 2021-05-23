package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string
}

func main() {
	a := make(map[string]employee)

	a["1"] = employee{Name: "John"}

	j, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}
}
