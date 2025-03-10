package main

import (
	"ProgrammierkonzepteTask2/datastructures"
	"fmt"
)

func main() {
	mylist := new(datastructures.LinkedList[int])
	fmt.Println("Add the Values 10, 11, 12, 13, 14, 15")
	mylist.Add(10, 11, 12, 13, 14, 15)
	mylist.Print()
	fmt.Println("----------------------------------------")
	val, err := mylist.Get(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 2:", val)
	}
	
	fmt.Println("Länge der Liste", mylist.Size())

	fmt.Println("Liste ist leer", mylist.IsEmpty())

	fmt.Println("Länge der voll", mylist.IsFull())

	fmt.Println("Removing Element with Value 14")
	mylist.Remove(14)

	fmt.Println("Removing Element at position 1")
	mylist.RemoveAt(1)

	fmt.Println("Inserting 999 at position 0")
	mylist.Insert(0, 999)

	fmt.Println("Replacing Element at position 1 with value 100")
	mylist.Replace(1, 100)

	mylist.Print()

	fmt.Println("----------------------------------------")
	mystack := new(datastructures.Stack[int])
	mystack.Push(10)
	mystack.Push(11)
	mystack.Push(12)
	mystack.Pop()
	mystack.Pop()
	value, err := mystack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
			fmt.Println("Top Element: ", value)
	}

	fmt.Println("----------------------------------------")
	myqueue := new(datastructures.Queue[int])
	myqueue.Enqueue(10)
	myqueue.Enqueue(11)
	myqueue.Enqueue(12)
	value, err = myqueue.Dequeue()
	if err != nil {
			fmt.Println("Error:", err)
	} else {
			fmt.Println("Pop first element should be 10:", value)
	}
	fmt.Println(myqueue.Dequeue())
	fmt.Println(myqueue.Dequeue())
}
