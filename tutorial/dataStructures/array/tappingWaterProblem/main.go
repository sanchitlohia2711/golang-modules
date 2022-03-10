package main

import (
	"fmt"
	"sync"
)

func main() {
	output := trap([]int{2, 0, 2, 1, 3, 1})
	fmt.Println(output)

	output = trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(output)

	output = trap([]int{4, 2, 0, 3, 2, 5})
	fmt.Println(output)
}

func trap(height []int) int {

	output := 0
	heightArrayLen := len(height)
	customStack := customStack{
		stack: make([]int, 0),
	}

	customStack.Push(0)

	for i := 1; i < heightArrayLen; i++ {

		for customStack.Size() != 0 {
			front, _ := customStack.Front()
			if height[front] <= height[i] {
				output = output + (i-front-1)*(height[front]-max(height, front+1, i-1))
				customStack.Pop()
			} else {
				output = output + (i-front-1)*(height[i]-max(height, front+1, i-1))
				break
			}
		}
		customStack.Push(i)
	}
	return output

}

func max(input []int, start, end int) int {

	if start > end {
		return 0
	}

	max := 0

	for i := start; i <= end; i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

type customStack struct {
	stack []int
	lock  sync.RWMutex
}

func (c *customStack) Push(name int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStack) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (int, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	return 0, fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	return len(c.stack) == 0
}
