package datastructure

import (
	"testing"
)

func setupLinkedList() *LinkedList[int] {
	list := &LinkedList[int]{}
	return list
}

func TestLinkedListAdd(t *testing.T) {
	list := setupLinkedList()

	// Füge mehrere Elemente hinzu
	list.Add(1, 2, 3, 4, 5)

	// Überprüfe, ob die Länge korrekt ist
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, got %d", list.Size())
	}

	// Überprüfe, ob die Elemente in der richtigen Reihenfolge eingefügt wurden
	expectedValues := []int{1, 2, 3, 4, 5}
	for i, expected := range expectedValues {
		if list.GetValue(i) != expected {
			t.Errorf("Expected value at index %d to be %d, got %d", i, expected, list.GetValue(i))
		}
	}
}
