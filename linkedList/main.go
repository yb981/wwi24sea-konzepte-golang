package main

import (
	"ProgrammierkonzepteTask2/datastructures"
	"fmt"
)

func printValue(value int) {
	fmt.Print(value, ", ")
}

func doubleValue(value int) int {
	value = value * 2
	return value
}

func isOdd(value int) bool {
	return value%2 != 0
}

func Add(first int, second int) int {
	return first + second
}

func stringify(first string, second int) string {
	return first + fmt.Sprintf("%v", second)
}

func main() {
	/*
		mylist := new(datastructures.LinkedList[int])
		fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")
		mylist.Add(0, 1, 2, 3, 4, 5, 6)

		doubleValueList := mylist.Map(doubleValue)
		fmt.Println("List with doubled Value: ", doubleValueList.ToString())

		oddList := mylist.LazyFilter(isOdd)
		fmt.Println("This is the lazy list but not executed: ", oddList)
		fmt.Println("now executed")
		result := oddList.Execute()
		fmt.Println("Executed Lazy Filter: ", result.ToString())

		myQueue := new(datastructures.Queue[int])
		myQueue.Enqueue(0)
		myQueue.Enqueue(1)
		myQueue.Enqueue(2)
		myQueue.Enqueue(3)
		myQueue.Enqueue(4)
		myQueue.Enqueue(5)
		myQueue.Enqueue(6)
		fmt.Println("Add the Values 0, 1, 2, 3, 4, 5, 6")

		newQueue := myQueue.Map(doubleValue)
		fmt.Println("Queue with doubled Value: ", newQueue.ToString())

		oddQueue := myQueue.Filter(isOdd)
		fmt.Println("List with odd Value: ", oddQueue.ToString())

		//reduce operation on int list
		reduceList := new(datastructures.LinkedList[int])
		reduceList.Add(1, 2, 3, 4, 5, 6, 7)

			resultnew, _ := datastructures.Reduce[int, int](*reduceList, Add)
			fmt.Println("REduce FUnktion: ", resultnew)

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

		chainCalls := &datastructures.LinkedList[int]{}
		chainCalls.Add(1, 2, 3, 4, 5, 6, 7, 8)
		test := chainCalls.Map(doubleValue).Filter(is6).Map(doubleValue)
		fmt.Println("verkettung: ", test.ToString())
	*/
	functionalProgrammingDemo()
}

func functionalProgrammingDemo() {
	//init and fill list
	demoList := &datastructures.LinkedList[int]{}
	demoList.Add(1, 2, 3, 4, 5, 6, 7, 8)

	//init and fill stack
	demoStack := &datastructures.Stack[int]{}
	demoStack.PushAll(1, 2, 3, 4, 5, 6, 7, 8)

	//init and fill queue
	demoQueue := &datastructures.Queue[int]{}
	for i := 1; i < 9; i++ {
		demoQueue.Enqueue(i)
	}

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("myList: ", demoList.ToString())
	fmt.Println("myStack: ", demoStack.ToString())
	fmt.Println("myQueue: ", demoQueue.ToString())
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Demo der ForEach Methode anhand einer Print Funktion: ")
	fmt.Println()
	fmt.Print("ForEach bei der Liste: ")
	demoList.ForEach(printValue)
	fmt.Println()
	fmt.Print("ForEach beim Stack: ")
	demoStack.ForEach(printValue)
	fmt.Println()
	fmt.Print("ForEach bei der Queue: ")
	demoQueue.ForEach(printValue)
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Demo der Filter Methode anhand einer isOdd Funktion: ")
	fmt.Println()
	fmt.Print("Filter bei der Liste: ")
	fmt.Println(demoList.Filter(isOdd).ToString())
	fmt.Print("Filter beim Stack: ")
	fmt.Println(demoStack.Filter(isOdd).ToString())
	fmt.Print("Filter bei der Queue: ")
	fmt.Println(demoQueue.Filter(isOdd).ToString())
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Demo der Map Methode anhand einer doubleValue Funktion: ")
	fmt.Println()
	fmt.Print("Map bei der Liste: ")
	fmt.Println(demoList.Map(doubleValue).ToString())
	fmt.Print("Map beim Stack: ")
	fmt.Println(demoStack.Map(doubleValue).ToString())
	fmt.Print("Map bei der Queue: ")
	fmt.Println(demoQueue.Map(doubleValue).ToString())
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Demo der Reduce Methode anhand einer Add Funktion: ")
	fmt.Println()
	fmt.Print("Reduce bei der Liste: ")
	result, _ := demoList.Reduce(Add)
	fmt.Println(result)
	fmt.Print("Reduce beim Stack: ")
	result, _ = demoStack.Reduce(Add)
	fmt.Println(result)
	fmt.Print("Reduce bei der Queue: ")
	result, _ = demoQueue.Reduce(Add)
	fmt.Println(result)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Demo der Lazy Variante bei der Filter Funktion mit isOdd auf einer Liste: ")
	fmt.Println()
	unexecutedLazyFilter := demoList.LazyFilter(isOdd)
	fmt.Println("Unexecuted LazyFilter: ", unexecutedLazyFilter.Operations)
	fmt.Println("Executed LazyFilter: ", unexecutedLazyFilter.Execute().ToString())
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Demo der Lazy Variante bei der Map Funktion mit doubleValue auf einer Liste: ")
	fmt.Println()
	unexecutedLazyMap := demoList.LazyMap(doubleValue)
	fmt.Println("Unexecuted LazyMap: ", unexecutedLazyMap.Operations)
	fmt.Println("Executed LazyMap: ", unexecutedLazyMap.ExecuteMap().ToString())
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------")
}
