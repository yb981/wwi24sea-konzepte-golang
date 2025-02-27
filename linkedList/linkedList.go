package main

import (
	"fmt"
)

type LinkedList[T any] struct {
	head *Node[T]
}

// adds a new element at the front of the list 
func (list *LinkedList[T]) addFront(data T) {
	newNode := &Node[T]{data: data, next: list.head}
	list.head = newNode
}

// adds a new element at the end of the list
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

// prints all elements on the console
func (list *LinkedList[T]) print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// is true if the list is empty
func (list *LinkedList[T]) isEmpty() bool {
	return list.head == nil
}

// is true if the list is not empty 
func (list *LinkedList[T]) isFull() bool {
	return list.head != nil
}

// get the current size of the list
func (list *LinkedList[T]) size() int {
	current := list.head
	size := 0
	for current != nil {
		size++
		current = current.next
	}
	return size
}

// insert an element at a position
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

// get the element at position pos
func (list *LinkedList[T]) get(pos int) T {
	current := list.head
	currentPosition := 0
	for currentPosition < pos {
		current = current.next
		currentPosition++
	}
	return current.data
}

// add one or multiple elements to the list
func (list* LinkedList[T]) add(datas ... T){
	for _, data := range datas {
        list.addBack(data)
    }
}
