// stack.go
// Package stack implements a basic stack data structure.
//
// Author: Lukas Gröning
// Date: 22.02.2025
//
// This file contains methods for stack operations.

package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(input T) {
	*s = append(*s, input)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}

	lastValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return lastValue
}

func (s Stack[T]) Top() T {
	if len(s) != 0 {
		return s[len(s)-1]
	}
	var zero T
	return zero
}

func (s Stack[T]) Print() {
	for _, value := range s {
		fmt.Print(value, " ")
	}
}