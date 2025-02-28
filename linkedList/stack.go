package main

type Stack[T comparable] struct {
	list LinkedList[T]
}

func (stack *Stack[T]) push(data T) {
	stack.list.addFront(data)
}

func (stack *Stack[T]) pop() T {
	output := stack.list.getValue(0)
	stack.list.removeAt(0)
	return output
}

func (stack *Stack[T]) peek() T {
	return stack.list.getValue(0)
}
