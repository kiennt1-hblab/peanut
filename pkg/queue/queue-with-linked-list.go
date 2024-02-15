package queue

import (
	"container/list"
	"fmt"
)

type LQueue struct {
	queue *list.List
}

func NewLQueue() *LQueue {
	return &LQueue{
		queue: list.New(),
	}
}

func (c *LQueue) Enqueue(value any) {
	c.queue.PushBack(value)
}

func (c *LQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *LQueue) Peek() (any, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (c *LQueue) Size() int {
	return c.queue.Len()
}

func (c *LQueue) IsEmpty() bool {
	return c.queue.Len() == 0
}

func (c *LQueue) Display() {
	fmt.Print("Values stored in queue are: ")
	front := c.queue.Front()
	for front != nil {
		fmt.Print(front.Value, " ")
		front = front.Next()
	}
	fmt.Println()
}
