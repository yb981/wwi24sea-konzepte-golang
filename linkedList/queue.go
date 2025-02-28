package main

type Queue[T comparable] struct {
	list LinkedList[T]
}

func (queue *Queue[T]) enqueue(data T) {
	queue.list.addFront(data)
}
