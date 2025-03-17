package datastructures

import (
	"errors"
)

type FilterFunc[T any] func() bool

type LazyFilterList[T comparable] struct {
	data       []T
	Operations []FilterFunc[T]
}

type MapFunc[T any] func() T

type LazyMapList[T comparable] struct {
	Operations []MapFunc[T]
}

type Collection[T any] interface {
	ToString() string
}

type CollectionType int

const (
	LinkedListType CollectionType = iota
	QueueType
	StackType
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

func (queue *Queue[T]) Filter(operation func(T) bool) *Queue[T] {
	return &Queue[T]{list: *queue.list.Filter(operation)}
}

func (stack *Stack[T]) Filter(operation func(T) bool) *Stack[T] {
	return &Stack[T]{list: *stack.list.Filter(operation)}
}

//------------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------------
// Map Method
// -----------------------------------------------------------------------------------------
// Define the Map method on the generic LinkedList type
func (list *LinkedList[T]) Map(operation func(T) T) *LinkedList[T] {
	current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
		newList.Append(operation(current.data))
		current = current.next
	}
	return newList
}

func (queue *Queue[T]) Map(operation func(T) T) *Queue[T] {
	return &Queue[T]{list: *queue.list.Map(operation)}
}

func (stack *Stack[T]) Map(operation func(T) T) *Stack[T] {
	return &Stack[T]{list: *stack.list.Map(operation)}
}

// Lazy Map missing
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

func (l LazyMapList[T]) ExecuteMap() *LinkedList[T] {
	output := &LinkedList[T]{}
	for _, op := range l.Operations {
		value := op()
		output.Append(value)
	}
	return output
}

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

