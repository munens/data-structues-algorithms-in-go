package main

import (
	doublelinkedlist "github.com/munens/data-structures-algorithms/double-linked-list"
	linkedlist "github.com/munens/data-structures-algorithms/linked-list"
)

func main() {

	dll := doublelinkedlist.DoubleLinkedList[int]{}
	dll.Add(1)
	dll.Add(2)
	dll.Add(3)
	dll.Add(4)
	dll.Add(10)

	dll.Print()

	dll.Reverse()

	dll.Print()

	dll.Reverse()

	dll.Print()

	ll := linkedlist.LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	ll.Print()

	ll.Reverse()

	ll.Print()
}
