// main.go
//
// Startet die Demo der Datenstrukturen
//
// Author: Till Burdorf, Lukas Gröning
// Date: 30.02.2025

package main

import (
	"fmt"

	"github.com/yb981/wwi24sea-konzepte-golang/datastructures"
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

func main() {
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
