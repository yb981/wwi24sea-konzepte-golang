package main

import (
	"fmt"
	"math"
)

func (stack *Stack) Add() {
	switch {
	case len(stack.data) >= 2:
		//Adding last two numbers of stack
		sum := stack.data[len(stack.data)-2] + stack.data[len(stack.data)-1]

		//Updating stack
		stack.Pop()
		stack.data[len(stack.data)-1] = sum

	//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Sub() {
	switch {
	case len(stack.data) >= 2:
		//Subtracting last two numbers of stack
		sub := stack.data[len(stack.data)-2] - stack.data[len(stack.data)-1]

		//Updating stack
		stack.Pop()
		stack.data[len(stack.data)-1] = sub

	//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Mult() {
	switch {
	case len(stack.data) >= 2:
		//Multiplying last two numbers of stack
		mult := stack.data[len(stack.data)-2] * stack.data[len(stack.data)-1]

		//Updating Stack
		stack.Pop()
		stack.data[len(stack.data)-1] = mult

		//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Div() {
	switch {
	case len(stack.data) >= 2:
		switch {
		case stack.data[len(stack.data)-1] != 0:
			//dividing last two numbers of stack
			div := stack.data[len(stack.data)-2] / stack.data[len(stack.data)-1]

			//Updating stack
			stack.Pop()
			stack.data[len(stack.data)-1] = div
		default:
			fmt.Println("Can´t devide by ZERO!")
		}
		//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Pow() {
	switch {
	case len(stack.data) >= 2:
		//Potentiate second last number to the power of the last number of the stack
		result := math.Pow(stack.data[len(stack.data)-2], stack.data[len(stack.data)-1])

		//Updating stack
		stack.Pop()
		stack.data[len(stack.data)-1] = result

	//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Sqrt() {
	switch {
	case len(stack.data) >= 1:
		//Potentiate with 0.5
		sqrt := math.Pow(stack.data[len(stack.data)-1], 0.5)

		//Updating stack
		stack.data[len(stack.data)-1] = sqrt

	//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Abs() {
	switch {
	case len(stack.data) >= 1:

		//Turning positiv :)
		if stack.data[len(stack.data)-1] < 0 {
			//Updating stack
			stack.data[len(stack.data)-1] = -1 * stack.data[len(stack.data)-1]
		}

	//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
	}
}

func (stack *Stack) Fact() bool {
	result := 1

	switch {
	case len(stack.data) >= 1:
		switch {
		case stack.data[len(stack.data)-1] < 0:
			fmt.Println("Fakultät nur für positive Zahlen möglich!")
			return true

		//Checkt auf Nachkommastellen
		case math.Floor(stack.data[len(stack.data)-1]) != stack.data[len(stack.data)-1]:
			fmt.Println("Fakultät kann nur auf Ganzzahlen angwandt werden!")
			return true

		default:
			if stack.data[len(stack.data)-1] == 0 || stack.data[len(stack.data)-1] == 1 {
				stack.data[len(stack.data)-1] = 1
				return false
			} else {
				for i := 2; i <= int(stack.data[len(stack.data)-1]); i++ {
					result *= i
				}
			}
			stack.data[len(stack.data)-1] = float64(result)
			return false
		}
	//Stack to small
	default:
		fmt.Println("Not enough numbers in Stack")
		return false //don´t have to Pop from History --> already covered
	}
	//ToDo Update History bei Fehlschlag
}
