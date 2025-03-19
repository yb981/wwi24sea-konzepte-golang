package main

import (
	"sync"
)

type ArrayList[T comparable] struct {
	list []T
}

func (al *ArrayList[T]) Append(value T) {
	al.list = append(al.list, value)
}


func (al *ArrayList[T]) parallelMap(f func(T) T) *ArrayList[T] {
	var wg sync.WaitGroup
	newList := &ArrayList[T]{list: make([]T, len(al.list))}

	for i, v := range al.list {
		wg.Add(1)
		go func(i int, v T) {
			defer wg.Done()
			newList.list[i] = f(v)
		}(i, v)
	}
	wg.Wait()
	return newList
}
