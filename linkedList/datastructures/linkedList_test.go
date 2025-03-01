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
	if list.Get(0) != 1 {
		t.Errorf("Expected first element 1, got %d", list.Get(0))
	}
	if list.Get(3) != 4 {
		t.Errorf("Expected last element 4, got %d", list.Get(3))
	}
}

func TestInsert(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Insert(2, 99)
	if list.Get(2) != 99 {
		t.Errorf("Expected inserted element 99 at position 2, got %d", list.Get(2))
	}
}

func TestRemove(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Remove(2)
	if list.Size() != 2 {
		t.Errorf("Expected list size 2 after removal, got %d", list.Size())
	}
}

func TestRemoveAt(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.RemoveAt(1)
	if list.Size() != 2 {
		t.Errorf("Expected list size 2 after RemoveAt, got %d", list.Size())
	}
}

func TestReplace(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Replace(1, 42)
	if list.Get(1) != 42 {
		t.Errorf("Expected replaced value 42 at position 1, got %d", list.Get(1))
	}
}

func TestAddFront(t *testing.T) {
	list := LinkedList[int]{}
	list.Prepend(0)
	if list.Get(0) != 0 {
		t.Errorf("Expected first element 0 after AddFront, got %d", list.Get(0))
	}
}
