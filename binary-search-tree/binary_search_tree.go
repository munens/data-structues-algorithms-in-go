package binary_search_tree

import (
	"errors"
	"fmt"
	"reflect"
)

type Node[T interface{}] struct {
	Value  T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

type compare[T interface{}] func(v1, v2 T) bool

type BinarySearchTree[T interface{}] struct {
	Root    *Node[T]
	compare compare[T]
}

// NewBinarySearchTree returns a binary tree with a single node (the root). Takes in a value, v and a compare func as arguments.
// The value, v will be used to set a value in the root node. And the compare function, c will be used to determine whether new inserted
// nodes should appear on the left/right. Right if the compare function returns true; left otherwise.
func NewBinarySearchTree[T interface{}](v T, c compare[T]) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		Root: &Node[T]{
			Value: v,
		},
		compare: c,
	}
}

func (bst *BinarySearchTree[T]) isEqual(t1, t2 T) bool {
	return reflect.ValueOf(t1).Equal(reflect.ValueOf(t2))
}

// traverseTo takes a value, v and a pointer to a node, n. If n is nil traverseTo will return the root node of the tree
// and traverse the tree until it finds a node in the tree with value, v. If it cant find a node with value, v then traverseTo
// will return the last node it traversed to.
func (bst *BinarySearchTree[T]) traverseTo(v T, n *Node[T]) *Node[T] {
	if n == nil {
		return bst.traverseTo(v, bst.Root)
	}

	fmt.Println(v, n, n.Left, n.Right)

	if bst.isEqual(n.Value, v) {
		return n
	}

	if bst.hasNoChildren(n) {
		return n
	}

	if bst.hasLeftOnly(n) {
		if bst.compare(v, n.Value) {
			return n
		} else {
			return bst.traverseTo(v, n.Left)
		}
	}

	if bst.hasRightOnly(n) {
		if bst.compare(v, n.Value) {
			return bst.traverseTo(v, n.Right)
		} else {
			return n
		}
	}

	if bst.compare(v, n.Value) {
		return bst.traverseTo(v, n.Right)
	} else {
		return bst.traverseTo(v, n.Left)
	}
}

// Insert adds a new node to the binary search tree with value, v. It will return an error if a node with a value, v
// already exists.
func (bst *BinarySearchTree[T]) Insert(v T) error {
	b := bst.traverseTo(v, nil)
	if bst.isEqual(b.Value, v) {
		return errors.New(fmt.Sprintf("Value, %s already exists in binary search tree", v))
	}

	n := &Node[T]{
		Value:  v,
		Parent: b,
	}

	if bst.compare(v, b.Value) {
		b.Right = n
	} else {
		b.Left = n
	}

	return nil
}

// Lookup looks up a value, v in the binary search tree. It will return an error if a value cant be found.
func (bst *BinarySearchTree[T]) Lookup(v T) (*Node[T], error) {
	if n := bst.traverseTo(v, nil); n == nil {
		return nil, errors.New(fmt.Sprintf("Unable to find value, %s in binary search tree", v))
	} else {
		return n, nil
	}
}

// Remove takes a value, v and removes the node related to it from the tree. If it is unable to find a node that has
// value, v it returns an error.
func (bst *BinarySearchTree[T]) Remove(v T) error {
	n := bst.traverseTo(v, nil)
	if n == nil {
		return errors.New(fmt.Sprintf("Unable to find value, %s in binary search tree", v))
	}

	if bst.hasNoChildren(n) {
		bst.updateChild(n.Parent, n.Value, nil)
	}

	if bst.hasLeftOnly(n) {
		bst.updateChild(n.Parent, n.Value, n.Left)
	}

	if bst.hasRightOnly(n) {
		bst.updateChild(n.Parent, n.Value, n.Right)
	}

	if bst.hasBothChildren(n) {
		succ := bst.findInOrderSuccessorD(n)
		succ.Parent = n.Parent
		succ.Left = n.Left
		succ.Right = n.Right

		bst.updateChild(n.Parent, n.Value, succ)
	}

	return nil
}

func (bst *BinarySearchTree[T]) hasBothChildren(n *Node[T]) bool {
	return n.Left != nil && n.Right != nil
}

func (bst *BinarySearchTree[T]) hasNoChildren(n *Node[T]) bool {
	return n.Left == nil && n.Right == nil
}

func (bst *BinarySearchTree[T]) hasLeftOnly(n *Node[T]) bool {
	return n.Left != nil && n.Right == nil
}

func (bst *BinarySearchTree[T]) hasRightOnly(n *Node[T]) bool {
	return n.Left == nil && n.Right != nil
}

func (bst *BinarySearchTree[T]) hasChild(n *Node[T]) bool {
	return bst.hasLeftOnly(n) || bst.hasRightOnly(n)
}

func (bst *BinarySearchTree[T]) updateChild(n *Node[T], v T, c *Node[T]) {
	if n.Left != nil {
		if bst.isEqual(n.Left.Value, v) {
			n.Left = c
		}
	}

	if n.Right != nil {
		if bst.isEqual(n.Right.Value, v) {
			n.Right = c
		}
	}
}

func (bst *BinarySearchTree[T]) findInOrderSuccessorD(n *Node[T]) *Node[T] {
	fmt.Println("findInOrderSuccessorD", n)
	if bst.hasNoChildren(n) {
		return n
	} else {

		if bst.hasBothChildren(n) {
			return bst.minValue(n, 0)
		}

		if bst.hasLeftOnly(n) {
			return bst.findInOrderSuccessorD(n.Left)
		} else {
			return bst.findInOrderSuccessorD(n.Right)
		}
	}
}

func (bst *BinarySearchTree[T]) minValue(n *Node[T], i int) *Node[T] {
	if bst.hasNoChildren(n) {
		return n
	} else {
		if i == 0 {
			return bst.minValue(n.Right, i+1)
		} else {
			if bst.hasBothChildren(n) {
				return bst.minValue(n.Left, i)
			}

			if bst.hasLeftOnly(n) {
				return bst.minValue(n.Left, i)
			} else {
				return bst.minValue(n.Right, i)
			}
		}
	}
}
