package main

import (
	"concurrency/concurrency"
	"testing"
)

// Benchmark für die sequenzielle Reduce-Funktion
func BenchmarkReduce(b *testing.B) {
	myList := &concurrency.ArrayList[int]{}
	for i := 0; i < 100000000; i++ {
		myList.Append(i)
	}

	b.ResetTimer() // Starte den Timer für den eigentlichen Benchmark
	for i := 0; i < b.N; i++ {
		_, _ = myList.Reduce(add)
	}
}

// Benchmark für die parallele Reduce-Funktion
func BenchmarkParallelReduce(b *testing.B) {
	myList := &concurrency.ArrayList[int]{}
	for i := 0; i < 100000000; i++ {
		myList.Append(i)
	}

	b.ResetTimer() // Starte den Timer für den eigentlichen Benchmark
	for i := 0; i < b.N; i++ {
		_, _ = myList.ParallelReduce(add)
	}
}

func BenchmarkMap(b *testing.B) {
	myList := &concurrency.ArrayList[int]{}
	for i := 0; i < 100000000; i++ {
		myList.Append(i)
	}

	b.ResetTimer() // Starte den Timer für den eigentlichen Benchmark
	for i := 0; i < b.N; i++ {
		_, _ = myList.Map(doubleValue)
	}
}

func BenchmarkParallelMap(b *testing.B) {
	myList := &concurrency.ArrayList[int]{}
	for i := 0; i < 100000000; i++ {
		myList.Append(i)
	}

	b.ResetTimer() // Starte den Timer für den eigentlichen Benchmark
	for i := 0; i < b.N; i++ {
		_, _ = myList.ParallelMap(doubleValue)
	}
}
