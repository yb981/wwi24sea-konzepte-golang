package main

import (
	"fmt"
)

type LinkedList[T any] struct {
	head *Node[T]
}

func (list *LinkedList[T]) addFront(data T) {
	newNode := &Node[T]{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList[T]) addBack(data T) {
	
	newNode := &Node[T]{data: data, next: nil}

	if list.head == nil{
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList[T]) print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
