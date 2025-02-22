package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []float64
}

type History struct {
	data []string
}

func (stack *Stack) Push(input float64) {
	stack.data = append(stack.data, input)
}

func (stack *Stack) Pop() {
	if len(stack.data) == 0 {
		err := errors.New("Cannot pop from an empty stack")
		fmt.Println(err.Error())
	} else {
		stack.data = (stack.data)[:len(stack.data)-1]
	}
}

func (stack *History) Push(input string) {
	stack.data = append(stack.data, input)
}

func (stack *History) Pop() {
	if len(stack.data) == 0 {
		err := errors.New("Cannot pop from an empty stack")
		fmt.Println(err.Error())
	} else {
		stack.data = (stack.data)[:len(stack.data)-1]
	}
}

func userInput() (float64, string, bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var valid bool

	fmt.Print("Eingabe: ")
	operandAndHistory, inputErr := reader.ReadString('\n')
	operandAndHistory = strings.TrimSpace(operandAndHistory)
	number, parseErr := strconv.ParseFloat(operandAndHistory, 64)
	fmt.Printf("Zahl: _%v_ <%T>\n", number, number)
	fmt.Printf("Input: _%v_ <%T>\n\n", operandAndHistory, operandAndHistory)

	//Validation
	switch {
	case inputErr != nil:
		valid = false
		fmt.Println("Fehler bei der Eingabe:", inputErr)
		return number, "", valid, parseErr

	case parseErr != nil && !(contains(allOperands, operandAndHistory)):
		valid = false
		fmt.Println("Fehlerhafte Eingabe: Nur Zahlen und Operanden erlaubt!")
		return number, "", valid, parseErr

	case operandAndHistory == "exit":
		return number, operandAndHistory, valid, parseErr

	//Valid Input --> textInput also carries operands
	default:
		valid = true
		return number, operandAndHistory, valid, parseErr
	}
}

func contains(list [11]string, input string) bool {
	for _, operand := range list {
		if input == operand {
			return true
		}
	}
	return false
}

func isSpecialOperand(input string) bool {
	specialOperands := [4]string{"sqrt", "log", "!", "abs"}
	for _, operand := range specialOperands {
		if input == operand {
			return true
		}
	}
	return false
}
