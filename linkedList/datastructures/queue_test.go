package datastructures

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	if queue.IsEmpty() {
		t.Errorf("Expected queue not to be empty after enqueue operations")
	}

	value, err := queue.Peek()
	if err != nil || value != 10 {
		t.Errorf("Expected Peek value to be 10, got %v", value)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(10)
	queue.Enqueue(20)

	value, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Expected no error when dequeueing from non-empty queue")
	}
	if value != 10 {
		t.Errorf("Expected Dequeue value to be 10, got %v", value)
	}

	value, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Expected no error when dequeueing from non-empty queue")
	}
	if value != 20 {
		t.Errorf("Expected Dequeue value to be 20, got %v", value)
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeueing from empty queue")
	}
}

func TestQueue_Peek(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(10)

	value, err := queue.Peek()
	if err != nil || value != 10 {
		t.Errorf("Expected Peek value to be 10, got %v", value)
	}

	queue.Dequeue()  // Remove the only element

	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty queue")
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	queue := Queue[int]{}

	if !queue.IsEmpty() {
		t.Errorf("Expected new queue to be empty")
	}

	queue.Enqueue(10)
	if queue.IsEmpty() {
		t.Errorf("Expected queue not to be empty after enqueue")
	}

	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeueing all elements")
	}
}

func TestQueueIsFull(t *testing.T) {
	queue := Queue[int]{}
	// Assuming the list implementation does not constrain size, hence always flexible
	queue.Enqueue(10)
	if !queue.IsFull() {
		t.Errorf("Expected IsFull to return true when queue has elements")
	}

	queue.Dequeue()
	if queue.IsFull() {
		t.Errorf("Expected IsFull to return false when queue is empty")
	}
}

func TestQueueSize(t *testing.T) {
	tests := []struct {
		input []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1,2}, 2},
	}

	for _, test := range tests {
		queue := Queue[int]{}
		for _, item := range test.input {
			queue.Enqueue(item)
		}
		result := queue.Size()
		if result != test.expected {
			t.Errorf("Expected size: %v, but got %v", test.expected, result)
		}
	}
}

func TestQueueToString(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{1, 2, 3}, "[1, 2, 3]"},
	}

	for _, test := range tests {
		queue := Queue[int]{}
		for _, item := range test.input {
			queue.Enqueue(item)
		}
		result := queue.ToString()
		if result != test.expected {
			t.Errorf("got %s, want %s", result, test.expected)
		}
	}
}

func TestEqual(t *testing.T) {
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
		list1 := Queue[int]{}
		for _, item := range test.first {
			list1.Enqueue(item)
		}
		list2 := Queue[int]{}
		for _, item := range test.second {
			list2.Enqueue(item)
		}
		result := list1.Equals(&list2)
		if result != test.expected {
			t.Errorf("Equals(%v, %v) got %v, want %v", test.first, test.second, result, test.expected)
		}
	}
}