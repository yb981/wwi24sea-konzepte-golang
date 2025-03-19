package main

import (
	"concurrency/concurrency"
	"fmt"
)

func doubleValue(input int) int {
	return input * 2
}

func add(input1 int, input2 int) int {
	return input1 + input2
}

func main() {
	myList := &concurrency.ArrayList[int]{}
	for i := range 1000000 {
		myList.Append(i)
	}

	reduceList := myList.ParallelReduce(add)

	fmt.Println(reduceList)
}
