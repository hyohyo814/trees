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
	root := &BinaryNode[int]{value: 15}
	root.left = &BinaryNode[int]{value: 7}
	root.right = &BinaryNode[int]{value: 51}
	root.left.left = &BinaryNode[int]{value: 4}
	root.right.left = &BinaryNode[int]{value: 25}
	root.right.right = &BinaryNode[int]{value: 100}
	root.right.left.right = &BinaryNode[int]{value: 37}

	root.printTree("", true)
	target := 37
	res := root.search(target)
	fmt.Printf("Found match for %v: %v\n", target, res)
}

func (tree *BinaryNode[T]) search(target int) bool {
	if tree == nil {
		return false
	}

	if tree.value == target {
		return true
	}

	if target < tree.value {
		return tree.left.search(target)
	} else {
		return tree.right.search(target)
	}
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
