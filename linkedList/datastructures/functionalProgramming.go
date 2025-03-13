package datastructures

// For Each Method
func (list *LinkedList[T]) ForEach(operation func(T)){
	current := list.head
	for current != nil {
		operation(current.data)
		current = current.next
    }
}

// Filter Method 
func (list *LinkedList[T]) Filter(operation func(T) bool) LinkedList[T]{
    current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
        if operation(current.data){
            newList.Append(current.data)
        }
		current = current.next
	}
	return *newList
}

// Map Method
func (list *LinkedList[T]) Map(operation func(T) T) LinkedList[T] {
	current := list.head
	newList := &LinkedList[T]{}
	for current != nil {
		element := operation(current.data)
		newList.Append(element)
		current = current.next
	}
	return *newList
}

func (queue *Queue[T]) Map(operation func(T) T) Queue[T] {
	newQueue := &Queue[T]{}
	newQueue.list = queue.list.Map(operation)
	return *newQueue
}

func (stack *Stack[T]) Map(operation func(T) T) Stack[T] {
	newStack := &Stack[T]{}
	newStack.list = stack.list.Map(operation)
	return *newStack
}