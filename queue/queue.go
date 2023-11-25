package queue

import (
	"errors"
	"fmt"
	linkedlist "github.com/munens/data-structures-algorithms/linked-list"
)

type Queue[T interface{}] struct {
	Size       int
	linkedlist linkedlist.LinkedList[T]
}

func (q *Queue[T]) Lookup() error {
	if q.linkedlist.Size == 0 {
		return errors.New("queue has no items")
	}

	q.linkedlist.Traverse()
	return nil
}

func (q *Queue[T]) Enqueue(value T) {
	q.linkedlist.Add(value)
	q.Size += 1
}

func (q *Queue[T]) Dequeue() (*T, error) {
	v := &q.linkedlist.Head.Value
	if err := q.linkedlist.RemoveByIndex(0); err != nil {
		return nil, err
	}

	q.Size -= 1
	return v, nil
}

func (q *Queue[T]) Peek() error {
	if q.linkedlist.Size == 0 {
		return errors.New("queue has no items")
	}

	fmt.Print(q.linkedlist.Head.Value)
	return nil
}
