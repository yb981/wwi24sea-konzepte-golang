package main

import (
	"concurrency/concurrency"
	"fmt"
	"runtime"
)

/*
func ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if m > 0 && n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}
*/

func doubleValue(input int) int {
	return input * 2
}

func add(input1 int, input2 int) int {
	return input1 + input2
}

func concatToString(a string, b string) string {
	return a + b
}

func main() {
	createDemoOutput()
}

func createDemoOutput() {
	// initialize list and fill with values
	myList := &concurrency.ArrayList[int]{}
	for i := range 10 {
		myList.Append(i)
	}
	fmt.Println("My List: ", myList)

	parMapList, err := myList.ParallelMap(runtime.NumCPU(), doubleValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	mapList, err := myList.Map(doubleValue)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Test sequencial Map with double Value Function. Result: ", *mapList)
	fmt.Println("Test parallel Map with double Value Function. Result:   ", parMapList)

	parReduce, err := myList.ParallelReduce(runtime.NumCPU(), add)
	if err != nil {
		fmt.Println(err)
		return
	}

	reduce, err := myList.Reduce(add)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Test sequencial Reduce with Add Function. Result: ", reduce)
	fmt.Println("Test parallel Reduce with Add Function. Result:   ", parReduce)
}