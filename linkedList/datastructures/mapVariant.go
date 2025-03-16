package datastructures

type CollectionType int

const (
	LinkedListType CollectionType = iota
	QueueType
	StackType
)

func (list *LinkedList[T]) MapVariant(operation func(T) any, collectionType CollectionType) any {
	current := list.head

	switch collectionType {
	case LinkedListType:
		newList := LinkedList[any]{}
		for current != nil {
			newList.Append(operation(current.data))
			current = current.next
		}
		return newList

	case QueueType:
		newQueue := Queue[any]{}
		for current != nil {
			newQueue.Enqueue(operation(current.data))
			current = current.next
		}
		return newQueue

	case StackType:
		newStack := Stack[any]{}
		for current != nil {
			newStack.Push(operation(current.data))
			current = current.next
		}
		return newStack

	default:
		println("Type not found")
		return nil
	}
}
