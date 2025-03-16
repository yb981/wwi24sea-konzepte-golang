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

func Add(first int, second int) int {
	return first + second
}

func stringify(first string, second int) string {
	return first + fmt.Sprintf("%v", second)
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
	fmt.Println("This is the lazy list but not executed: ", oddList)
	fmt.Println("now executed")
	result := oddList.Execute()
	fmt.Println("Executed Lazy Filter: ", result.ToString())

	myQueue := new(datastructures.Queue[int])
	fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")

	newQueue := myQueue.Map(doubleValue[int])
	fmt.Println("Queue with doubled Value: ", newQueue.ToString())

	oddQueue := myQueue.Filter(isOdd[int])
	fmt.Println("List with odd Value: ", oddQueue.ToString())

	//reduce operation on int list
	reduceList := new(datastructures.LinkedList[int])
	reduceList.Add(1, 2, 3, 4, 5, 6, 7)
	/*
		resultnew, _ := datastructures.Reduce[int, int](*reduceList, Add)
		fmt.Println("REduce FUnktion: ", resultnew)
	*/
	ergebnis, _ := reduceList.Reduce(Add)
	fmt.Println("Reduce ergebnis:", ergebnis)

	reduceQueue := &datastructures.Queue[int]{}
	reduceQueue.Enqueue(1)
	reduceQueue.Enqueue(2)
	reduceQueue.Enqueue(3)
	value, _ := reduceQueue.Reduce(Add)
	fmt.Println("REduzierte Queue: ", value)

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
