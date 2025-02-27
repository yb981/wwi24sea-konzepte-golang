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

	if list.head == nil {
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

func (list *LinkedList[T]) isEmpty() bool {
	return list.head == nil
}

func (list *LinkedList[T]) isFull() bool {
	return list.head != nil
}

func (list *LinkedList[T]) size() int {
	current := list.head
	size := 0
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (list *LinkedList[T]) insert(position int, data T) {
	current := list.head
	currentPosition := 0
	for currentPosition < position-1 {
		current = current.next
		currentPosition++
	}
	newNode := &Node[T]{data: data, next: current.next}
	current.next = newNode
}

func (list *LinkedList[T]) get(pos int) T {
	current := list.head
	currentPosition := 0
	for currentPosition < pos {
		current = current.next
		currentPosition++
	}
	return current.data
}

func (list* LinkedList[T]) add(datas ... T){
	for _, data := range datas {
        list.addBack(data)
    }
}