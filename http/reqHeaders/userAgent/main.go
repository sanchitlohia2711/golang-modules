package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	userAgent := r.UserAgent()
	fmt.Printf("UserAgent is: %s", userAgent)
}

//curl -v -X POST http://localhost:8080/example -H "user-agent: Mozialla -1.0"
