package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	handler := http.HandlerFunc(handleRequest)
// 	http.Handle("/example", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	headers := r.Header
// 	fmt.Println(headers)

// 	fmt.Println(headers["Content-Type"])
// 	fmt.Println(headers["Foo"])

// 	fmt.Println(headers["content-type"])
// 	fmt.Println(headers["foo"])
// }

// //curl -v -X POST http://localhost:8080/example -H "content-type: application/json"
