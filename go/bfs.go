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
	q := Queue[*BinaryNode[T]]{}

	if tree == nil {
		return false
	}
	
	q.enqueue(tree)

	for q.length > 0 {
		curr := q.deque()

		if curr.val == nil {
			continue
		}

		if curr.val.value == needle {
			return true
		}

		q.enqueue(curr.val.left)
		q.enqueue(curr.val.right)
	}
	return false 
}

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

func (q *Queue[T]) deque() *QueueNode[T]{
	q.length--

	if q == nil {
		return nil
	}

	tmp := q.start
	q.start = q.start.next
	tmp.next = nil


	return tmp
}
