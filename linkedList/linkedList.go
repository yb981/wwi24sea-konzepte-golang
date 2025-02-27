package main

import (
	"fmt"
)

type LinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

// get the node at position pos | helping function to iterate over the list
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
	return list.getNode(pos).data
}

// add one or multiple elements to the list
func (list *LinkedList[T]) add(datas ...T) {
	for _, data := range datas {
		list.append(data)
	}
}

// insert an element at a position
func (list *LinkedList[T]) insert(pos int, data T) {
	if pos == 0 {
		newNode := &Node[T]{data: data, next: list.head}
		list.head = newNode
		list.length++
		return
	}
	// the node with position - 1 next pointer is set to a new node which contains data and the node after the chosen  position as next
	list.getNode(pos -1).next = &Node[T]{data: data, next: list.getNode(pos + 1)}
}

// removes the element elem from the list
func (list *LinkedList[T]) remove(elem T) {
	current := list.head
	for current.next.data != elem {
		current = current.next
	}
	current.next = current.next.next
	list.length--
}

// removes the element at position pos from the list
func (list *LinkedList[T]) removeAt(pos int) {
	current := list.getNode(pos - 1)
	newNext := list.getNode(pos + 1)
	current.next = newNext
	list.length--
}

// repleaces the element at position pos with new element with value val
func (list *LinkedList[T]) replace(pos int, val T) {

	if pos == 0 {
		list.head.data = val
		return
	}
	list.getNode(pos).data = val
}

// adds a new element at the front of the list
func (list *LinkedList[T]) addFront(data T) {
	list.head = &Node[T]{data: data, next: list.head}
	list.length++
}

// adds a new element at the end of the list
func (list *LinkedList[T]) append(data T) {

	if list.head == nil {
		list.head = &Node[T]{data : data, next : nil}
		list.length++
		return
	}

	// füge eine Node hinzu setze 
	list.getNode(list.length -1).next = &Node[T]{data : data, next:nil}
	list.length++
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
func (list *LinkedList[T]) size() int {
	return list.length
}

// is true if the list is empty
func (list *LinkedList[T]) isEmpty() bool {
	return list.head == nil
}

// is true if the list is not empty
func (list *LinkedList[T]) isFull() bool {
	return list.head != nil
}
