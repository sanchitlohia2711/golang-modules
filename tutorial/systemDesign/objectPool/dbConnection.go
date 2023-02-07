package main

import "fmt"

type dbconnection struct {
	id string
}

func (c *dbconnection) getID() string {
	return c.id
}

func (c *dbconnection) doSomething() {
	fmt.Printf("Connection with id %s in action\n", c.getID())
}
