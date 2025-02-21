package main

import "fmt"

type Stack []float64

func (s *Stack) Push(input float64) {
	*s = append(*s, input)
}

func (s *Stack) Pop() float64 {
	if len(*s) == 0 {
		return 0
	}

	lastValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return lastValue
}

func (s *Stack) Print() {
	for _, v := range *s {
		fmt.Print(v, " ")
	}
}
