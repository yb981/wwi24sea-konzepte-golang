package main

import(
	"fmt"
)

func main(){
	mylist := new(LinkedList)
	mylist.addFront(10)
	fmt.Println(mylist.print())
}
