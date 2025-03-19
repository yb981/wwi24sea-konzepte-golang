package main

import (
	"fmt"
)

func doubleValue(input int) int {
	return input * 2
}

func main() {
	myList := &ArrayList[int]{}
	for i := range 10 {
		myList.Append(i)
	}
	parallelList := myList.parallelMap(doubleValue)

	fmt.Println(parallelList)
}
