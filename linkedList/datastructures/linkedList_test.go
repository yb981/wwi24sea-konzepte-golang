package datastructures

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	list := LinkedList[int]{}
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}
}

func TestAddAndAppend(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Append(4)

	if list.Size() != 4 {
		t.Errorf("Expected list size 4, got %d", list.Size())
	}
	val, _ := list.Get(0)
	if val != 1 {
		t.Errorf("Expected first element 1, got %d", val)
	}
	val, _ = list.Get(3)
	if val != 4 {
		t.Errorf("Expected last element 4, got %d", val)
	}

	_, err := list.Get(5)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestInsert(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Insert(2, 99)
	val, _ := list.Get(2)
	if val != 99 {
		t.Errorf("Expected inserted element 99 at position 2, got %d", val)
	}
}

func TestRemove(t *testing.T) {
	list := LinkedList[int]{}

	err := list.Remove(5)
	if err == nil {
		t.Errorf("Expected error")
	}

	list.Add(1, 2, 3)
	list.Remove(2)
	if list.Size() != 2 {
		t.Errorf("Expected list size 2 after removal, got %d", list.Size())
	}
}

func TestRemoveAt(t *testing.T) {
	list := LinkedList[int]{}

	// Test empty list
	err := list.RemoveAt(0)
	if err == nil {
		t.Errorf("Expected error")
	}

	list.Add(1, 2, 3)
	list.RemoveAt(1)
	if list.Size() != 2 {
		t.Errorf("Expected list size 2 after RemoveAt, got %d", list.Size())
	}

	err = list.RemoveAt(5)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestReplace(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Replace(1, 42)
	val , _ := list.Get(1)
	if val != 42 {
		t.Errorf("Expected replaced value 42 at position 1, got %d", val)
	}
}

func TestAddFront(t *testing.T) {
	list := LinkedList[int]{}
	list.Prepend(0)
	val , _ := list.Get(0)
	if val!= 0 {
		t.Errorf("Expected first element 0 after AddFront, got %d", val)
	}
}
