package main

import (
	"concurrency/concurrency"
	"fmt"
)

func doubleValue(input int) int {
	return input * 2
}

func ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if m > 0 && n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}

func add(input1 int, input2 int) int {
	return input1 + input2
}

func main() {
	myList := &concurrency.ArrayList[int]{}
	for i := range 10000000 {
		myList.Append(i)
	}

	parMapList, err := myList.ParallelReduce(add)
	mapList, err := myList.Reduce(add)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parMapList)
	fmt.Println(mapList)
}