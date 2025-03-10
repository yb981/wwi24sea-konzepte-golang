package datastructures

import "errors"

type Queue[T comparable] struct {
	list LinkedList[T]
}

// adds new Element to the Queue
func (queue *Queue[T]) Enqueue(data T) {
	queue.list.Append(data)
}

// removes last element from the Queue and returns it
func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.list.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	output, _ := queue.list.Get(0)
	queue.list.RemoveAt(0)
	return output, nil
}

func (queue *Queue[T]) Peek() (T, error) {
	if queue.list.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	output, _ := queue.list.Get(0)
	return output, nil
}

func (queue* Queue[T]) Size() int{
	return queue.list.Size()
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *Queue[T]) IsFull() bool {
	return queue.list.IsFull()
}

func (queue* Queue[T]) ToString() string{
	return queue.list.ToString()
}

func (queue* Queue[T]) Equals(compare* Queue[T]) bool{
	return queue.list.Equals(&compare.list)
}