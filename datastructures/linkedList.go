// linkedlist.go
// Generische, einfach verkettete Liste mit grundlegenden Methoden.
//
// Author: Till Burdorf, Lukas Gröning, Daniel Brecht
// Date: 10.03.2025

package datastructures

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	data T
}

// LinkedList ist eine generische einfach verkettete Liste.
type LinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

// getNode gibt den Node an der angegebenen Position zurück.
// Hilfsfunktion zum durchlaufen der Liste
func (list *LinkedList[T]) getNode(pos int) *Node[T] {
	current := list.head
	currentPosition := 0
	for currentPosition < pos {
		current = current.next
		currentPosition++
	}
	return current
}

// Get gibt das Element an der angegebenen Position zurück.
func (list *LinkedList[T]) Get(pos int) (T, error) {
	if pos < 0 || pos >= list.length {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return list.getNode(pos).data, nil
}

// Add fügt ein oder mehrere Elemente ans Ende der Liste an.
func (list *LinkedList[T]) Add(datas ...T) {
	for _, data := range datas {
		list.Append(data)
	}
}

// Insert fügt ein Element an der angegebenen Position ein.
func (list *LinkedList[T]) Insert(pos int, data T) error {
	if pos < 0 || pos >= list.length {
		return errors.New("index out of bounds")
	}
	if pos == 0 {
		list.Prepend(data)
		return nil
	}
	list.getNode(pos - 1).next = &Node[T]{data: data, next: list.getNode(pos + 1)}
	list.length++
	return nil
}

// Remove entfernt das erste Vorkommen des angegebenen Elements.
func (list *LinkedList[T]) Remove(elem T) error {
	if list.Size() == 0 {
		return errors.New("list is empty")
	}

	if list.head.data == elem {
		list.head = list.head.next
		list.length--
		return nil
	}
	current := list.head
	for current.next != nil && current.next.data != elem {
		current = current.next
	}

	if current.next == nil {
		return errors.New("element not found")
	}

	current.next = current.next.next
	list.length--
	return nil
}

// RemoveAt entfernt das Element an der angegebenen Position.
func (list *LinkedList[T]) RemoveAt(pos int) error {
	if list.Size() == 0 {
		return errors.New("list is empty")
	} else if pos < 0 || pos > list.Size() {
		return errors.New("index out of bounds")
	}

	if pos == 0 {
		list.head = list.head.next
	} else {
		list.getNode(pos - 1).next = list.getNode(pos + 1)
	}
	list.length--
	return nil
}

// Replace ersetzt das Element an der angegebenen Position mit dem neuen Wert.
func (list *LinkedList[T]) Replace(pos int, val T) {
	if pos == 0 {
		list.head.data = val
		return
	}
	list.getNode(pos).data = val
}

// Prepend fügt ein neues Element am Anfang der Liste ein.
func (list *LinkedList[T]) Prepend(data T) {
	list.head = &Node[T]{data: data, next: list.head}
	list.length++
}

// Append fügt ein neues Element am Ende der Liste ein.
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

// Print gibt alle Elemente der Liste in der Konsole aus (zeilenweise).
func (list *LinkedList[T]) Print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Size gibt die Anzahl der Elemente in der Liste zurück.
func (list *LinkedList[T]) Size() int {
	return list.length
}

// IsEmpty gibt true zurück, wenn die Liste leer ist.
func (list *LinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

// IsFull gibt true zurück, wenn die Liste mindestens ein Element enthält.
func (list *LinkedList[T]) IsFull() bool {
	return list.head != nil
}

// ToString gibt eine String-Repräsentation der Liste im Format [a, b, c] zurück.
func (list *LinkedList[T]) ToString() string {
	stringifiedList := "["
	current := list.head
	for current != nil {
		stringifiedList += fmt.Sprintf("%v", current.data)
		if current.next != nil {
			stringifiedList += ", "
		}
		current = current.next
	}
	stringifiedList += "]"
	return stringifiedList
}

// Equals prüft, ob zwei Listen inhaltlich gleich sind.
func (list *LinkedList[T]) Equals(secondList *LinkedList[T]) bool {
	if list == nil || secondList == nil {
		return list == secondList
	}

	firstNode := list.head
	secondNode := secondList.head

	for firstNode != nil && secondNode != nil {
		if firstNode.data != secondNode.data {
			return false
		}
		firstNode = firstNode.next
		secondNode = secondNode.next
	}
	return firstNode == nil && secondNode == nil
}
