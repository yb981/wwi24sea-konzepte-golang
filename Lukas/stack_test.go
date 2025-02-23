package main

import (
	"testing"
)

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
		if testStack[len(testStack)-1] != number.expected {
			t.Errorf("Expected %v on last position in Stack but got %v", number.expected, number.input)
		}
	}
}
