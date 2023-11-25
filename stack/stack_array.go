package stack

import (
	"errors"
	"fmt"
	"reflect"
)

type StackArray[T interface{}] struct {
	Size  int
	array []T
}

func New[T interface{}](size int) *StackArray[T] {
	return &StackArray[T]{
		Size:  0,
		array: make([]T, size),
	}
}

func (s *StackArray[T]) Lookup() {
	for _, v := range s.array {
		if !reflect.ValueOf(v).IsZero() {
			fmt.Println(v)
		}
	}
}

func (s *StackArray[T]) Push(value T) error {
	s.array = append(s.array, value)
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

func (s *StackArray[T]) Peek() error {
	if s.Size == 0 {
		return errors.New("stack has no items")
	}

	fmt.Print(s.array[len(s.array)-1])
	return nil
}
