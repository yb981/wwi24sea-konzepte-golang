package datastructures

type Queue[T comparable] struct {
	list LinkedList[T]
}

func (queue *Queue[T]) Enqueue(data T) {
	queue.list.AddFront(data)
}
