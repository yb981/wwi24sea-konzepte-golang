// stack.go
// Generischer Stack (LIFO-Datenstruktur) basierend auf einer verketteten Liste.
//
// Author: Till Burdorf, Lukas Gröning, Daniel Brecht
// Date: 10.03.2025

package datastructures

import (
	"errors"
)

// Stack ist eine generische Stack-Implementierung nach dem LIFO-Prinzip.
type Stack[T comparable] struct {
	list LinkedList[T]
}

// Push fügt ein Element oben auf den Stack.
func (stack *Stack[T]) Push(data T) {
	stack.list.Prepend(data)
}

// PushAll fügt mehrere Elemente in umgekehrter Reihenfolge auf den Stack.
func (stack *Stack[T]) PushAll(datas ...T) {
	for _, data := range datas {
		stack.list.Prepend(data)
	}
}

// Pop entfernt und gibt das oberste Element des Stacks zurück.
// Gibt einen Fehler zurück, wenn der Stack leer ist.
func (stack *Stack[T]) Pop() (T, error) {
	if stack.list.Size() == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	output, _ := stack.list.Get(0)
	stack.list.RemoveAt(0)
	return output, nil
}

// Peek gibt das oberste Element zurück, ohne es zu entfernen.
func (stack *Stack[T]) Peek() (T, error) {
	if stack.list.Size() == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	output, _ := stack.list.Get(0)
	return output, nil
}

// IsEmpty gibt true zurück, wenn der Stack leer ist.
func (stack *Stack[T]) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// IsFull gibt true zurück, wenn der Stack mindestens ein Element enthält.
func (stack *Stack[T]) IsFull() bool {
	return stack.list.IsFull()
}

// Size gibt die Anzahl der Elemente im Stack zurück.
func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

// ToString gibt eine String-Repräsentation des Stacks zurück.
func (stack *Stack[T]) ToString() string {
	return stack.list.ToString()
}

// Equals prüft, ob zwei Stacks denselben Inhalt und dieselbe Reihenfolge haben.
func (stack *Stack[T]) Equals(compare *Stack[T]) bool {
	return stack.list.Equals(&compare.list)
}
