package datastructures

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.Size() != 3 {
		t.Errorf("Expected stack size to be 3, got %v", stack.Size())
	}

	value, err := stack.Peek()
	if err != nil || value != 30 {
		t.Errorf("Expected Peek value to be 30, got %v", value)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)

	value, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error when popping from non-empty stack")
	}
	if value != 20 {
		t.Errorf("Expected Pop value to be 20, got %v", value)
	}

	value, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error when popping from non-empty stack")
	}
	if value != 10 {
		t.Errorf("Expected Pop value to be 10, got %v", value)
	}

	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from empty stack")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(10)

	value, err := stack.Peek()
	if err != nil || value != 10 {
		t.Errorf("Expected Peek value to be 10, got %v", value)
	}

	stack.Pop() // Remove the only element

	_, err = stack.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty stack")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := Stack[int]{}

	if !stack.IsEmpty() {
		t.Errorf("Expected new stack to be empty")
	}

	stack.Push(10)
	if stack.IsEmpty() {
		t.Errorf("Expected stack not to be empty after push")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements")
	}
}

func TestStack_IsFull(t *testing.T) {
	stack := Stack[int]{}
	// Assuming the list implementation does not constrain size, hence always reports true for IsFull
	stack.Push(10)
	if !stack.IsFull() {
		t.Errorf("Expected IsFull to return true when stack has elements")
	}

	stack.Pop()
	if stack.IsFull() {
		t.Errorf("Expected IsFull to return false when stack is empty")
	}
}

func TestStack_Size(t *testing.T) {
	stack := Stack[int]{}

	if stack.Size() != 0 {
		t.Errorf("Expected stack size to be 0, got %v", stack.Size())
	}

	stack.Push(10)
	if stack.Size() != 1 {
		t.Errorf("Expected stack size to be 1, got %v", stack.Size())
	}

	stack.Push(20)
	if stack.Size() != 2 {
		t.Errorf("Expected stack size to be 2, got %v", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Expected stack size to be 1 after pop, got %v", stack.Size())
	}
}