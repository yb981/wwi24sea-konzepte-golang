package main

import (
	"datastructures/datastructure"
	"fmt"
)

func main() {
	mylist := new(datastructure.LinkedList[int])
	fmt.Println("Add the Values 10, 11, 12, 13, 14, 15")
	mylist.Add(10, 11, 12, 13, 14, 15)

	fmt.Println("Element at index 2:", mylist.GetValue(2))

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
	mystack := new(datastructure.Stack[int])
	mystack.Push(10)
	mystack.Push(11)
	mystack.Push(12)
	mystack.Pop()
	mystack.Pop()
	fmt.Println("Top Element: ", mystack.Peek())
}
