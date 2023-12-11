package main

import (
	"fmt"
	bst "github.com/munens/data-structures-algorithms/binary-search-tree"
)

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

	//qs := queue.NewQueueStack[string](10)
	//qs.Enqueue("jaymo")
	//qs.Enqueue("babu")
	//
	//qs.Lookup()
	//
	//qs.Dequeue()
	//qs.Peek()
	//qs.Lookup()

	t := bst.NewBinarySearchTree[int](5, func(v1 int, v2 int) bool {
		return v1 > v2
	})

	fmt.Println("inserting 7")
	if err := t.Insert(7); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 7")

	fmt.Println("inserting 6")
	if err := t.Insert(6); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 6")

	fmt.Println("inserting 3")
	if err := t.Insert(3); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 3")

	fmt.Println("inserting 4")
	if err := t.Insert(4); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 4")

	fmt.Println("inserting 10")
	if err := t.Insert(10); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 10")

	fmt.Println("inserting 9")
	if err := t.Insert(9); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 9")

	fmt.Println("inserting 11")
	if err := t.Insert(11); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("inserted 11")

	fmt.Println("looking up 5")
	if _, err := t.Lookup(5); err != nil {
		fmt.Println("unable to lookup")
	}
	fmt.Println("looked up 5")

	fmt.Println("looking up 7")
	if _, err := t.Lookup(7); err != nil {
		fmt.Println("unable to lookup")
	}
	fmt.Println("looked up 7")

	fmt.Println("looking up 6")
	if _, err := t.Lookup(6); err != nil {
		fmt.Println("unable to lookup")
	}
	fmt.Println("looked up 6")

	fmt.Println("looking up 4")
	if _, err := t.Lookup(4); err != nil {
		fmt.Println("unable to lookup")
	}
	fmt.Println("looked up 4")

	//           5
	//         /   \
	//       3      7
	//       \     /  \
	//       4    6   10
	//               /  \
	//              9    11

	fmt.Println("removing 4")
	if err := t.Remove(4); err != nil {
		fmt.Print("unable to remove")
	}
	fmt.Println("removed 4")

	fmt.Println("removing 7")
	if err := t.Remove(7); err != nil {
		fmt.Print("unable to remove")
	}
	fmt.Println("removed 7")

	fmt.Println("looking up 9")
	if _, err := t.Lookup(9); err != nil {
		fmt.Println("unable to lookup")
	}
	fmt.Println("looked up 9")
}
