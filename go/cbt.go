package main

import "fmt"

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

func main()  {
	tree1 := &BinaryNode[int]{value: 5}
	tree1.left = &BinaryNode[int]{value: 3}
	tree1.right = &BinaryNode[int]{value: 1}
	tree1.left.left = &BinaryNode[int]{value: 7}
	tree1.left.right = &BinaryNode[int]{value: 9}
	tree1.right.left = &BinaryNode[int]{value: 14}
	tree1.printTree("", true)

	tree2 := &BinaryNode[int]{value: 5}
	tree2.left = &BinaryNode[int]{value: 3}
	tree2.right = &BinaryNode[int]{value: 1}
	tree2.left.left = &BinaryNode[int]{value: 7}
	tree2.left.right = &BinaryNode[int]{value: 9}
	tree2.printTree("", true)

	tree3 := &BinaryNode[int]{value: 5}
	tree3.left = &BinaryNode[int]{value: 3}
	tree3.left.left = &BinaryNode[int]{value: 7}
	tree3.printTree("", true)

	tree4 := &BinaryNode[int]{value: 5}
	tree4.left = &BinaryNode[int]{value: 3}
	tree4.right = &BinaryNode[int]{value: 1}
	tree4.left.left = &BinaryNode[int]{value: 7}
	tree4.left.right = &BinaryNode[int]{value: 9}
	tree4.right.left = &BinaryNode[int]{value: 14}
	tree4.printTree("", true)


	res1 := compare(tree1, tree2)
	res2 := compare(tree2, tree3)
	res3 := compare(tree1, tree4)	
	
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

func compare(a *BinaryNode[int], b *BinaryNode[int]) bool {
	if a == nil  && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
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
