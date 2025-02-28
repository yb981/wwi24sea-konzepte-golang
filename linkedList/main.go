package main

import (
	"fmt"
)

func main() {
	mylist := new(LinkedList[int])
	fmt.Println("Add the Values 10, 11, 12, 13, 14, 15")
	mylist.add(10, 11, 12, 13, 14, 15)

	fmt.Println("Element at index 2:", mylist.getValue(2))

	fmt.Println("Länge der Liste", mylist.size())

	fmt.Println("Liste ist leer", mylist.isEmpty())

	fmt.Println("Länge der voll", mylist.isFull())

	fmt.Println("Removing Element with Value 14")
	mylist.remove(14)

	fmt.Println("Removing Element at position 1")
	mylist.removeAt(1)

	fmt.Println("Inserting 999 at position 0")
	mylist.insert(0, 999)

	fmt.Println("Replacing Element at position 1 with value 100")
	mylist.replace(1, 100)

	mylist.print()

	fmt.Println("----------------------------------------")
	mystack := new(Stack[int])
	mystack.push(10)
	mystack.push(11)
	mystack.push(12)
	mystack.pop()
	mystack.pop()
	fmt.Println("Top Element: ", mystack.peek())
}
