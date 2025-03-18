package main

type ArrayList[T comparable] struct{
	list[] T
}

func(al* ArrayList[T]) Append(value T){
	al.list = append(al.list, value)
}

func(al* ArrayList[T]) Insert(value T){
}
