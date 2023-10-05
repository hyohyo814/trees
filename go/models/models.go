package models

type BinaryNode[T any] struct {
	left, right *BinaryNode[T]
	val         T
}

type Queue[T any] struct {
	head, tail *BinaryNode[T]
}

type QueueNode[T any] struct {
	next *QueueNode[T]
	val  T
}
