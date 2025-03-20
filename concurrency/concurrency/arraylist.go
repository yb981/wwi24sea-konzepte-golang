package concurrency

import (
	"errors"
	"runtime"
	"sync"
)

type ArrayList[T comparable] struct {
	list []T
}

func (al *ArrayList[T]) Append(value T) {
	al.list = append(al.list, value)
}

func (al *ArrayList[T]) ParallelMap(f func(T) T) *ArrayList[T] {
	//init waitgroup to sync the go Routines
	var wg sync.WaitGroup

	output := &ArrayList[T]{list: make([]T, len(al.list))}

	for i, v := range al.list {
		wg.Add(1)             // increase waitgroup counter for every iteration
		go func(i int, v T) { //start go Routine
			defer wg.Done() // reduce waitgroup counter after function is executed
			output.list[i] = f(v)
		}(i, v)
	}

	wg.Wait() // wait till all functions are executed
	return output
}

func (al *ArrayList[T]) ParallelReduce(operation func(a, b T) T) (T, error) {
	if len(al.list) == 0 {
		var zero T
		return zero, errors.New("Reduce not possible for empty List!")
	}
	chunk := len(al.list) / runtime.NumCPU() // split the array into chunks which equal the number of the cpus available
	output := &ArrayList[T]{list: make([]T, runtime.NumCPU())}
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1) // increase waitgroup counter for every iteration
		// start a GO Routine for every CPU Core available in the system
		go func(i int) {
			defer wg.Done()
			start := i * chunk
			end := start + chunk

			// last chunk could be smaller
			if i == runtime.NumCPU()-1 {
				end = len(al.list)
			}

			result := al.list[start]
			for j := start + 1; j < end; j++ {
				result = operation(result, al.list[j])
			}
			output.list[i] = result
		}(i)
	}
	wg.Wait()

	finalResult := output.list[0]
	for i := 1; i < len(output.list); i++ {
		finalResult = operation(finalResult, output.list[i])
	}

	return finalResult, nil
}
