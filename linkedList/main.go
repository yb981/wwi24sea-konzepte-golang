package main

import (
	"fmt"
)

func main() {
	mylist := new(LinkedList[int])
	mylist.addBack(12)
	mylist.addBack(13)
	mylist.addBack(14)
	mylist.insert(2, 888)
	fmt.Println("Länge der Liste", mylist.size())
	fmt.Println("Liste ist leer", mylist.isEmpty())
	fmt.Println("Länge der voll", mylist.isFull())
	mylist.print()
}
