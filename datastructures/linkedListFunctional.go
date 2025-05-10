// linkedListFunctional.go
//
// Funktionale Erweiterungen für generische LinkedLists, Queues und Stacks.
// Beinhaltet Implementierungen für Filter, Map und Reduce sowie Lazy Evaluation.
//
// Author: Till Burdorf, Lukas Gröning
// Date: 17.03.2025

package datastructures

import (
	"errors"
)

// -------------------------------------------
// Typdefinitionen für Lazy Evaluation
// -------------------------------------------

// FilterFunc definiert eine verzögerte Filterfunktion ohne Parameter
type FilterFunc[T any] func() bool

// LazyFilterList speichert Daten und eine Liste verzögerter Filterfunktionen
type LazyFilterList[T comparable] struct {
	data       []T
	Operations []FilterFunc[T]
}

// MapFunc definiert eine verzögerte Map-Funktion ohne Parameter
type MapFunc[T any] func() T

// LazyMapList speichert verzögerte Map-Funktionen
type LazyMapList[T comparable] struct {
	Operations []MapFunc[T]
}

// -------------------------------------------
// Gemeinsames Collection-Interface
// -------------------------------------------

type Collection[T any] interface {
	ToString() string
}

// Typenumwandlung für polymorphe Map/Filter-Ergebnisse
type CollectionType int

const (
	LinkedListType CollectionType = iota
	QueueType
	StackType
)

// -----------------------------------------------------------------------------------------
// For Each Methode
// -----------------------------------------------------------------------------------------
func (list *LinkedList[T]) ForEach(operation func(T)) {
	current := list.head
	for current != nil {
		operation(current.data)
		current = current.next
	}
}

func (queue *Queue[T]) ForEach(operation func(T)) {
	queue.list.ForEach(operation)
}

func (stack *Stack[T]) ForEach(operation func(T)) {
	stack.list.ForEach(operation)
}

// -------------------------------------------
// Filter-Methoden (eager und lazy)
// -------------------------------------------

// Gibt eine neue LinkedList zurück, gefiltert nach der Bedingung
func (list *LinkedList[T]) Filter(operation func(T) bool) *LinkedList[T] {
	current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
		if operation(current.data) {
			newList.Append(current.data)
		}
		current = current.next
	}
	return newList
}

// Gibt je nach CollectionType eine neue gefilterte Collection zurück
func (list *LinkedList[T]) FilterVariant(operation func(T) bool, collectionType CollectionType) Collection[any] {
	current := list.head

	switch collectionType {
	case LinkedListType:
		newList := LinkedList[any]{}
		for current != nil {
			if operation(current.data) {
				newList.Append(current.data)
			}
			current = current.next
		}
		return &newList

	case QueueType:
		newQueue := Queue[any]{}
		for current != nil {
			if operation(current.data) {
				newQueue.Enqueue(current.data)
			}
			current = current.next
		}
		return &newQueue

	case StackType:
		newStack := Stack[any]{}
		for current != nil {
			if operation(current.data) {
				newStack.Push(current.data)
			}
			current = current.next
		}
		return &newStack
	}

	return nil
}

// Erzeugt eine LazyFilterList, die Operationen speichert, aber nicht sofort ausführt
func (list *LinkedList[T]) LazyFilter(operation func(T) bool) LazyFilterList[T] {
	current := list.head
	lazyOps := []FilterFunc[T]{}
	var datalist []T

	for current != nil {
		value := current.data
		datalist = append(datalist, value)
		lazyOps = append(lazyOps, func() bool {
			return operation(value)
		})
		current = current.next
	}

	return LazyFilterList[T]{
		Operations: lazyOps,
		data:       datalist,
	}
}

// Führt die gespeicherten LazyFilter-Funktionen aus und gibt gefilterte LinkedList zurück
func (l LazyFilterList[T]) Execute() *LinkedList[T] {
	result := make([]bool, len(l.Operations))
	for i, op := range l.Operations {
		result[i] = op()
	}

	outputList := &LinkedList[T]{}

	for i := 0; i < len(result); i++ {
		if result[i] {
			outputList.Append(l.data[i])
		}
	}
	return outputList
}

// Varianten für Queue und Stack mit Filterfunktion
func (queue *Queue[T]) Filter(operation func(T) bool) *Queue[T] {
	return &Queue[T]{list: *queue.list.Filter(operation)}
}

