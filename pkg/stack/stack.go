package stack

import "fmt"

type (
	Stack struct {
		top  *node
		size int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// New Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Size Return the number of items in the stack
func (s *Stack) Size() int {
	return s.size
}

// Peek View the top item on the stack
func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return fmt.Errorf("peep Error: Stack is empty")
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() error {
	if s.size == 0 {
		return fmt.Errorf("pop Error: Stack is empty")
	}

	n := s.top
	s.top = n.prev
	s.size--
	return nil
}

// Push a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.size++
}

// IsEmpty check empty stack
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Display() {
	temp := s.top
	fmt.Print("Values stored in stack are: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.prev
	}
	fmt.Println()
}
