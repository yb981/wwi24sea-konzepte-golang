package datastructures

import (
	"bytes"
	"os"
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

	// insert infront
	list.Insert(0, 88)
	val, _ = list.Get(0)
	if val != 88 {
		t.Errorf("Expected inserted element 88 at position 0, got %d", val)
	}

	// testing out of bounds
	err := list.Insert(5, 100)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestRemove(t *testing.T) {
	list := LinkedList[int]{}

	err := list.Remove(5)
	if err == nil {
		t.Errorf("Expected error")
	}

	list.Add(1, 2, 3)
	list.Remove(3)
	if list.Size() != 2 {
		t.Errorf("Expected list size 2 after removal, got %d", list.Size())
	}

	// Remove head element
	list.Remove(1)
	if list.Size() != 1 {
		t.Errorf("Expected list size 1 after removal, got %d", list.Size())
	}

	err = list.Remove(8)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestRemoveAt(t *testing.T) {
	list := LinkedList[int]{}

	// test empty list
	err := list.RemoveAt(0)
	if err == nil {
		t.Errorf("Expected error")
	}

	list.Add(1, 2, 3)
	list.RemoveAt(1)
	if list.Size() != 2 {
		t.Errorf("Expected list size 2 after RemoveAt, got %d", list.Size())
	}

	// test head element
	list.RemoveAt(0)
	if list.Size() != 1 {
		t.Errorf("Expected list size 1 after RemoveAt, got %d", list.Size())
	}

	// testing out of bounds
	err = list.RemoveAt(5)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestReplace(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)
	list.Replace(1, 42)
	val, _ := list.Get(1)
	if val != 42 {
		t.Errorf("Expected replaced value 42 at position 1, got %d", val)
	}

	// Replace head element
	list.Replace(0, 50)
	val, _ = list.Get(0)
	if val != 50 {
		t.Errorf("Expected replaced value 42 at position 1, got %d", val)
	}
}

func TestAddFront(t *testing.T) {
	list := LinkedList[int]{}
	list.Prepend(0)
	val, _ := list.Get(0)
	if val != 0 {
		t.Errorf("Expected first element 0 after AddFront, got %d", val)
	}
}

func TestPrint(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1, 2, 3)

	// Redirect the stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	list.Print()

	// Capture the output
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	// Verify the output
	expectedOutput := "1\n2\n3\n"
	actualOutput := buf.String()

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %v, Actual output: %v", expectedOutput, actualOutput)
	}
}

// test for the ToString method
func TestLinkedListToString(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{1, 2, 3}, "[1, 2, 3]"},
	}

	for _, test := range tests {
		list := LinkedList[int]{}
		for _, item := range test.input {
			list.Append(item)
		}
		result := list.ToString()
		if result != test.expected {
			t.Errorf("got %s, want %s", result, test.expected)
		}
	}
}

// test for the Equals method
func TestLinkedListEquals(t *testing.T) {
	tests := []struct {
		first    []int
		second   []int
		expected bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
	}

	for _, test := range tests {
		list1 := LinkedList[int]{}
		for _, item := range test.first {
			list1.Append(item)
		}
		list2 := LinkedList[int]{}
		for _, item := range test.second {
			list2.Append(item)
		}
		result := list1.Equals(&list2)
		if result != test.expected {
			t.Errorf("Equals(%v, %v) got %v, want %v", test.first, test.second, result, test.expected)
		}
	}

	// test for nil list
	list := LinkedList[int]{}
	result := list.Equals(nil)
	if result {
		t.Errorf("Equals(%v, %v) got %v, want %v", list, nil, result, false)
	}
}
