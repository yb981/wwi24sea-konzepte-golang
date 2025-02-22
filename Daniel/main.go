package main

import (
	"fmt"
)

var allOperands = [11]string{"+", "-", "*", "/", "^", "sqrt", "log", "!", "abs", "++", "**"} //11 Operanden

func main() {
	fmt.Println(allOperands, "\n")
	var stack Stack
	var history History

Loop:
	for {
		input, operandAndHistory, valid, parseErr := userInput()

		if valid {
			switch {
			case !(isSpecialOperand(operandAndHistory)):
				if parseErr == nil || len(stack.data) >= 2 {
					history.Push(operandAndHistory)
				}

			//sqrt, log, !, abs
			case len(stack.data) >= 1:
				history.Push(operandAndHistory)
			}
			switch {
			case operandAndHistory == "exit":
				break Loop

			case parseErr == nil:
				stack.Push(input)

			case contains(allOperands, operandAndHistory):
				switch operandAndHistory {
				case "+":
					stack.Add()

				case "-":
					stack.Sub()

				case "*":
					stack.Mult()

				case "/":
					stack.Div()

				case "^":
					stack.Pow()

				case "sqrt":
					stack.Sqrt()

				case "abs":
					stack.Abs()

				case "!":
					noFact := stack.Fact()
					if noFact {
						history.Pop()
					}
				}
			}
		}
		fmt.Println("Stack:  ", stack.data)
		fmt.Println("History:", history.data)
		fmt.Println()
	}
}
