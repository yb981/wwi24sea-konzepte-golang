package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RPN Taschenrechner. Geben Sie eine RPN-Ausdruck ein:")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Fehler beim Lesen der Eingabe:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Beenden...")
			break
		}

		_, latexExpr, err := evaluateRPN(input)
		if err != nil {
			fmt.Println("Fehler:", err)
		} else {
			fmt.Println("RPN:", input)
			fmt.Println("Latex:", latexExpr)
			fmt.Println("Latex (rendered):")
			fmt.Println("\\[", latexExpr, "\\]")
		}
	}
}

func evaluateRPN(expression string) (float64, string, error) {
	tokens := strings.Split(expression, " ")
	var stack []float64
	var latexStack []string

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, "", fmt.Errorf("nicht genügend Operanden für %s", token)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			bStr := latexStack[len(latexStack)-1]
			aStr := latexStack[len(latexStack)-2]
			latexStack = latexStack[:len(latexStack)-2]

			var result float64
			var latexExpr string
			switch token {
			case "+":
				result = a + b
				latexExpr = fmt.Sprintf("%s + %s", aStr, bStr)
			case "-":
				result = a - b
				latexExpr = fmt.Sprintf("%s - %s", aStr, bStr)
			case "*":
				result = a * b
				latexExpr = fmt.Sprintf("%s \\cdot %s", aStr, bStr)
			case "/":
				if b == 0 {
					return 0, "", fmt.Errorf("Division durch Null")
				}
				result = a / b
				latexExpr = fmt.Sprintf("\\frac{%s}{%s}", aStr, bStr)
			case "^":
				result = math.Pow(a, b)
				latexExpr = fmt.Sprintf("%s^{%s}", aStr, bStr)
			}
			stack = append(stack, result)
			latexStack = append(latexStack, latexExpr)
		case "sqrt":
			if len(stack) < 1 {
				return 0, "", fmt.Errorf("nicht genügend Operanden für sqrt")
			}
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			aStr := latexStack[len(latexStack)-1]
			latexStack = latexStack[:len(latexStack)-1]
			stack = append(stack, math.Sqrt(a))
			latexStack = append(latexStack, fmt.Sprintf("\\sqrt{%s}", aStr))
		case "log":
			if len(stack) < 1 {
				return 0, "", fmt.Errorf("nicht genügend Operanden für log")
			}
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, math.Log10(a))
			latexStack = append(latexStack, fmt.Sprintf("\\log_{10}(%s)", latexStack[len(latexStack)-1]))
		case "abs":
			if len(stack) < 1 {
				return 0, "", fmt.Errorf("nicht genügend Operanden für abs")
			}
			stack[len(stack)-1] = math.Abs(stack[len(stack)-1])
		case "++":
			sum := 0.0
			for _, num := range stack {
				sum += num
			}
			stack = []float64{sum}
		case "**":
			prod := 1.0
			for _, num := range stack {
				prod *= num
			}
			stack = []float64{prod}
		case "!":
			if len(stack) < 1 {
				return 0, "", fmt.Errorf("nicht genügend Operanden für !")
			}
			n := int(stack[len(stack)-1])
			stack[len(stack)-1] = float64(factorial(n))
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, "", fmt.Errorf("ungültige Zahl: %s", token)
			}
			stack = append(stack, num)
			latexStack = append(latexStack, fmt.Sprintf("%.0f", num))
		}
	}

	if len(stack) != 1 {
		return 0, "", fmt.Errorf("ungültiger Ausdruck")
	}

	return stack[0], latexStack[0], nil
}
