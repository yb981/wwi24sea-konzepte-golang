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
	latex     	Stack[string]
	//latex       LaTeXConverter
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
		fmt.Println("\\["+c.latex.Top()+"\\]")
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
	case "++", "**":
		c.performMultiOperation(input)
	default:
		c.handleNumberInput(input)
	}
}

func (c *calculator) performBinaryOperation(op string) {
	if len(c.numberStack) < 2 {
		fmt.Println("Error: Need at least 2 numbers on the stack.")
		return
	}

	latex2 := c.latex.Pop()
	latex1 := c.latex.Pop()

	term2 := c.history.Pop()
	term1 := c.history.Pop()

	var result float64
	secondOp := c.numberStack.Pop()
	firstOp := c.numberStack.Pop()

	switch op {
	case "+":
		result = firstOp + secondOp
		c.latex.Push(fmt.Sprintf("({%s} + {%s})", latex1, latex2))
	case "-":
		result = firstOp - secondOp
		c.latex.Push(fmt.Sprintf("({%s} - {%s})", latex1, latex2))
	case "*":
		result = firstOp * secondOp
		c.latex.Push(fmt.Sprintf("({%s} \\cdot {%s})", latex1, latex2))
	case "/":
		result = firstOp / secondOp
		c.latex.Push(fmt.Sprintf("\\frac{%s}{%s}", latex1, latex2))
	case "^":
		result = math.Pow(firstOp, secondOp)
		c.latex.Push(fmt.Sprintf("{%s}^{%s}", latex1, latex2))
	}

	termNew := fmt.Sprintf("(%s %s %s)", term1, op, term2)
	c.history.Push(termNew)
	fmt.Printf("current calculation: %s = %v\n", termNew, result)
	c.numberStack.Push(result)
}

func (c *calculator) performUnaryOperation(op string) {
	if len(c.numberStack) < 1 {
		fmt.Println("Error: Need at least 1 number on the stack.")
		return
	}

	latex1 := c.latex.Pop()
	term1 := c.history.Pop()

	var result float64

	switch op {
	case "abs":
		result = math.Abs(c.numberStack.Pop())
		c.history.Push(fmt.Sprintf("abs(%s)", term1))
		c.latex.Push(fmt.Sprintf("\\rvert{%s}", latex1))
	case "sqrt":
		result = math.Sqrt(c.numberStack.Pop())
		c.history.Push(fmt.Sprintf("sqrt(%s)", term1))
		c.latex.Push(fmt.Sprintf("\\sqrt{%s}", latex1))
	case "log":
		result = math.Log(c.numberStack.Pop())
		c.history.Push(fmt.Sprintf("log(%s)", term1))
		c.latex.Push(fmt.Sprintf("log{%s}", latex1))
	case "!":
		current := c.numberStack.Pop()
		if current < 0 || current != math.Floor(current) { // Check for non-negative integer
			fmt.Println("Error: Factorial is not defined for negative numbers or non-integers.")
			c.numberStack.Push(current) // Push number back on stack
			return
		}
		result = c.factorial(current)
		c.history.Push(fmt.Sprintf("%d!", int(current)))
		c.latex.Push(fmt.Sprintf("%d!", int(current)))
	}
	fmt.Printf("current calculation: %s = %v\n", c.history.Top(), result)
	c.numberStack.Push(result)
}

func (c *calculator) performMultiOperation(op string) {
	if len(c.numberStack) < 2 {
		fmt.Println("Error: Need at least 2 numbers on the stack.")
		return
	}

	n := len(c.numberStack)

	tempSlice := make([]float64, n)
	tempSliceHistory := make([]string, n)
	tempSliceLatex := make([]string, n)

	for i := n-1; i >= 0; i-- {
		tempSlice[i] = c.numberStack.Pop()
		tempSliceHistory[i] = c.history.Pop()
		tempSliceLatex[i] = c.latex.Pop()
	}
	
	historyOutput := "("
	latexOutput := "("

	var result float64
	if op == "++" {
		result = 0.0
	} else {
		result = 1.0
	}
	

	for i := 0; i < n; i++ {
		current := tempSlice[i]
		if op == "++" {
			result += current
		} else {
			result *= current
		}
		
		historyOutput += tempSliceHistory[i]
		latexOutput += tempSliceLatex[i]

		if i != n-1 {
			if op == "++" {
				historyOutput += " + "
				latexOutput += " + "
			} else {
				historyOutput += " * "
				latexOutput += " \\cdot "	
			}
		}
	}

	historyOutput += ")"
	latexOutput += ")"

	c.numberStack.Push(result)
	c.history.Push(historyOutput)
	c.latex.Push(latexOutput)
	fmt.Printf("current calculation: %s = %v\n", historyOutput, result)
}

func (c *calculator) handleNumberInput(input string) {
	number, err := strconv.ParseFloat(input, 64)
	if err == nil {
		c.numberStack.Push(number)
		c.history.Push(input)
		c.latex.Push(input)
	} else {
		fmt.Println("Error: Wrong Input")
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