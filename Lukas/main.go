package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack []float64

func (s *Stack) Push(input float64) {
	*s = append(*s, input)
}

func (s *Stack) Pop() float64 {
	if(len(*s) == 0) {
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

type LaTeXOutput struct {
	Expression string
}

// Function to format the input string into LaTeX
func (l LaTeXOutput) formatToLatex(input string) string {
	var latexBuilder strings.Builder
	tokens := strings.Split(input, " ")

	stack := []float64{}

	for _, token := range tokens {
		switch token {
		case "+":
			b := stack[len(stack)-1] // The second operand
			stack = stack[:len(stack)-1] // Pop the stack
			a := stack[len(stack)-1] // The first operand
			stack = stack[:len(stack)-1] // Pop the stack
			latexBuilder.WriteString(fmt.Sprintf("{%s} + {%s}", l.formatToLatexNumber(a), l.formatToLatexNumber(b)))
			stack = append(stack, 0) // Push a placeholder result

		case "-":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			latexBuilder.WriteString(fmt.Sprintf("{%s} - {%s}", l.formatToLatexNumber(a), l.formatToLatexNumber(b)))
			stack = append(stack, 0) // Push a placeholder result

		case "*":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			latexBuilder.WriteString(fmt.Sprintf("{%s} \\cdot {%s}", l.formatToLatexNumber(a), l.formatToLatexNumber(b)))
			stack = append(stack, 0) // Push a placeholder result

		case "/":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			latexBuilder.WriteString(fmt.Sprintf("\\frac{%s}{%s}", l.formatToLatexNumber(a), l.formatToLatexNumber(b)))
			stack = append(stack, 0) // Push a placeholder result

		case "^":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			latexBuilder.WriteString(fmt.Sprintf("{%s}^{%s}", l.formatToLatexNumber(a), l.formatToLatexNumber(b)))
			stack = append(stack, 0) // Push a placeholder result

		case "sqrt":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			latexBuilder.WriteString(fmt.Sprintf("\\sqrt{%s}", l.formatToLatexNumber(a)))
			stack = append(stack, 0) // Push a placeholder result 

		default:
			// Parse and push number onto the stack
			var num float64
			fmt.Sscanf(token, "%g", &num)
			stack = append(stack, num)
			latexBuilder.WriteString(l.formatToLatexNumber(num))
		}
		latexBuilder.WriteString(" ") // Add space for clarity
	}

	finalLatex := fmt.Sprintf("\\[%s\\]", latexBuilder.String())
	return finalLatex
}

// Formatting for numbers to LaTeX
func (l LaTeXOutput) formatToLatexNumber(val float64) string {
	return fmt.Sprintf("%g", val)
}

func main() {

	stack := Stack{}
	latex := LaTeXOutput{}

	printWelcomeMessage()

	for {
		stack.Print()

		input := getInput()

		checkInput(input, &stack, &latex)
	}
}

func printWelcomeMessage() {
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


func getInput() string {
		reader := bufio.NewReader(os.Stdin)
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Wrong input: ", err)
			return ""
		}
		input = strings.TrimSpace(input)
		return input
}

func checkInput(input string, stack *Stack, latex *LaTeXOutput) {

	// commands
	if input == "exit" {
		fmt.Println("\nQuitting Application. See you soon!")
		os.Exit(0)
	} else if input == "list" {
		fmt.Println("printing input")
		stack.Print()
		return
	} else if input == "latex" {
		fmt.Println(latex.formatToLatex(latex.Expression))
		return
	} else if input == "help" {
		printWelcomeMessage()
		return
	}

	// operators
	switch input {
	case "+": 
		result := stack.Pop() + stack.Pop()
		stack.Push(result)
	case "-": 
		secondOp := stack.Pop()
		firstOp := stack.Pop()
		result := firstOp - secondOp
		stack.Push(result)
	case "*": 
		result := stack.Pop() * stack.Pop()
		stack.Push(result)
	case "/": 
		secondOp := stack.Pop()
		firstOp := stack.Pop()
		result := firstOp / secondOp
		stack.Push(result)
	case "abs":
		result := math.Abs(stack.Pop())
		stack.Push(result)
	case "sqrt":
		result := math.Sqrt(stack.Pop())
		stack.Push(result)
	case "log":
		result := math.Log(stack.Pop())
		stack.Push(result)
	case "^":
		exponent := stack.Pop()
		base := stack.Pop()
		result := math.Pow(base,exponent)
		stack.Push(result)
	case "++":
		n := len(*stack)
		result := 0.0
		for i := 0; i < n; i++ {
			result += stack.Pop()
		}
		stack.Push(result)
	case "**":
		n := len(*stack)
		result := 1.0
		for i := 0; i < n; i++ {
			result *= stack.Pop()
		}
		if n > 0 {
			stack.Push(result)
		}
	}

	// numbers
	number, err := strconv.ParseFloat(input, 64)
	if  err != nil {
		
	} else {
		stack.Push(number)
	}

	// add to history
	latex.Expression = latex.Expression + " " +input
	//fmt.Println("current latexExp ", latex.Expression)
}