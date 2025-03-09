package datastructures

import (
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	data T
}

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
func (list *LinkedList[T]) Get(pos int) T {
	return list.getNode(pos).data
}

// add one or multiple elements to the list
func (list *LinkedList[T]) Add(datas ...T) {
	for _, data := range datas {
		list.Append(data)
	}
}

// insert an element at a position
func (list *LinkedList[T]) Insert(pos int, data T) {
	if pos == 0 {
		list.Prepend(data)
		return
	}
	// the node with position - 1 next pointer is set to a new node which contains data and the node after the chosen  position as next
	list.getNode(pos - 1).next = &Node[T]{data: data, next: list.getNode(pos + 1)}
	list.length++
}

// removes the element elem from the list
func (list *LinkedList[T]) Remove(elem T) {
	current := list.head
	for current.next.data != elem {
		current = current.next
	}
	current.next = current.next.next
	list.length--
}

// removes the element at position pos from the list
func (list *LinkedList[T]) RemoveAt(pos int) {
	if pos == 0 {
		list.head = list.head.next
	} else {
		list.getNode(pos - 1).next = list.getNode(pos + 1)
	}
	list.length--
}

// repleaces the element at position pos with new element with value val
func (list *LinkedList[T]) Replace(pos int, val T) {

	if pos == 0 {
		list.head.data = val
		return
	}
	list.getNode(pos).data = val
}

// adds a new element at the front of the list
func (list *LinkedList[T]) Prepend(data T) {
	list.head = &Node[T]{data: data, next: list.head}
	list.length++
}

// adds a new element at the end of the list
func (list *LinkedList[T]) Append(data T) {

	if list.head == nil {
		list.head = &Node[T]{data: data, next: nil}
		list.length++
		return
	}

	// füge eine Node hinzu setze
	list.getNode(list.length - 1).next = &Node[T]{data: data, next: nil}
	list.length++
}

// prints all elements on the console
func (list *LinkedList[T]) Print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// get the current size of the list
func (list *LinkedList[T]) Size() int {
	return list.length
}

// is true if the list is empty
func (list *LinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

// is true if the list is not empty
func (list *LinkedList[T]) IsFull() bool {
	return list.head != nil
}