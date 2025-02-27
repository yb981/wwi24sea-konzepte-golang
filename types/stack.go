// stack.go
// Package stack implements a basic stack data structure.
//
// Author: Lukas Gröning
// Date: 26.02.2025
//
// This file contains methods for stack operations.

package types

import (
	"errors"
)


type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(input T) {
	s.items = append(s.items, input)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("can't pop from empty stack")
	}

	lastValue := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lastValue, nil
}

func (s Stack[T]) Peek() (T, error) {
	if len(s.items) != 0 {
		return s.items[len(s.items)-1], nil
	}

	var zero T
	return zero, errors.New("Stack is empty")
}

func (s Stack[T]) Size() int {
	return len(s.items)
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Stack[T]) IsFull() bool {
	return len(s.items) != 0
}