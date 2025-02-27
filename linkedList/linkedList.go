package main

type LinkedList struct{
	head* Node 
}

func(list* LinkedList)addFront(data int){
	newNode := &Node{data : data, next: list.head}
	list.head = newNode
}

func(list* LinkedList)print()int{
	return list.head.data
}