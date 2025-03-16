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

func Add(first int, second int) string {
	return string(first + second)
}

func doubleValueVLukas(value int) int {
	value = value * 2
	return value
}

func main() {
	mylist := new(datastructures.LinkedList[int])
	fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")
	mylist.Add(0, 1, 2, 3, 4, 5, 6)

	doubleValueList := mylist.Map(doubleValue[int])
	fmt.Println("List with doubled Value: ", doubleValueList.ToString())

	oddList := mylist.LazyFilter(isOdd[int])
	fmt.Println("List with odd Value: ", oddList.ToString())

	myQueue := new(datastructures.Queue[int])
	fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")
	for i := 0; i < 7; i++ {
		myQueue.Enqueue(i)
	}

	newQueue := myQueue.Map(doubleValue[int])
	fmt.Println("Queue with doubled Value: ", newQueue.ToString())

	oddQueue := myQueue.Filter(isOdd[int])
	fmt.Println("List with odd Value: ", oddQueue.ToString())

	//reduce operation on int list
	reduceList := new(datastructures.LinkedList[int])
	reduceList.Add(1, 2, 3, 4, 5, 6, 7)
	//result, _ := datastructures.Reduce[int, string](*reduceList, Add)
	//fmt.Println("REduce FUnktion: ", result)

	myListForMapping := new(datastructures.LinkedList[int])
	fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")
	myListForMapping.Add(0, 1, 2, 3, 4, 5, 6)

	mappedList := myListForMapping.MapVariant(
		func(x int) any { return x * 2 },
		datastructures.LinkedListType,
	).(*datastructures.LinkedList[any])

	mappedQueue := myListForMapping.MapVariant(
		func(x int) any { return x * 2 },
		datastructures.QueueType,
	).(*datastructures.Queue[any])

	mappedStack := myListForMapping.MapVariant(
		func(x int) any { return x * 2 },
		datastructures.StackType,
	).(*datastructures.Stack[any])

	fmt.Println(mappedList.ToString())
	fmt.Println(mappedQueue.ToString())
	fmt.Println(mappedStack.ToString())
}
