package queue

import (
	"fmt"
	"sync"
)

type SQueue struct {
	queue []any
	lock  sync.RWMutex
}

func NewSStack() *SQueue {
	return &SQueue{
		queue: make([]any, 0),
	}
}

func (c *SQueue) Enqueue(value any) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue, value)
}

func (c *SQueue) Dequeue() error {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.queue = c.queue[1:]
		return nil
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *SQueue) Peek() any {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.queue[0]
	}
	return fmt.Errorf("peep Error: Queue is empty")
}

func (c *SQueue) Size() int {
	return len(c.queue)
}

func (c *SQueue) Empty() bool {
	return len(c.queue) == 0
}

func (c *SQueue) Display() {
	size := 0
	fmt.Print("Values stored in stack are: ")
	for size < len(c.queue) {
		fmt.Print(c.queue[size], " ")
		size = size + 1
	}
	fmt.Println()
}
