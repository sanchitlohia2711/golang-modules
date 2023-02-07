package main

import (
	"fmt"
	"log"
	"strconv"
)

type client struct {
	pool *pool
}

func (c *client) init() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &dbconnection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	var err error
	c.pool, err = initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
}

func (c *client) doWork() {
	fmt.Printf("Capacity: %d\n\n", c.pool.capacity)

	conn1, err := c.pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn1.doSomething()

	fmt.Printf("InUse: %d\n\n", c.pool.inUse)
	conn2, err := c.pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2.doSomething()
	fmt.Printf("InUse: %d\n\n", c.pool.inUse)

	c.pool.receive(conn1)
	fmt.Printf("InUse: %d\n\n", c.pool.inUse)

	c.pool.receive(conn2)
	fmt.Printf("InUse: %d\n", c.pool.inUse)
}
