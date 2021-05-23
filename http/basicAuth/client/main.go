package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	call("https://localhost:8080/example", "POST")
}

func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}

//curl -v -X POST http://localhost:8085/example -H "content-type: application/json" -H "foo: bar1" -H "foo: bar2"

//curl -v -X GET http://localhost:8085/employee -d '{"Name":"s"}'
