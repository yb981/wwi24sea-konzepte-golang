package datastructures

import (
	"fmt"
	"reflect"
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
	original := &LinkedList[int]{}

	list.Add(1, 2, 3)
	original.Add(1, 2, 3)

	invalidType := CollectionType(999) // Ungültiger Typ
	result := list.MapVariant(func(x int) any { return x * 2 }, invalidType)

	if result != nil {
		t.Errorf("Expected nil for invalid CollectionType, but got %v", result)
	}

	resultLL := list.MapVariant(func(x int) any { return x * 2 }, LinkedListType).(*LinkedList[any])
	if !list.Equals(original) {
		t.Errorf("Expected original list not to change, but got %v", list)
	}
	expected := LinkedList[any]{}
	expected.Append(2)
	expected.Append(4)
	expected.Append(6)

	if !resultLL.Equals(&expected) {
		t.Errorf("Expected %v, but got %v", expected.ToString(), resultLL.ToString())
	}
}

// Testdatei für ForEachMethode
func TestForEach(t *testing.T) {

	expected := []int{1, 2, 3} //slice, den wir erwarten (slice= array mit flexibler größe)

	// Test für LinkedList
	ll := LinkedList[int]{}
	ll.Add(1, 2, 3)

	// LinkedList in Slice
	result := []int{}
	ll.ForEach(func(value int) {
		result = append(result, value)
	})

	// Vergleich der Längen
	if len(result) != len(expected) {
		t.Errorf("LinkedList ForEach failed. Length mismatch. Expected %v, got %v", len(expected), len(result))
	}

	// Vergleich von jedem i
	for i := range expected {
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

func TestFilterVariantLinkedList(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	isEven := func(x int) bool { return x%2 == 0 }

	result := list.FilterVariant(isEven, LinkedListType)
	filteredList, ok := result.(*LinkedList[any])
	if !ok {
		t.Errorf("Expected result to be of type LinkedList, but got %T", result)
	}

	expected := LinkedList[any]{}
	expected.Append(2)

	if !reflect.DeepEqual(filteredList, &expected) {
		t.Errorf("Expected %v, but got %v", expected, filteredList)
	}
}

func TestFilterVariantQueue(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	isEven := func(x int) bool { return x%2 == 0 }

	result := list.FilterVariant(isEven, QueueType)
	filteredQueue, ok := result.(*Queue[any])
	if !ok {
		t.Errorf("Expected result to be of type Queue, but got %T", result)
	}

	expected := Queue[any]{}
	expected.Enqueue(2)

	if !reflect.DeepEqual(filteredQueue, &expected) {
		t.Errorf("Expected %v, but got %v", expected, filteredQueue)
	}
}

func TestFilterVariantStack(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	isEven := func(x int) bool { return x%2 == 0 }

	result := list.FilterVariant(isEven, StackType)
	filteredStack, ok := result.(*Stack[any])
	if !ok {
		t.Errorf("Expected result to be of type Stack, but got %T", result)
	}

	expected := Stack[any]{}
	expected.Push(2)

	if !reflect.DeepEqual(filteredStack, &expected) {
		t.Errorf("Expected %v, but got %v", expected, filteredStack)
	}
}

func TestFilterVariantEmptyList(t *testing.T) {
	list := LinkedList[int]{}

	isEven := func(x int) bool { return x%2 == 0 }

	result := list.FilterVariant(isEven, LinkedListType)
	filteredList, ok := result.(*LinkedList[any])
	if !ok {
		t.Errorf("Expected result to be of type LinkedList, but got %T", result)
	}

	expected := LinkedList[any]{}

	if !reflect.DeepEqual(filteredList, &expected) {
		t.Errorf("Expected %v, but got %v", expected, filteredList)
	}
}

func TestFilterVariantNoMatch(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(1)
	list.Append(3)
	list.Append(5)

	isEven := func(x int) bool { return x%2 == 0 }

	result := list.FilterVariant(isEven, LinkedListType)
	filteredList, ok := result.(*LinkedList[any])
	if !ok {
		t.Errorf("Expected result to be of type LinkedList, but got %T", result)
	}

	expected := LinkedList[any]{}

	if !reflect.DeepEqual(filteredList, &expected) {
		t.Errorf("Expected %v, but got %v", expected, filteredList)
	}
}

func TestFilterVariantUnknownType(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	isEven := func(x int) bool { return x%2 == 0 }

	result := list.FilterVariant(isEven, CollectionType(-1)) // Unknown collection type

	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
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
