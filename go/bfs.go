package main

import (
	"fmt"
)

type BinaryNode[T int] struct {
	left, right *BinaryNode[T]
	value int
}

type Queue[T int] struct {
	start, end *QueueNode[T]
	length int
}

type QueueNode[T int] struct {
	next *QueueNode[T]
	val int
}

func main() {
	root := &BinaryNode[int]{value: 7}
	root.left = &BinaryNode[int]{value: 23}
	root.right = &BinaryNode[int]{value: 8}
	root.left.left = &BinaryNode[int]{value: 5}
	root.left.right = &BinaryNode[int]{value: 4}
	root.right.left = &BinaryNode[int]{value: 21}
	root.right.right = &BinaryNode[int]{value: 15}

	root.printTree("", true)


	fmt.Println("queue display:")
	queue := Queue[int]{length: 0}
	queue.enqueue(1)
	queue.enqueue(4)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.deque()
	fmt.Println(queue.length)

	q := queue.getAll()
	fmt.Println(q)
}

func (q *Queue[T])getAll() []int{
	var elems []int
	for e := q.start; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (q *Queue[T])enqueue(el int) {
	if q.length == 0 {
		q.start = &QueueNode[T]{val: el}
		q.end = q.start 
	} else {
		q.end.next = &QueueNode[T]{val: el}
		q.end = q.end.next 
	}
	q.length++
}

func (q *Queue[T])deque() int  {
	if q == nil {
		return -1
	}

	q.length--	
	tmp := q.start
	q.start = q.start.next

	tmp.next = nil

	return tmp.val
}

/*func (tree *BinaryNode[T])bfs(needle int) bool {
	q := []int{tree}

}*/

func (tree *BinaryNode[T])printTree(prefix string, isLeft bool) {
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

