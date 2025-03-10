package datastructures

type Stack[T comparable] struct {
	list LinkedList[T]
}

func (stack *Stack[T]) Push(data T) {
	stack.list.Prepend(data)
}

func (stack *Stack[T]) PushAll(datas ...T) {
	for _, data := range datas {
		stack.list.Prepend(data)
	}
}

func (stack *Stack[T]) Pop() T {
	output := stack.list.Get(0)
	stack.list.RemoveAt(0)
	return output
}

func (stack *Stack[T]) Peek() T {
	return stack.list.Get(0)
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

func (stack *Stack[T]) ToString() string {
	return stack.list.ToString()
}

func (stack *Stack[T]) Equals(compare *Stack[T]) bool {
	return stack.list.Equals(&compare.list)
}
