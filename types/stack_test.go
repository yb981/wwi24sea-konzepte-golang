// stack_test.go
// Testing for stack.go file
//
// Author: Lukas Gröning
// Date: 26.02.2025
//
// This file contains tests.

package types

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	q := NewStack[float64]()
	if q == nil {
		t.Error("Expected a new Stack, instead recieved nil")
	}
}

func TestPush(t *testing.T) {
	var tests = []struct {
		input    float64
		expected float64
	}{
		{2, 2},
		{-1, -1},
		{20.9, 20.9},
		{-13.01, -13.01},
	}

	testStack := Stack[float64]{}

	for _, number := range tests {
		testStack.Push(number.input)
		if actual, err := testStack.Peek(); err != nil {
			t.Errorf("Expected %v, but got Error %v instead", number.expected, err)
		} else if actual != number.expected {
			t.Errorf("Expected %v on last position in Stack but got %v", number.expected, number.input)
		}
	}
}

func TestPop(t *testing.T) {
	var tests = []struct {
		input    float64
		expected float64
	}{
		{-15, -15},
		{2, 2},
		{14.02, 14.02},
		{0, 0},
		{99999.999, 99999.999},
	}
	testStack := Stack[float64]{}
	for _, input := range tests {
		testStack.Push(input.input)
	}

	for index, _ := range testStack.items {
		input, _ := testStack.Pop()
		if input != tests[len(tests)-(index+1)].input && input != tests[len(tests)-index].expected {
			t.Errorf("Got %v but expected %v", input, tests[len(tests)-index].expected)
		}
	}

	_, err := testStack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping empty stack")
	}
}

func TestTop(t *testing.T) {
	var tests = []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{1.0123456789, 1.0123456789},
	}

	testStack := Stack[float64]{}
	for _, input := range tests {
		testStack.Push(input.input)
	}

	for index, _ := range testStack.items {
		input, _ := testStack.Peek()
		testStack.Pop()

		if input != tests[len(tests)-(index+1)].input && input != tests[len(tests)-index].expected {
			t.Errorf("Got %v but expected %v", input, tests[len(tests)-index].expected)
		}
	}

	input, _ := testStack.Peek()
	if input != 0 {
		t.Errorf("Got %v but expected ZERO", input)
	}
}