// calculator.go
// RPN Calculator
//
// Author: Lukas Gröning
// Date: 22.02.2025
//
// RPN Calculator with basic operations.

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type calculator struct {
	numberStack Stack[float64]
	history     Stack[string]
	latex       LaTeXOutput
}

func (c *calculator) checkInput(input string) {
	// Commands
	switch input {
	case "exit":
		fmt.Println("\nQuitting Application. See you soon!")
		os.Exit(0)
	case "list":
		fmt.Println("printing input")
		c.numberStack.Print()
		return
	case "latex":
		c.latex.Expression = c.history.Top()
		fmt.Println(c.history.Top())
		fmt.Println(c.latex.Expression)
		fmt.Println(c.latex.formatToLatex(c.latex.Expression))
		return
	case "help":
		c.printWelcomeMessage()
		return
	}

	// Operations handling
	switch input {
	case "+", "-", "*", "/", "^":
		c.performBinaryOperation(input)
	case "abs", "sqrt", "log", "!":
		c.performUnaryOperation(input)
	case "++":
		c.performSumOperation()
	case "**":
		c.performProductOperation()
	default:
		c.handleNumberInput(input)
	}
}

func (c *calculator) performBinaryOperation(op string) {
	term2 := c.history.Pop()
	term1 := c.history.Pop()
	var result float64

	switch op {
	case "+":
		result = c.numberStack.Pop() + c.numberStack.Pop()
	case "-":
		secondOp := c.numberStack.Pop()
		firstOp := c.numberStack.Pop()
		result = firstOp - secondOp
	case "*":
		result = c.numberStack.Pop() * c.numberStack.Pop()
	case "/":
		secondOp := c.numberStack.Pop()
		firstOp := c.numberStack.Pop()
		result = firstOp / secondOp
	case "^":
		exponent := c.numberStack.Pop()
		base := c.numberStack.Pop()
		result = math.Pow(base, exponent)
	}

	termNew := fmt.Sprintf("(%s %s %s)", term1, op, term2)
	c.history.Push(termNew)
	fmt.Printf("current calculation: %s = %v\n", termNew, result)
	c.numberStack.Push(result)
}

func (c *calculator) performUnaryOperation(op string) {
	term1 := c.history.Pop()
	var result float64

	switch op {
	case "abs":
		result = math.Abs(c.numberStack.Pop())
		c.history.Push(fmt.Sprintf("abs(%s)", term1))
	case "sqrt":
		result = math.Sqrt(c.numberStack.Pop())
		c.history.Push(fmt.Sprintf("sqrt(%s)", term1))
	case "log":
		result = math.Log(c.numberStack.Pop())
		c.history.Push(fmt.Sprintf("log(%s)", term1))
	case "!":
		current := c.numberStack.Pop()
		if current < 0 || current != math.Floor(current) { // Check for non-negative integer
			fmt.Println("Error: Factorial is not defined for negative numbers or non-integers.")
			c.numberStack.Push(current) // Push number back on stack
			return
		}
		result = c.factorial(current)
		c.history.Push(fmt.Sprintf("%d!", int(current))) // Store the factorial expression
	}
	fmt.Printf("current calculation: %s = %v\n", c.history.Top(), result)
	c.numberStack.Push(result)
}

func (c *calculator) performSumOperation() {
	n := len(c.numberStack)
	result := 0.0
	
	for i := 0; i < n; i++ {
		result += c.numberStack.Pop()
	}
	c.numberStack.Push(result)
	fmt.Printf("current calculation: sum = %v\n", result)
}

func (c *calculator) performProductOperation() {
	n := len(c.numberStack)
	result := 1.0
	
	for i := 0; i < n; i++ {
		result *= c.numberStack.Pop()
	}
	if n > 0 {
		c.numberStack.Push(result)
		fmt.Printf("current calculation: product = %v\n", result)
	}
}

func (c *calculator) handleNumberInput(input string) {
	number, err := strconv.ParseFloat(input, 64)
	if err == nil {
		c.numberStack.Push(number)
		c.history.Push(input)
	} else {
		// Fehlerbehandlung
	}
}

// Helper function to compute factorial of a non-negative integer
func (c calculator) factorial(n float64) float64 {
	if n == 0 {
		return 1 // 0! is 1
	}
	var result float64 = 1
	for i := 1; i <= int(n); i++ {
		result *= float64(i)
	}
	return result
}

func (c calculator) printWelcomeMessage() {
	fmt.Println("=========================================")
	fmt.Println("        Welcome to the RPN Calculator   ")
	fmt.Println("=========================================")
	fmt.Println("Functionality:")
	fmt.Println(" +  Addition")
	fmt.Println(" -  Subtraction")
	fmt.Println(" *  Multiplication")
	fmt.Println(" /  Division")
	fmt.Println(" ^  Exponentiation")
	fmt.Println(" sqrt  Square Root")
	fmt.Println(" log  Logarithm (Base 10)")
	fmt.Println(" !  Factorial")
	fmt.Println(" abs  Absolute Value")
	fmt.Println(" ++  Sum All Numbers on the Stack")
	fmt.Println(" **  Multiply All Numbers on the Stack")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println(" - Type 'help' for assistance")
	fmt.Println(" - Type 'latex' for LaTeX formatted output")
	fmt.Println(" - Type 'exit' to close the application")
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("Please enter your command:")
}