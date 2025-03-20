package concurrency

import (
	"errors"
	"sync"
)

type ArrayList[T comparable] struct {
	list []T
}

func (al *ArrayList[T]) Append(value T) {
	al.list = append(al.list, value)
}

func (al *ArrayList[T]) isEmpty() bool {
	return len(al.list) == 0
}

func (al *ArrayList[T]) ParallelMap(workerNum int, operation func(a T) T) (ArrayList[T], error) {
	if al.isEmpty() {
		var zero ArrayList[T]
		return zero, errors.New("Reduce not possible for empty List")
	}
	chunk := len(al.list) / workerNum // split the array into chunks which equal the number of the cpus available
	output := &ArrayList[T]{list: make([]T, len(al.list))}
	var wg sync.WaitGroup

	for i := 0; i < workerNum; i++ {
		wg.Add(1) // increase waitgroup counter for every iteration
		// start a GO Routine for every CPU Core available in the system
		go func(i int) {
			defer wg.Done()
			start := i * chunk
			end := start + chunk

			// last chunk could be smaller
			if i == workerNum-1 {
				end = len(al.list)
			}

			for j := start; j < end; j++ {
				output.list[j] = operation(al.list[j])
			}
		}(i)
	}
	wg.Wait()

	return *output, nil
}

func (al *ArrayList[T]) ParallelReduce(workerNum int, operation func(a, b T) T) (T, error) {
	if al.isEmpty() {
		var zero T
		return zero, errors.New("Reduce not possible for empty List")
	}
	chunk := len(al.list) / workerNum // split the array into chunks which equal the number of the cpus available
	output := &ArrayList[T]{list: make([]T, workerNum)}
	var wg sync.WaitGroup

	for i := 0; i < workerNum; i++ {
		wg.Add(1) // increase waitgroup counter for every iteration
		// start a GO Routine for every CPU Core available in the system
		go func(i int) {
			defer wg.Done()
			start := i * chunk
			end := start + chunk

			// last chunk could be smaller
			if i == workerNum-1 {
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

func (al *ArrayList[T]) Map(operation func(T) T) (*ArrayList[T], error) {
	if al.isEmpty() {
		var zero *ArrayList[T]
		return zero, errors.New("empty list does not allow map function")
	}
	output := &ArrayList[T]{list: make([]T, len(al.list))}
	for i, v := range al.list {
		output.list[i] = operation(v)
	}
	return output, nil
}

func (al *ArrayList[T]) Reduce(f func(a, b T) T) (T, error) {
	if al.isEmpty() {
		var zero T
		return zero, errors.New("empty list does not allow reduce operation")
	}

	result := al.list[0]

	for i := 1; i < len(al.list); i++ {
		result = f(result, al.list[i])
	}

	return result, nil
}
