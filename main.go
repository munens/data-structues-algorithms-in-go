package main

import "github.com/munens/data-structures-algorithms/queue"

func main() {

	//dll := doublelinkedlist.DoubleLinkedList[int]{}
	//dll.Add(1)
	//dll.Add(2)
	//dll.Add(3)
	//dll.Add(4)
	//dll.Add(10)
	//
	//dll.Print()
	//
	//dll.Reverse()
	//
	//dll.Print()
	//
	//dll.Reverse()
	//
	//dll.Print()
	//
	//ll := linkedlist.LinkedList[int]{}
	//ll.Add(1)
	//ll.Add(2)
	//ll.Add(3)
	//
	//ll.Print()
	//
	//ll.Reverse()
	//
	//ll.Print()

	//s := stack.Stack[string]{}
	//s.Push("google")
	//s.Push("udemy")
	//
	//s.Lookup()
	//s.Pop()
	//s.Lookup()

	//s := stack.NewStackArray[string](5)
	//s.Push("google")
	//s.Push("udemy")
	//
	//s.Lookup()
	//s.Pop()
	//s.Lookup()

	qs := queue.NewQueueStack[string](10)
	qs.Enqueue("jaymo")
	qs.Enqueue("babu")

	qs.Lookup()

	qs.Dequeue()
	qs.Peek()
	qs.Lookup()
}
