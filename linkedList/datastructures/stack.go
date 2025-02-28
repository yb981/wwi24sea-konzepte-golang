package datastructures

type Stack[T comparable] struct {
	list LinkedList[T]
}

func (stack *Stack[T]) Push(data T) {
	stack.list.AddFront(data)
}

func (stack *Stack[T]) Pop() T {
	output := stack.list.GetValue(0)
	stack.list.RemoveAt(0)
	return output
}

func (stack *Stack[T]) Peek() T {
	return stack.list.GetValue(0)
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
