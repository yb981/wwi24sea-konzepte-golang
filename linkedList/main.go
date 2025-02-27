package main

import (
	"fmt"
)

func main() {
	mylist := new(LinkedList[int])
	mylist.add(10, 11, 12, 13, 14, 15)
	fmt.Println("Element at index 2:", mylist.get(2))
	fmt.Println("Länge der Liste", mylist.size())
	fmt.Println("Liste ist leer", mylist.isEmpty())
	fmt.Println("Länge der voll", mylist.isFull())
	mylist.print()
}
