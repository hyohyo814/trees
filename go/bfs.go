package main

import (
	"fmt"
)

type BinaryNode[T any] struct {
	left, right *BinaryNode[T]
	value       int
}

type Queue[T any] struct {
	start, end *QueueNode[T]
	length     int
}

type QueueNode[T any] struct {
	next *QueueNode[T]
	val  T
}

func main() {
	root := &BinaryNode[int]{value: 7}
	root.left = &BinaryNode[int]{value: 23}
	root.right = &BinaryNode[int]{value: 8}
	root.left.left = &BinaryNode[int]{value: 5}
	root.left.right = &BinaryNode[int]{value: 4}
	root.right.left = &BinaryNode[int]{value: 21}
	root.right.right = &BinaryNode[int]{value: 15}

	fmt.Println("tree display:")
	root.printTree("", true)

	target := 12
	res := root.bfs(target)
	fmt.Printf("bfs result for %v: %v\n", target, res)
}

func (tree *BinaryNode[T]) bfs(needle int) bool {
	// initialize Queue containing BinaryNodes
	q := Queue[*BinaryNode[T]]{}

	// exit if Node does not exist
	if tree == nil {
		return false
	}

	// insert root of tree into Queue
	q.enqueue(tree)

	for q.length > 0 {
		// Deque for search operation
		curr := q.deque()

		if curr.val == nil {
			// continue dereferencing nil pointers
			continue
		}

		if curr.val.value == needle {
			// match found
			return true
		}

		// enqueue children to continue search operation
		q.enqueue(curr.val.left)
		q.enqueue(curr.val.right)
	}

	// exit if no match found
	return false
}

// Method to visualize tree connections
func (tree *BinaryNode[T]) printTree(prefix string, isLeft bool) {
	if tree == nil {
		return
	}

	if isLeft {
		fmt.Printf("%s|--%d\n", prefix, tree.value)
	} else {
		fmt.Printf("%s└──%d\n", prefix, tree.value)
	}

	childPrefix := prefix
	if isLeft {
		childPrefix += "|  "
	} else {
		childPrefix += "   "
	}

	tree.left.printTree(childPrefix, true)

	tree.right.printTree(childPrefix, false)

}

// Method to get all values in tree in a slice
func (q *Queue[T]) getAll() []T {
	var elems []T
	for e := q.start; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (q *Queue[T]) enqueue(el T) {
	if q.length == 0 {
		q.start = &QueueNode[T]{val: el}
		q.end = q.start
	} else {
		q.end.next = &QueueNode[T]{val: el}
		q.end = q.end.next
	}
	q.length++
}

func (q *Queue[T]) deque() *QueueNode[T] {
	q.length--

	if q == nil {
		return nil
	}

	tmp := q.start
	q.start = q.start.next
	tmp.next = nil

	return tmp
}
