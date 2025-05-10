// queue.go
//
// Generische Queue-Implementierung mit FIFO-Prinzip.
//
// Author: Till Burdorf, Lukas Gröning, Daniel Brecht
// Date: 10.03.2025

package datastructures

import "errors"

// Queue ist eine generische Warteschlange mit FIFO-Verhalten.
type Queue[T comparable] struct {
	list LinkedList[T]
}

// Enqueue fügt ein Element am Ende der Queue hinzu.
func (queue *Queue[T]) Enqueue(data T) {
	queue.list.Append(data)
}

// Dequeue entfernt und gibt das erste Element (vorderstes) zurück.
// Gibt einen Fehler zurück, wenn die Queue leer ist.
func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.list.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	output, _ := queue.list.Get(0)
	queue.list.RemoveAt(0)
	return output, nil
}

// Peek gibt das vorderste Element zurück, ohne es zu entfernen.
func (queue *Queue[T]) Peek() (T, error) {
	if queue.list.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	output, _ := queue.list.Get(0)
	return output, nil
}

// Size gibt die Anzahl der Elemente in der Queue zurück.
func (queue *Queue[T]) Size() int {
	return queue.list.Size()
}

// IsEmpty gibt true zurück, wenn die Queue leer ist.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

// IsFull gibt true zurück, wenn die Queue Elemente enthält.
func (queue *Queue[T]) IsFull() bool {
	return queue.list.IsFull()
}

// ToString gibt eine String-Darstellung der Queue zurück.
func (queue *Queue[T]) ToString() string {
	return queue.list.ToString()
}

// Equals prüft, ob zwei Queues denselben Inhalt haben.
func (queue *Queue[T]) Equals(compare *Queue[T]) bool {
	return queue.list.Equals(&compare.list)
}
