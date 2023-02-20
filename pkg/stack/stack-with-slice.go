package stack

import (
	"fmt"
	"sync"
)

type SStack struct {
	stack []any
	lock  sync.RWMutex
}

func NewSStack() *SStack {
	return &SStack{
		stack: make([]any, 0),
	}
}

func (c *SStack) Push(value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, value)
}

func (c *SStack) Pop() error {
	size := len(c.stack)
	if size > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:size-1]
		return nil
	}
	return fmt.Errorf("pop Error: Stack is empty")
}

func (c *SStack) Peek() any {
	size := len(c.stack)
	if size > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[size-1]
	}
	return fmt.Errorf("peep Error: Stack is empty")
}

func (c *SStack) Size() int {
	return len(c.stack)
}

func (c *SStack) IsEmpty() bool {
	return len(c.stack) == 0
}

func (c *SStack) Display() {
	size := len(c.stack)
	fmt.Print("Values stored in stack are: ")
	for size > 0 {
		fmt.Print(c.stack[size-1], " ")
		size = size - 1
	}
	fmt.Println()
}
