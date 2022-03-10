package main

import "fmt"

func main() {
	output := largestRectangleArea([]int{3, 1, 5, 6, 2, 3})
	fmt.Println(output)

	output = largestRectangleArea([]int{2, 5, 4, 4, 4, 5, 5})
	fmt.Println(output)

	output = largestRectangleArea([]int{1, 1, 4})
	fmt.Println(output)

	output = largestRectangleArea([]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6})
	fmt.Println(output)
}

type customStack struct {
	stack []int
}

func (c *customStack) Push(num int) {
	c.stack = append(c.stack, num)
}

func (c *customStack) Pop() (int, error) {
	length := len(c.stack)
	poppedItem := 0
	if length > 0 {
		poppedItem = c.stack[length-1]
		c.stack = c.stack[:length-1]
		return poppedItem, nil
	}
	return 0, fmt.Errorf("stack is empty")
}

func (c *customStack) Front() (int, error) {
	length := len(c.stack)
	if length > 0 {
		return c.stack[length-1], nil
	}
	return 0, fmt.Errorf("stack is empty")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func largestRectangleArea(heights []int) int {
	customStack := &customStack{}

	lenHeights := len(heights)

	customStack.Push(0)

	maxRectangleSize := heights[0]

	for i := 1; i < lenHeights; i++ {

		for customStack.Size() != 0 {
			current, _ := customStack.Front()
			if heights[current] > heights[i] {
				var rectangleUsingCurrentBar int
				current, _ := customStack.Pop()
				//Calcualte max rectangle using the current front
				previous, err := customStack.Front()
				if err != nil {
					previous = -1
				}
				rectangleUsingCurrentBar = (i - previous - 1) * heights[current]
				if rectangleUsingCurrentBar > maxRectangleSize {
					maxRectangleSize = rectangleUsingCurrentBar
				}
			} else {
				break
			}
		}
		customStack.Push(i)
	}

	front, err := customStack.Front()
	if err != nil {
		return maxRectangleSize
	}
	var rectangleUsingCurrentBar int
	for customStack.Size() != 0 {
		current, _ := customStack.Pop()
		previous, err := customStack.Front()
		if err != nil {
			previous = -1
		}
		rectangleUsingCurrentBar = (front - previous) * heights[current]
		if rectangleUsingCurrentBar > maxRectangleSize {
			maxRectangleSize = rectangleUsingCurrentBar
		}
	}
	return maxRectangleSize
}