func (stack *Stack[T]) Filter(operation func(T) bool) *Stack[T] {
	return &Stack[T]{list: *stack.list.Filter(operation)}
}

// -------------------------------------------
// Map-Methoden (eager und lazy)
// -------------------------------------------

// Führt eine Map-Operation auf jedes Element aus und gibt neue LinkedList zurück
func (list *LinkedList[T]) Map(operation func(T) T) *LinkedList[T] {
	current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
		newList.Append(operation(current.data))
		current = current.next
	}
	return newList
}

// Varianten für Queue und Stack mit Map-Funktion
func (queue *Queue[T]) Map(operation func(T) T) *Queue[T] {
	return &Queue[T]{list: *queue.list.Map(operation)}
}

func (stack *Stack[T]) Map(operation func(T) T) *Stack[T] {
	return &Stack[T]{list: *stack.list.Map(operation)}
}

// Erzeugt eine LazyMapList, bei der Operationen erst bei Execute ausgeführt werden
func (list *LinkedList[T]) LazyMap(operation func(T) T) LazyMapList[T] {
	current := list.head
	lazyOps := []MapFunc[T]{}

	for current != nil {
		value := current.data
		lazyOps = append(lazyOps, func() T {
			return operation(value)
		})
		current = current.next
	}

	return LazyMapList[T]{
		Operations: lazyOps,
	}
}

// Führt alle LazyMap-Operationen aus und gibt eine LinkedList zurück
func (l LazyMapList[T]) ExecuteMap() *LinkedList[T] {
	output := &LinkedList[T]{}
	for _, op := range l.Operations {
		value := op()
		output.Append(value)
	}
	return output
}

// Gibt je nach CollectionType eine neue gemappte Collection zurück
func (list *LinkedList[T]) MapVariant(operation func(T) any, collectionType CollectionType) Collection[any] {
	current := list.head

	switch collectionType {
	case LinkedListType:
		newList := LinkedList[any]{}
		for current != nil {
			newList.Append(operation(current.data))
			current = current.next
		}
		return &newList

	case QueueType:
		newQueue := Queue[any]{}
		for current != nil {
			newQueue.Enqueue(operation(current.data))
			current = current.next
		}
		return &newQueue

	case StackType:
		newStack := Stack[any]{}
		for current != nil {
			newStack.Push(operation(current.data))
			current = current.next
		}
		return &newStack

	default:
		println("Type not found")
		return nil
	}
}

// Generische Map-Funktion außerhalb von Methoden (für separate Nutzung)
func Map[T comparable, U comparable](list LinkedList[T], operation func(T) U) LinkedList[U] {
	current := list.head
	newList := &LinkedList[U]{}
	for current != nil {
		newList.Append(operation(current.data))
		current = current.next
	}
	return *newList
}

// -------------------------------------------
// Reduce-Methoden
// -------------------------------------------

// Reduziert die LinkedList auf einen einzelnen Wert
func (list *LinkedList[T]) Reduce(operation func(T, T) T) (T, error) {
	if list.head == nil {
		var zero T
		return zero, errors.New("Reduce Function not allowed on empty List")
	}
	current := list.head
	result := current.data

	for current.next != nil {
		current = current.next
		result = operation(result, current.data)
	}
	return result, nil
}

// Varianten für Queue und Stack mit Reduce-Funktion
func (queue *Queue[T]) Reduce(operation func(T, T) T) (T, error) {
	value, err := Reduce(queue.list, operation)
	var zero T
	if err != nil {
		return zero, err
	} else {
		return value, nil
	}
}

func (stack *Stack[T]) Reduce(operation func(T, T) T) (T, error) {
	value, err := Reduce(stack.list, operation)
	var zero T
	if err != nil {
		return zero, err
	} else {
		return value, nil
	}
}

// Alternative Reduce-Variante mit initialem neutralem Wert vom Typ U
func Reduce[U comparable, T comparable](list LinkedList[T], operation func(U, T) U) (U, error) {
	if list.head == nil {
		var zero U
		return zero, errors.New("Reduce Function not allowed on empty List")
	}
	current := list.head
	var initial U
	result := operation(initial, current.data)

	for current.next != nil {
		current = current.next
		result = operation(result, current.data)
	}
	return result, nil
}
