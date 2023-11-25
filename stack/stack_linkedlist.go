package stack

import (
	"errors"
	"fmt"
	linkedlist "github.com/munens/data-structures-algorithms/linked-list"
)

type Stack[T interface{}] struct {
	Size       int
	linkedlist linkedlist.LinkedList[T]
}

func (s *Stack[T]) Lookup() error {
	if s.linkedlist.Size == 0 {
		return errors.New("stack has no items")
	}

	s.linkedlist.Traverse()
	return nil
}

func (s *Stack[T]) Push(value T) error {
	if err := s.linkedlist.Insert(0, value); err != nil {
		return err
	}

	s.Size += 1
	return nil
}

func (s *Stack[T]) Pop() (*T, error) {
	v := &s.linkedlist.Head.Value
	if err := s.linkedlist.RemoveByIndex(0); err != nil {
		return nil, err
	}

	s.Size -= 1
	return v, nil
}

func (s *Stack[T]) Peak() error {
	if s.linkedlist.Size == 0 {
		return errors.New("stack has no items")
	}

	fmt.Print(s.linkedlist.Head.Value)
	return nil
}
