package datastructures

import (
	"fmt"
	"testing"
)

/*
Commands für Test
Zuerst:
cd .\datastructures\
go test -v
go test -cover
go test -coverprofile=coverage
go tool cover -html=coverage -o coverage.html
*/

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
		testList = *testList.Filter(isEvenFunc)

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

func TestLazyFilter(t *testing.T) {
	isEvenFunc := func(a int) bool { return a%2 == 0 }

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4}},
		{[]int{10, 15, 20}, []int{10, 20}},
		{[]int{5, 7, 9}, []int{}}, // keine geraden Zahlen
	}

	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}

		lazyResult := testList.LazyFilter(isEvenFunc).Execute()

		var expectedList LinkedList[int]
		for _, v := range test.expected {
			expectedList.Append(v)
		}

		if !lazyResult.Equals(&expectedList) {
			t.Errorf("LazyFilter failed. Expected %v, got %v", expectedList.ToString(), lazyResult.ToString())
		}
	}
}
func TestMap(t *testing.T) {
	// Funktion zum Verdoppeln einer Zahl
	doubleFunc := func(a int) int { return a * 2 }

	// Funktion zur Typumwandlung von int nach string
	toStringFunc := func(a int) string { return fmt.Sprintf("Num:%d", a) }

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{-3, -2, -1}, []int{-6, -4, -2}},
		{[]int{0, 5, 10}, []int{0, 10, 20}},
	}

	// Test für LinkedList, Queue und Stack
	for _, test := range tests {
		var list LinkedList[int]
		for _, v := range test.input {
			list.Append(v)
		}
		mappedList := *list.Map(doubleFunc)

		var expectedList LinkedList[int]
		for _, v := range test.expected {
			expectedList.Append(v)
		}

		if !mappedList.Equals(&expectedList) {
			t.Errorf("LinkedList Map failed. Expected %v, got %v", expectedList.ToString(), mappedList.ToString())
		}

		// Gleiches für Queue
		var queue Queue[int]
		for _, v := range test.input {
			queue.Enqueue(v)
		}
		mappedQueue := *queue.Map(doubleFunc)

		var expectedQueue Queue[int]
		for _, v := range test.expected {
			expectedQueue.Enqueue(v)
		}

		if !mappedQueue.Equals(&expectedQueue) {
			t.Errorf("Queue Map failed. Expected %v, got %v", expectedQueue.ToString(), mappedQueue.ToString())
		}

		// Gleiches für Stack
		var stack Stack[int]
		for _, v := range test.input {
			stack.Push(v)
		}
		mappedStack := *stack.Map(doubleFunc)

		var expectedStack Stack[int]
		for _, v := range test.expected {
			expectedStack.Push(v)
		}

		if !mappedStack.Equals(&expectedStack) {
			t.Errorf("Stack Map failed. Expected %v, got %v", expectedStack.ToString(), mappedStack.ToString())
		}
	}

	// Test für generische Map-Funktion mit Typumwandlung (int zu string)
	intInput := []int{1, 2, 3}
	expectedStrings := []string{"Num:1", "Num:2", "Num:3"}

	var intList LinkedList[int]
	for _, v := range intInput {
		intList.Append(v)
	}

	stringList := Map(intList, toStringFunc)

	var expectedStringList LinkedList[string]
	for _, v := range expectedStrings {
		expectedStringList.Append(v)
	}

	if !stringList.Equals(&expectedStringList) {
		t.Errorf("Generic Map (int -> string) failed. Expected %v, got %v", expectedStringList.ToString(), stringList.ToString())
	}
}

func TestLazyMap(t *testing.T) {
	// Funktion zum Quadrieren einer Zahl
	squareFunc := func(a int) int { return a * a }

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 4, 9}},
		{[]int{-3, -2, -1}, []int{9, 4, 1}},
		{[]int{0, 5, 10}, []int{0, 25, 100}},
	}

	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}

		lazyResult := testList.LazyMap(squareFunc).ExecuteMap()

		var expectedList LinkedList[int]
		for _, v := range test.expected {
			expectedList.Append(v)
		}

		if !lazyResult.Equals(&expectedList) {
			t.Errorf("LazyMap failed. Expected %v, got %v", expectedList.ToString(), lazyResult.ToString())
		}
	}
}


func TestReduce(t *testing.T) {
	// Testfunktion: Summe berechnen
	sumFunc := func(a, b int) int { return a + b }

	tests := []struct {
		input    []int
		expected int
		hasError bool
	}{
		{[]int{1, 2, 3}, 6, false},
		{[]int{10, 20, 30}, 60, false},
		{[]int{}, 0, true}, // Leere Liste sollte einen Fehler werfen
	}

	// Test für LinkedList Reduce
	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}

		result, err := testList.Reduce(sumFunc)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected an error for empty list, but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		}
	}

	// Test für Queue Reduce
	for _, test := range tests {
		queue := &Queue[int]{}
		for _, v := range test.input {
			queue.Enqueue(v)
		}

		result, err := queue.Reduce(sumFunc)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected an error for empty queue, but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		}
	}

	// Test für Stack Reduce
	for _, test := range tests {
		stack := &Stack[int]{}
		for _, v := range test.input {
			stack.Push(v)
		}

		result, err := stack.Reduce(sumFunc)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected an error for empty stack, but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		}
	}
}
