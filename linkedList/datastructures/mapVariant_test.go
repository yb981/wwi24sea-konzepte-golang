package datastructures

import (
	"testing"
)

func TestMapVariant(t *testing.T) {
	doubleFunc := func(a int) any { return a * 2 }

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{-3, -2, -1}, []int{-6, -4, -2}},
		{[]int{0, 5, 10}, []int{0, 10, 20}},
	}

	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}

		// Test für LinkedList
		mappedList := testList.MapVariant(doubleFunc, LinkedListType).(*LinkedList[any])
		var expectedList LinkedList[any]
		for _, v := range test.expected {
			expectedList.Append(v)
		}
		if mappedList.ToString() != expectedList.ToString() {
			t.Errorf("LinkedList MapVariant failed. Expected %v, got %v", expectedList.ToString(), mappedList.ToString())
		}

		// Test für Queue
		mappedQueue := testList.MapVariant(doubleFunc, QueueType).(*Queue[any])
		var expectedQueue Queue[any]
		for _, v := range test.expected {
			expectedQueue.Enqueue(v)
		}
		if mappedQueue.ToString() != expectedQueue.ToString() {
			t.Errorf("Queue MapVariant failed. Expected %v, got %v", expectedQueue.ToString(), mappedQueue.ToString())
		}

		// Test für Stack
		mappedStack := testList.MapVariant(doubleFunc, StackType).(*Stack[any])
		var expectedStack Stack[any]
		for _, v := range test.expected {
			expectedStack.Push(v)
		}
		if mappedStack.ToString() != expectedStack.ToString() {
			t.Errorf("Stack MapVariant failed. Expected %v, got %v", expectedStack.ToString(), mappedStack.ToString())
		}
	}
}

func TestMapVariant_DefaultCase(t *testing.T) {
	list := &LinkedList[int]{}
	list.Add(1, 2, 3)

	invalidType := CollectionType(999) // Ungültiger Typ
	result := list.MapVariant(func(x int) any { return x * 2 }, invalidType)

	if result != nil {
		t.Errorf("Expected nil for invalid CollectionType, but got %v", result)
	}
}