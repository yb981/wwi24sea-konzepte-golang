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

func TestQueue_IsFull(t *testing.T) {
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