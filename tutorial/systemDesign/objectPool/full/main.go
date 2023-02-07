package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	inUse    int
	mulock   *sync.Mutex
}

//InitPool Initialize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot craete a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) loan() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	p.inUse = p.inUse + 1
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
	return obj, nil
}

func (p *pool) receive(target iPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	p.inUse = p.inUse - 1
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Targe pool object doesn't belong to the pool")
}

func (p *pool) setMaxCapacity(capacity int) {
	p.capacity = capacity
}

type iPoolObject interface {
	getID() string //This is any id which can be used to compare two different pool objects
	doSomething()
}

type dbconnection struct {
	id string
}

func (c *dbconnection) getID() string {
	return c.id
}

func (c *dbconnection) doSomething() {
	fmt.Printf("Connection with id %s in action\n", c.getID())
}

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

func main() {
	client := &client{}
	client.init()
	client.doWork()
}
