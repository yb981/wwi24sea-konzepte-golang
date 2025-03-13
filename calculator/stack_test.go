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
		testStack = append(testStack, input.input)
	}

	for index, _ := range testStack {
		input := testStack.Pop()

		if len(testStack) > 0 {
			if testStack[len(testStack)-1] != tests[len(tests)-(index+2)].expected {
				t.Errorf("Expected number in Stack %v but got %v", tests[len(tests)-index+2].expected, testStack[len(testStack)-1])
			}
		}
		if input != tests[len(tests)-(index+1)].input && input != tests[len(tests)-index].expected {
			t.Errorf("Got %v but expected %v", input, tests[len(tests)-index].expected)
		}
	}

	//testStack ist leer
	input := testStack.Pop()
	if input != 0 {
		t.Errorf("Expected for empty Stack %v but got %v", 0, input)
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
		testStack = append(testStack, input.input)
	}

	for index, _ := range testStack {
		input := testStack.Top()
		testStack = testStack[:len(testStack)-1]

		if input != tests[len(tests)-(index+1)].input && input != tests[len(tests)-index].expected {
			t.Errorf("Got %v but expected %v", input, tests[len(tests)-index].expected)
		}
	}

	input := testStack.Top()
	if input != 0 {
		t.Errorf("Got %v but expected ZERO", input)
	}
}
