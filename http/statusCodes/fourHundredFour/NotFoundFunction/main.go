package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/example", http.NotFound)
	http.ListenAndServe(":8080", nil)
}

//curl -v -X POST http://localhost:8080/example
