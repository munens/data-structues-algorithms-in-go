package queue

import (
	"errors"
	"github.com/munens/data-structures-algorithms/stack"
)

type QueueStack[T interface{}] struct {
	s1   *stack.StackArray[T]
	s2   *stack.StackArray[T]
	Size int
}

func NewQueueStack[T interface{}](size int) *QueueStack[T] {
	return &QueueStack[T]{
		s1:   stack.NewStackArray[T](size),
		s2:   stack.NewStackArray[T](size),
		Size: 0,
	}
}

func (qs *QueueStack[T]) Lookup() error {
	if qs.s1.Size == 0 {
		return errors.New("queue stack has no items")
	}

	qs.s1.Lookup()
	return nil
}

func (qs *QueueStack[T]) Enqueue(value T) error {

	if qs.s1.Size > 0 {
		for i := 0; i < qs.s1.Size; i++ {
			v, err := qs.s1.Pop()

			if err != nil {
				return err
			}

			if err := qs.s2.Push(*v); err != nil {
				return err
			} else {
				qs.Size = qs.s1.Size
			}
		}
	}

	if err := qs.s1.Push(value); err != nil {
		return err
	}

	if qs.s2.Size > 0 {
		for i := 0; i < qs.s2.Size; i++ {
			v, err := qs.s2.Pop()

			if err != nil {
				return err
			}

			if err := qs.s1.Push(*v); err != nil {
				return err
			} else {
				qs.Size = qs.s1.Size
			}
		}
	}

	return nil
}

func (qs *QueueStack[T]) Dequeue() (*T, error) {
	if v, err := qs.s1.Pop(); err != nil {
		return nil, err
	} else {
		qs.Size = qs.s1.Size
		return v, nil
	}
}

func (qs *QueueStack[T]) Peek() error {
	if qs.s1.Size == 0 {
		return errors.New("queue stack has no items")
	}

	return qs.s1.Peek()
}

func (qs *QueueStack[T]) IsEmpty() bool {
	return qs.Size == 0
}
