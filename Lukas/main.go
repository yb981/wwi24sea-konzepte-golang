package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	numberStack := Stack[float64]{}
	latex := LaTeXOutput{}
	history := Stack[string]{}

	printWelcomeMessage()

	for {
		numberStack.Print()

		input := getInput()

		checkInput(input, &numberStack, &latex, &history)
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

func getInputByReader() string {
		reader := bufio.NewReader(os.Stdin)
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Wrong input: ", err)
			return ""
		}
		input = strings.TrimSpace(input)
		return input
}

func getInput() string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Eingabefehler")
	}
	return input
}

func checkInput(input string, numberStack *Stack[float64], latex *LaTeXOutput, history *Stack[string]) {

	// commands
	if input == "exit" {
		fmt.Println("\nQuitting Application. See you soon!")
		os.Exit(0)
	} else if input == "list" {
		fmt.Println("printing input")
		numberStack.Print()
		return
	} else if input == "latex" {
		latex.Expression = history.Top()
		fmt.Println(history.Top())
		fmt.Println(latex.Expression)
		fmt.Println(latex.formatToLatex(latex.Expression))
		return
	} else if input == "help" {
		printWelcomeMessage()
		return
	}

	// operators
	switch input {
	case "+": 
		result := numberStack.Pop() + numberStack.Pop()
		term2 := history.Pop()
		term1 := history.Pop()
		termNew := "(" + term1 + "+" +term2+")"
		history.Push(termNew)
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "-": 
		secondOp := numberStack.Pop()
		firstOp := numberStack.Pop()
		term2 := history.Pop()
		term1 := history.Pop()
		termNew := "(" + term1 + "-" +term2+")"
		history.Push(termNew)
		result := firstOp - secondOp
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "*": 
		result := numberStack.Pop() * numberStack.Pop()
		term2 := history.Pop()
		term1 := history.Pop()
		termNew := "(" + term1 + "*" +term2+")"
		history.Push(termNew)
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "/": 
		secondOp := numberStack.Pop()
		firstOp := numberStack.Pop()
		term2 := history.Pop()
		term1 := history.Pop()
		termNew := "(" + term1 + "/" +term2+")"
		history.Push(termNew)
		result := firstOp / secondOp
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "abs":
		result := math.Abs(numberStack.Pop())
		term1 := history.Pop()
		termNew := "abs(" + term1 +")"
		history.Push(termNew)
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "sqrt":
		result := math.Sqrt(numberStack.Pop())
		term1 := history.Pop()
		termNew := "sqrt(" + term1 +")"
		history.Push(termNew)
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "log":
		result := math.Log(numberStack.Pop())
		term1 := history.Pop()
		termNew := "log(" + term1 +")"
		history.Push(termNew)
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "^":
		exponent := numberStack.Pop()
		base := numberStack.Pop()
		result := math.Pow(base,exponent)
		term2 := history.Pop()
		term1 := history.Pop()
		termNew := "(" + term1 + "^" +term2+")"
		history.Push(termNew)
		fmt.Println("current calculation: " + termNew + " = ", result)
		numberStack.Push(result)
	case "++":
		n := len(*numberStack)
		result := 0.0
		// TODO history
		for i := 0; i < n; i++ {
			result += numberStack.Pop()
		}
		numberStack.Push(result)
	case "**":
		n := len(*numberStack)
		result := 1.0
		// TODO history
		for i := 0; i < n; i++ {
			result *= numberStack.Pop()
		}
		if n > 0 {
			numberStack.Push(result)
		}
	}

	// numbers
	number, err := strconv.ParseFloat(input, 64)
	if  err != nil {
		// TODO
	} else {
		numberStack.Push(number)
		history.Push(input)
	}
}