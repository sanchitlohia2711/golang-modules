package main

import (
	"encoding/json"
	"net/http"
)

var (
	employeeMap map[string]employee
)

type employee struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Photo string `json:"photo"`
}

func main() {
	employeeMap = make(map[string]employee)
	//Employee POST, PUT, PATCH, GET, DELETE, HEAD
	employeeHandler := http.HandlerFunc(employeeCRUD)
	http.Handle("/employee", employeeHandler)

	http.ListenAndServe(":8080", nil)
}

func postEmployee(w http.ResponseWriter, r *http.Request) {
	var e employee
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)
	employeeMap[e.Name] = e
	handleJSONResponse(w, "success", http.StatusOK)
	return
}

// func putEmployee(w http.ResponseWriter, r *http.Request) {
// 	var e employee
// 	decoder := json.NewDecoder(r.Body)
// 	decoder.Decode(&e)
// 	employeeMap[e.ID] = e
// 	w.WriteHeader(200)
// 	return
// }

func patchEmployee(w http.ResponseWriter, r *http.Request) {
	var e employee
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)
	currentEmployee, ok := employeeMap[e.Name]
	if ok {
		w.WriteHeader(404)
	}
	if e.Name != "" {
		currentEmployee.Name = e.Name
	}
	if e.Age != 0 {
		currentEmployee.Age = e.Age
	}
	w.WriteHeader(200)
	return
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	e, ok := employeeMap[name]
	if !ok {
		w.WriteHeader(404)
	}
	handleJSONResponse(w, e, 200)
	return
}

func headEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, ok := employeeMap[id]
	if ok {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
	return
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, ok := employeeMap[id]
	if ok {
		delete(employeeMap, id)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
	return
}

func employeeCRUD(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postEmployee(w, r)
		return
	} else if r.Method == "PATCH" {
		patchEmployee(w, r)
		return
	} else if r.Method == "GET" {
		getEmployee(w, r)
		return
	} else if r.Method == "HEAD" {
		headEmployee(w, r)
		return
	} else if r.Method == "DELETE" {
		deleteEmployee(w, r)
		return
	} else {
		w.WriteHeader(400)
	}
	return
}

func handleJSONResponse(w http.ResponseWriter, body interface{}, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	jsonResp, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(jsonResp)
}
