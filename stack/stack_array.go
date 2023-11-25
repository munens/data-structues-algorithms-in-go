package stack

import (
	"errors"
	"fmt"
)

type StackArray[T interface{}] struct {
	Size  int
	array []T
}

func New[T interface{}](size int) *StackArray[T] {
	return &StackArray[T]{
		Size:  size,
		array: make([]T, size),
	}
}

func (s *StackArray[T]) Lookup() {
	for v := range s.array {
		fmt.Println(v)
	}
}

func (s *StackArray[T]) Push(value T) error {
	if cap(s.array) == s.Size {
		return errors.New("stack is at capacity")
	}

	_ = append(s.array, value)
	s.Size += 1
	return nil
}

func (s *StackArray[T]) Pop() (*T, error) {
	if s.Size == 0 {
		return nil, errors.New("stack has no items")
	}

	v := &s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]

	s.Size -= 1
	return v, nil
}

func (s *StackArray[T]) Peak() error {
	if s.Size == 0 {
		return errors.New("stack has no items")
	}

	fmt.Print(s.array[len(s.array)-1])
	return nil
}
