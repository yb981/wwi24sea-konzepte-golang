package main

import (
	"ProgrammierkonzepteTask2/datastructures"
	"fmt"
)

func doubleValue[T any](value int) int {
	value = value * 2
	return value
}

func isOdd[T any](value int) bool {
	return value%2 != 0
}

func main() {
	mylist := new(datastructures.LinkedList[int])
	fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")
	mylist.Add(0, 1, 2, 3, 4, 5, 6)

	doubleValueList := mylist.Map(doubleValue[int])
	fmt.Println(doubleValueList.ToString())

	filteredList := mylist.Filter(isOdd[int])
	fmt.Println(filteredList.ToString())

	myQueue := new(datastructures.Queue[int])
	fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")
	for i := 0; i < 7; i++ {
		myQueue.Enqueue(i)
	}

	newQueue := myQueue.Map(doubleValue[int])
	fmt.Println(newQueue.ToString())
}
