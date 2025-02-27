package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

func (list *LinkedList) addFront(data int) {
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) addBack(data int) {
	
	newNode := &Node{data: data, next: nil}

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

func (list *LinkedList) print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
