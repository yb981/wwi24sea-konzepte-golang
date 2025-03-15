package datastructures

import (
	"testing"
)

func TestForEach(t *testing.T) {

	expected := []int{1, 2, 3} //slice, den wir erwarten (slice= array mit flexibler größe)

	// Test für LinkedList
	ll := LinkedList[int]{} //neue linkedlist(ll) wird erstellt
	ll.Add(1, 2, 3)

	result := []int{}            //leeres slice
	ll.ForEach(func(value int) { //mit foreach packen wir jedes element aus ll in slice
		result = append(result, value) //
	})

	if len(result) != len(expected) { //Go kann nicht result (als slice) == expected vergleichen, daher vergleichen wir erst die Länge
		t.Errorf("LinkedList ForEach failed. Length mismatch. Expected %v, got %v", len(expected), len(result))
	}

	for i := range expected { //für i wird jedes element verglichen
		if result[i] != expected[i] {
			t.Errorf("LinkedList ForEach failed. Expected %v, got %v", expected[i], result[i])
		}
	}

	// Test für Queue
	queue := &Queue[int]{list: ll}

	result = []int{}
	queue.ForEach(func(value int) {
		result = append(result, value)
	})

	if len(result) != len(expected) {
		t.Errorf("Queue ForEach failed. Length mismatch. Expected %v, got %v", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Queue ForEach failed. Expected %v, got %v", expected[i], result[i])
		}
	}

	// Test für Stack
	stack := &Stack[int]{list: ll}

	result = []int{}
	stack.ForEach(func(value int) {
		result = append(result, value)
	})

	if len(result) != len(expected) {
		t.Errorf("Stack ForEach failed. Length mismatch. Expected %v, got %v", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Stack ForEach failed at index %d. Expected %v, got %v", i, expected[i], result[i])
		}
	}
}

func TestFilter(t *testing.T) {
	isEvenFunc := func(a int) bool {
		return a%2 == 0
	}

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2}},
		{[]int{-3, -2, -1}, []int{-2}},
		{[]int{50, -2, 20}, []int{50, -2, 20}},
	}

	// test for linked list
	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}
		testList = testList.Filter(isEvenFunc)

		var resultList LinkedList[int]
		for _, v := range test.expected {
			resultList.Append(v)
		}
		if !testList.Equals(&resultList) {
			t.Errorf("Expected %v, but got %v", resultList.ToString(), testList.ToString())
		}
	}

	// test for queue
	for _, test := range tests {
		var testQueue Queue[int]
		for _, v := range test.input {
			testQueue.Enqueue(v)
		}
		testQueue = *testQueue.Filter(isEvenFunc)

		var resultList Queue[int]
		for _, v := range test.expected {
			resultList.Enqueue(v)
		}
		if !testQueue.Equals(&resultList) {
			t.Errorf("Expected %v, but got %v", resultList.ToString(), testQueue.ToString())
		}
	}

	// test for stack
	for _, test := range tests {
		var testStack Stack[int]
		for _, v := range test.input {
			testStack.Push(v)
		}
		testStack = *testStack.Filter(isEvenFunc)

		var resultStack Stack[int]
		for _, v := range test.expected {
			resultStack.Push(v)
		}
		if !testStack.Equals(&resultStack) {
			t.Errorf("Expected %v, but got %v", resultStack.ToString(), testStack.ToString())
		}
	}
}
