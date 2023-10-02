package main

import (
	"fmt"
)

type BinaryNode[T int] struct {
	left, right *BinaryNode[T]
	value int
}

func main() {
	root := &BinaryNode[int]{value: 5}
	root.left = &BinaryNode[int]{value: 3}
	root.right = &BinaryNode[int]{value: 8}
	root.left.left = &BinaryNode[int]{value: 1}
	root.left.right = &BinaryNode[int]{value: 4}
	root.right.left = &BinaryNode[int]{value: 7}
	root.right.right = &BinaryNode[int]{value: 9}

	fmt.Println("pre order search:")
	root.preOrderSearch()
	fmt.Println("in order search:")
	root.inOrderSearch()
	fmt.Println("post order search:")
	root.postOrderSearch()

	root.printTree("", true)
}

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

func (tree *BinaryNode[T]) appendLeft(val int) {
	tree.left.value = val
}

func (curr *BinaryNode[T]) preOrderWalk(path *[]int) []int {
	if curr == nil {
		return *path
	}

	*path = append(*path, curr.value)
	curr.left.preOrderWalk(path)
	curr.right.preOrderWalk(path)
	
	return *path
}

func (curr *BinaryNode[T]) inOrderWalk(path *[]int) []int {
	if curr == nil {
		return *path
	}

	curr.left.inOrderWalk(path)
	*path = append(*path, curr.value)
	curr.right.inOrderWalk(path)
	
	return *path
}

func (curr *BinaryNode[T]) postOrderWalk(path *[]int) []int {
	if curr == nil {
		return *path
	}

	curr.left.postOrderWalk(path)
	curr.right.postOrderWalk(path)
	*path = append(*path, curr.value)
	
	return *path
}

func (head *BinaryNode[T]) preOrderSearch() []int {
	trace := []int{}
	res := head.preOrderWalk(&trace)
	fmt.Println(trace)
	return res
}

func (head *BinaryNode[T]) inOrderSearch() []int {
	trace := []int{}
	res := head.inOrderWalk(&trace)
	fmt.Println(trace)
	return res
}

func (head *BinaryNode[T]) postOrderSearch() []int {
	trace := []int{}
	res := head.postOrderWalk(&trace)
	fmt.Println(trace)
	return res
}

