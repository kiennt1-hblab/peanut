package stack

import (
	"container/list"
	"fmt"
)

type LStack struct {
	stack *list.List
}

func NewLStack() *LStack {
	return &LStack{
		stack: list.New(),
	}
}

func (c *LStack) Push(value any) {
	c.stack.PushFront(value)
}

func (c *LStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("pop Error: Stack is empty")
}

func (c *LStack) Peek() any {
	if c.stack.Len() > 0 {
		return c.stack.Front().Value
	}
	return fmt.Errorf("peep Error: Stack is empty")
}

func (c *LStack) Size() int {
	return c.stack.Len()
}

func (c *LStack) IsEmpty() bool {
	return c.stack.Len() == 0
}

func (c *LStack) Display() {
	fmt.Print("Values stored in stack are: ")
	front := c.stack.Front()
	for front != nil {
		fmt.Print(front.Value, " ")
		front = front.Next()
	}
	fmt.Println()
}
