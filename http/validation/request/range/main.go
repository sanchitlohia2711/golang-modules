package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type employee struct {
	Age int `validate:"required,gte=10,lte=20"`
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

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}
	err = validateStruct(e)
	if err != nil {
		errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
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

func validateStruct(e employee) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Age": 5}'

//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Age": 21}'

//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Age": 15}'

//curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d ''
