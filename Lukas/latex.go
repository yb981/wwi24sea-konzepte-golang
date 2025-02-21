package main

import (
	"fmt"
	"strings"
)

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
			b := stack[len(stack)-1]     // The second operand
			stack = stack[:len(stack)-1] // Pop the stack
			a := stack[len(stack)-1]     // The first operand
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