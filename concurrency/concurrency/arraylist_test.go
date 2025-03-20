package concurrency

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := &ArrayList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if len(list.list) != 3 {
		t.Errorf("Expected length 3, got %d", len(list.list))
	}
}

func TestMap(t *testing.T) {
	list := &ArrayList[int]{list: []int{1, 2, 3, 4, 5}}
	mapped, err := list.Map(func(x int) int { return x * 2 })
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []int{2, 4, 6, 8, 10}
	for i, v := range mapped.list {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestReduce(t *testing.T) {
	list := &ArrayList[int]{list: []int{1, 2, 3, 4, 5}}
	sum, err := list.Reduce(func(a, b int) int { return a + b })
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}
}

func TestParallelMap(t *testing.T) {
	list := &ArrayList[int]{list: []int{1, 2, 3, 4, 5}}
	mapped, err := list.ParallelMap(2, func(x int) int { return x * 2 })
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []int{2, 4, 6, 8, 10}
	for i, v := range mapped.list {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestParallelMapEmptyList(t *testing.T) {
	list := &ArrayList[int]{} 
	mapped, err := list.ParallelMap(2, func(x int) int { return x * 2 })
	
	if err == nil {
		t.Errorf("Expected error for empty list, got nil")
	}

	if mapped.list != nil && len(mapped.list) != 0 {
		t.Errorf("Expected empty list, got %v", mapped.list)
	}
}

func TestParallelReduce(t *testing.T) {
	list := &ArrayList[int]{list: []int{1, 2, 3, 4, 5}}
	sum, err := list.ParallelReduce(2, func(a, b int) int { return a + b })
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}
}

func TestParallelReduceWorkerLimit(t *testing.T) {
	list := &ArrayList[int]{list: []int{1, 2, 3}}
	workerNum := 10   
	result, err := list.ParallelReduce(workerNum, func(a, b int) int { return a + b })

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 6
	if result != expected {
		t.Errorf("Expected sum %d, got %d", expected, result)
	}
}


func TestMapEmptyList(t *testing.T) {
	list := &ArrayList[int]{}
	_, err := list.Map(func(x int) int { return x * 2 })
	if err == nil {
		t.Errorf("Expected error for empty list, got nil")
	}
}

func TestReduceEmptyList(t *testing.T) {
	list := &ArrayList[int]{}
	_, err := list.Reduce(func(a, b int) int { return a + b })
	if err == nil {
		t.Errorf("Expected error for empty list, got nil")
	}
}

func TestParallelReduceEmptyList(t *testing.T) {
	list := &ArrayList[int]{}
	_, err := list.ParallelReduce(2, func(a, b int) int { return a + b })
	if err == nil {
		t.Errorf("Expected error for empty list, got nil")
	}
}
