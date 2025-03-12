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

func TestStackPushAll(t *testing.T) {
	stack1 := Stack[int]{}
	stack1.PushAll(1, 2, 3)
	stack2 := Stack[int]{}
	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)
	if !stack1.Equals(&stack2) {
		t.Errorf("Expected stacks to be the same. Stack1: %v, Stack2: %v", stack1.ToString(), stack2.ToString())
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

	// stack is not supposed to change size when using peek
	stack.Peek()
	if stack.Size() != 1 {
		t.Errorf("Expected Peek not to change stack size, got %v", stack.Size())
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

func TestStackToString(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{1, 2, 3}, "[3, 2, 1]"},
	}

	for _, test := range tests {
		stack := Stack[int]{}
		for _, item := range test.input {
			stack.Push(item)
		}
		result := stack.ToString()
		if result != test.expected {
			t.Errorf("got %s, want %s", result, test.expected)
		}
	}
}

func TestStackEqual(t *testing.T) {
	tests := []struct {
		first []int
		second []int
		expected bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
	}

	for _, test := range tests {
		list1 := Stack[int]{}
		for _, item := range test.first {
			list1.Push(item)
		}
		list2 := Stack[int]{}
		for _, item := range test.second {
			list2.Push(item)
		}
		result := list1.Equals(&list2)
		if result != test.expected {
			t.Errorf("Equals(%v, %v) got %v, want %v", test.first, test.second, result, test.expected)
		}
	}
}