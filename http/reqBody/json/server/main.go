package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	createEmployeeHanlder := http.HandlerFunc(createEmployee)
	http.Handle("/employee", createEmployeeHanlder)
	http.ListenAndServe(":8080", nil)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e employee
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Success", http.StatusOK)
	return
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

//Correct Request
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Name":"John", "Age": 21}'

//There is an extra field in the incoming JSON.
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Name":"John", "Age": 21, "Gender": "M"}'

//The body is not a valid JSON
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Name": "John", "Age":}'

//The body is empty
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d ''

//The body is empty
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee

//The type of one of the field is different than expected. For example sending a string value of age where int is expected
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Name":"John", "Age": "21"}'

//The type of one of the field is different than expected. For example sending a string value of age where int is expected
//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Name":"John", "Age": "21"}'

//Content-type is not provided or not application/json
//curl -v -X POST http://localhost:8080/employee -d '{"Name":"John", "Age": 21}'
