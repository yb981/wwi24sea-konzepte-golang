package datastructures

import (
	"errors"
)

// -----------------------------------------------------------------------------------------
// For Each Method
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

//------------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------------
// Filter Method
// -----------------------------------------------------------------------------------------
func (list *LinkedList[T]) Filter(operation func(T) bool) LinkedList[T] {
	current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
		if operation(current.data) {
			newList.Append(current.data)
		}
		current = current.next
	}
	return *newList
}

func (queue *Queue[T]) Filter(operation func(T) bool) *Queue[T] {
	return &Queue[T]{list: queue.list.Filter(operation)}
}

func (stack *Stack[T]) Filter(operation func(T) bool) *Stack[T] {
	return &Stack[T]{list: stack.list.Filter(operation)}
}

//------------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------------
// Map Method
// -----------------------------------------------------------------------------------------
// Define the Map method on the generic LinkedList type
func (list *LinkedList[T]) Map(operation func(T) T) LinkedList[T] {
	current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
		newList.Append(operation(current.data))
		current = current.next
	}
	return *newList
}

func (queue *Queue[T]) Map(operation func(T) T) *Queue[T] {
	return &Queue[T]{list: queue.list.Map(operation)}
}

func (stack *Stack[T]) Map(operation func(T) T) *Stack[T] {
	return &Stack[T]{list: stack.list.Map(operation)}
}

//------------------------------------------------------------------------------------------

func Map[T comparable, U comparable](list LinkedList[T], operation func(T) U) LinkedList[U] {
	current := list.head
	newList := &LinkedList[U]{}
	for current != nil {
		newList.Append(operation(current.data))
		current = current.next
	}
	return *newList
}

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

func Reduce[T comparable, U comparable](list LinkedList[T], operation func(T, T) U) (U, error) {
	if list.head == nil {
		var zero U
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
