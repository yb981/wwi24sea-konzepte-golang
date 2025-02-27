package main

import (
	"fmt"
)

type LinkedList[T comparable] struct {
	head *Node[T]
	size int 
}

// get the node at position pos
func (list *LinkedList[T]) getNode(pos int) *Node[T] {
	current := list.head
	currentPosition := 0
	for currentPosition < pos {
		current = current.next
		currentPosition++
	}
	return current
}

// get the element at position pos
func (list *LinkedList[T]) getValue(pos int) T {
	current := list.head
	currentPosition := 0
	for currentPosition < pos {
		current = current.next
		currentPosition++
	}
	return current.data
}

// add one or multiple elements to the list
func (list *LinkedList[T]) add(datas ...T) {
	for _, data := range datas {
		list.addBack(data)
		list.size++
	}
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
	list.size ++
}

// removes the element elem from the list
func (list *LinkedList[T]) remove(elem T) {
	current := list.head
	for current.next.data != elem {
		current = current.next
	}
	current.next = current.next.next
	list.size--
}

// removes the element at position pos from the list
func (list *LinkedList[T]) removeAt(pos int) {
	current := list.getNode(pos - 1)
	newNext := list.getNode(pos + 1)
	current.next = newNext
	list.size --
}

// repleaces the element at position pos with new element with value val
func (list *LinkedList[T]) replace(pos int, val T) {

	if pos == 0{
		list.head.data = val
		return
	}
	list.getNode(pos).data = val
}

// adds a new element at the front of the list
func (list *LinkedList[T]) addFront(data T) {
	list.head = &Node[T]{data: data, next: list.head}
	list.size ++
}

// adds a new element at the end of the list
func (list *LinkedList[T]) addBack(data T) {

	newNode := &Node[T]{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		list.size ++
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	list.size ++
}

// prints all elements on the console
func (list *LinkedList[T]) print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// get the current size of the list
func (list *LinkedList[T]) getSize() int {
	return list.size
}

// is true if the list is empty
func (list *LinkedList[T]) isEmpty() bool {
	return list.head == nil
}

// is true if the list is not empty
func (list *LinkedList[T]) isFull() bool {
	return list.head != nil
}