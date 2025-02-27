package main

import(
	//"fmt"
)

func main(){
	mylist := new(LinkedList[int])
	mylist.addBack(10)
	mylist.addBack(29)
	mylist.addBack(222)
	mylist.print()
}
