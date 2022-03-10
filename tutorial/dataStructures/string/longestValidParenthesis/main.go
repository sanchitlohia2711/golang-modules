package main

import (
	"fmt"
	"sync"
)

func main() {
	valid := longestValidParentheses("()(()")
	fmt.Println(valid)

	valid = longestValidParentheses("()()())")
	fmt.Println(valid)

	valid = longestValidParentheses("(()")
	fmt.Println(valid)

	valid = longestValidParentheses("(()(()")
	fmt.Println(valid)

	valid = longestValidParentheses("()()")
	fmt.Println(valid)

	valid = longestValidParentheses("(()()")
	fmt.Println(valid)

	valid = longestValidParentheses("((()))(()")
	fmt.Println(valid)

	valid = longestValidParentheses(")(((((()())()()))()(()))(")
	fmt.Println(valid)

	valid = longestValidParentheses("")
	fmt.Println(valid)

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

func (c *customStack) Pop() (int, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		val := c.stack[len-1]
		c.stack = c.stack[:len-1]
		return val, nil
	}
	return 0, fmt.Errorf("Pop Error: Stack is empty")
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

func (c *customStack) Flush() {
	c.stack = []int{}
}

func longestValidParentheses(s string) int {

	customStack := &customStack{
		stack: []int{},
	}

	longestValidParenthesisLength := 0
	currentValidParenthesisLength := 0
	customStack.Push(-1)
	for i, val := range s {

		if val == '(' {
			customStack.Push(i)
		} else if val == ')' {
			customStack.Pop()
			if customStack.Size() == 0 {
				customStack.Push(i)
			} else {
				index, _ := customStack.Front()
				currentValidParenthesisLength = i - index
				if currentValidParenthesisLength > longestValidParenthesisLength {
					longestValidParenthesisLength = currentValidParenthesisLength
				}
			}

		}

	}

	if currentValidParenthesisLength > longestValidParenthesisLength {
		longestValidParenthesisLength = currentValidParenthesisLength
	}

	return longestValidParenthesisLength
}
