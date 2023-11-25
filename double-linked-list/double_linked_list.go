package doublelinkedlist

import (
	"errors"
	"fmt"
)

type DoubleNode[T interface{}] struct {
	value T
	prev  *DoubleNode[T]
	next  *DoubleNode[T]
}

type DoubleLinkedList[T interface{}] struct {
	Head *DoubleNode[T]
	Tail *DoubleNode[T]
	Size int
}

func (l *DoubleLinkedList[T]) Add(value T) {
	node := &DoubleNode[T]{value: value, prev: nil, next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		node.prev = l.Tail
		l.Tail.next = node
		l.Tail = node
	}

	l.Size++
}

func (l *DoubleLinkedList[T]) insert(i int, value T) error {
	if i < 0 || i > l.Size {
		return errors.New("index is out of range")
	}

	if l.Head == nil || i == l.Size {
		l.Add(value)
		return nil
	}

	node := &DoubleNode[T]{value: value, prev: nil, next: nil}

	if i == 0 {
		node.next = l.Head
		l.Head.prev = node
		l.Head = node
		l.Size++
		return nil
	}

	curr := l.Head.next
	prev := l.Head
	currI := 1
	for curr != nil {
		if currI == i {
			node.prev = prev
			node.next = curr
			prev.next = node
			curr.prev = node
			l.Size++
			return nil
		}

		curr = curr.next
		prev = curr
		currI++
	}

	return nil
}

func (l *DoubleLinkedList[T]) Traverse() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (l *DoubleLinkedList[T]) Print() {
	s := ""
	curr := l.Head
	i := 0
	for curr != nil {
		backDir := ""
		if curr.prev != nil {
			backDir = "<-"
		}

		forDir := ""
		if curr.next != nil {
			forDir = "->"
		}

		v, _ := any(curr.value).(int)

		s = fmt.Sprintf("%s %s%d%s", s, backDir, v, forDir)

		i++
		curr = curr.next
	}

	fmt.Println(fmt.Sprintf("list: %s", s))
	fmt.Println(fmt.Sprintf("head -> %d", any(l.Head.value).(int)))
	fmt.Println(fmt.Sprintf("tail -> %d", any(l.Tail.value).(int)))
}

func (l *DoubleLinkedList[T]) TraverseReverse() {
	curr := l.Tail
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.prev
	}
}

func (l *DoubleLinkedList[T]) Reverse() {
	if l.Size < 2 {
		return
	}

	i := l.Size

	curr := l.Tail
	prev := l.Tail.prev
	prevCurr := l.Head

	for i > 1 {

		if i == l.Size {
			curr.next = l.Head
			curr.prev = nil
			prevCurr.prev = curr
			l.Head = curr
		} else {
			curr.prev = prevCurr
			curr.next = prevCurr.next
			prevCurr.next.prev = curr
			prevCurr.next = curr
		}

		l.Tail = prev
		prev.next = nil

		i--

		prevCurr = curr
		curr = l.Tail
		prev = l.Tail.prev
	}
}

func (l *DoubleLinkedList[T]) RemoveByIndex(i int) error {
	if l.Head == nil {
		return errors.New("DoubleLinkedList does not have nodes")
	}

	if i < 0 || i >= l.Size {
		return errors.New("Index is out of range")
	}

	if i == 0 {
		secNode := l.Head.next
		l.Head = secNode
		l.Size--

		if l.Size == 0 {
			l.Tail = nil
		} else {
			secNode.prev = nil
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
				next := curr.next
				prev.next = next
				next.prev = prev
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

func main() {
	linkedList := DoubleLinkedList[int]{}
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(10)

	linkedList.Print()

	linkedList.Reverse()

	linkedList.Print()

	linkedList.Reverse()

	linkedList.Print()
}
