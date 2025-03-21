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

func (al *ArrayList[T]) ParallelReduce(workerNum int, operation func(a, b T) T) (T, error) {
	if al.isEmpty() {
		var zero T
		return zero, errors.New("Reduce not possible for empty List")
	}

	n := len(al.list)
	if workerNum > n {
		workerNum = n // Nicht mehr Worker als Elemente
	}

	chunk := n / workerNum
	results := make(chan T, workerNum) // Channel für Zwischenergebnisse

	var wg sync.WaitGroup

	// Arbeiter starten
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * chunk
			end := start + chunk
			if i == workerNum-1 { // Letzter Worker bekommt den Rest
				end = n
			}

			result := al.list[start]
			for j := start + 1; j < end; j++ {
				result = operation(result, al.list[j])
			}
			results <- result
		}(i)
	}

	wg.Wait()
	close(results)

	// Endgültige Reduktion
	finalResult := <-results
	for res := range results {
		finalResult = operation(finalResult, res)
	}

	return finalResult, nil
}

func (al *ArrayList[T]) ParallelReduceJobChannel(workerNum int, operation func(a, b T) T) (T, error) {
	if al.isEmpty() {
		var zero T
		return zero, errors.New("Reduce not possible for empty List")
	}

	n := len(al.list)
	if workerNum > n {
		workerNum = n
	}

	jobs := make(chan [2]int, workerNum) // Channel für Start- und Endindex-Paare
	results := make(chan T, workerNum)   // Channel für Zwischenergebnisse

	var wg sync.WaitGroup

	// Worker starten
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				start, end := job[0], job[1]
				result := al.list[start]
				for j := start + 1; j < end; j++ {
					result = operation(result, al.list[j])
				}
				results <- result
			}
		}()
	}

	// Jobs an den Channel senden
	chunk := n / workerNum
	for i := 0; i < workerNum; i++ {
		start := i * chunk
		end := start + chunk
		if i == workerNum-1 { // Letzter Worker bekommt den Rest
			end = n
		}
		jobs <- [2]int{start, end}
	}
	close(jobs)

	// Warten auf die Worker
	wg.Wait()
	close(results)

	// Endgültige Reduktion
	finalResult := <-results
	for res := range results {
		finalResult = operation(finalResult, res)
	}

	return finalResult, nil
}

