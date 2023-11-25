package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

type Node[T interface{}] struct {
	Value T
	next  *Node[T]
}

type LinkedList[T interface{}] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{Value: value, next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.next = node
		l.Tail = node
	}

	l.Size += 1
}

func (l *LinkedList[T]) Insert(idx int, value T) error {
	if idx < 0 || idx > l.Size {
		return errors.New("index is out of range")
	}

	if idx == l.Size {
		l.Add(value)
		return nil
	}

	node := &Node[T]{Value: value}

	if idx == 0 {
		node.next = l.Head
		l.Head = node
		l.Size += 1
		return nil
	}

	curr := l.Head
	for i := 1; i < idx-1; i++ {
		if i == idx {
			node.next = curr.next
			curr.next = node
			l.Size += 1
			return nil
		}

		curr.next = curr
	}

	return nil
}

func (l *LinkedList[T]) Traverse() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.next
	}
}

func isEqual[T interface{}](v1 T, v2 T) bool {
	return reflect.ValueOf(v1).Interface() == reflect.ValueOf(v2).Interface()
}

func (l *LinkedList[T]) Print() {
	s := ""
	curr := l.Head
	i := 0
	for curr != nil {
		dir := ""
		if curr.next != nil {
			dir = "->"
		}

		v, _ := any(curr.Value).(int)

		s = fmt.Sprintf("%s %d%s", s, v, dir)

		i++
		curr = curr.next
	}

	fmt.Println(fmt.Sprintf("list: %s", s))
	fmt.Println(fmt.Sprintf("head -> %d", any(l.Head.Value).(int)))
	fmt.Println(fmt.Sprintf("tail -> %d", any(l.Tail.Value).(int)))
}

func (l *LinkedList[T]) Reverse() {
	if l.Size < 2 {
		return
	}

	i := 0

	curr := l.Head
	next := l.Head.next
	firstTail := l.Tail
	prevCurr := curr

	for i < l.Size-1 {

		firstTail.next = curr
		if i == 0 {
			curr.next = nil
			l.Tail = curr
		} else {
			curr.next = prevCurr
		}

		l.Head = next

		i++

		prevCurr = curr
		curr = l.Head
		next = l.Head.next
	}
}

func (l *LinkedList[T]) RemoveByIndex(i int) error {
	if l.Head == nil {
		return errors.New("LinkedList does not have nodes")
	}

	if i < 0 || i >= l.Size {
		return errors.New("index is out of range")
	}

	if i == 0 {
		l.Head = l.Head.next
		l.Size--

		if l.Size == 0 {
			l.Tail = nil
		}

		return nil
	}

	curr := l.Head.next
	prev := l.Head
	currI := 1
	for curr != nil {
		if i == currI {
			if currI == l.Size-1 {
				prev.next = nil
				l.Tail = prev
			} else {
				prev.next = curr.next
			}

			l.Size--
			return nil
		}

		currI++
		prev = curr
		curr = curr.next
	}

	return nil
}

func (l *LinkedList[T]) RemoveByVal(value T) error {
	if l.Head == nil {
		return errors.New("LinkedList does not have nodes")
	}

	if isEqual(l.Head.Value, value) {
		l.Head = l.Head.next
		l.Size--

		if l.Size == 0 {
			l.Tail = nil
		}

		return nil
	}

	curr := l.Head.next
	prev := l.Head
	for curr != nil {
		if isEqual(curr.Value, value) {

			if isEqual(l.Tail.Value, curr.Value) {
				prev.next = nil
				l.Tail = prev
			} else {
				prev.next = curr.next
			}

			l.Size--
			return nil
		}

		curr = curr.next
		prev = curr
	}

	return nil
}
