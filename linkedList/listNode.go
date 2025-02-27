package main

type Node[T any] struct {
	next *Node[T]
	data T
}
