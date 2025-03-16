package datastructures

import (
	"errors"
)


type FilterFunc[T any] func() bool

type LazyFilterList[T comparable] struct {
    data []T
	operations []FilterFunc[T]
}


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
		operations: lazyOps,
        data: datalist,
	}
}

func (l LazyFilterList[T]) Execute() LinkedList[T] {
	result := make([]bool, len(l.operations))
	for i, op := range l.operations {
		result[i] = op()
	}

    outputList := &LinkedList[T]{}

    for i := 0; i<  len(result); i++ {
        if result[i]{
            outputList.Append(l.data[i])
        }
    }
	return *outputList
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

func (list *LinkedList[T]) Reduce(operation func(T, T) T) (any, error) {
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
