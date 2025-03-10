package datastructures

import "errors"

type Stack[T comparable] struct {
	list LinkedList[T]
}

func (stack *Stack[T]) Push(data T) {
	stack.list.Prepend(data)
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.list.Size() == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	output, _ := stack.list.Get(0)
	stack.list.RemoveAt(0)
	return output, nil
}

func (stack *Stack[T]) Peek() (T, error) {
	if stack.list.Size() == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	output, _ := stack.list.Get(0)
	return output, nil
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.list.IsEmpty()
}

func (stack *Stack[T]) IsFull() bool {
	return stack.list.IsFull()
}

func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}
