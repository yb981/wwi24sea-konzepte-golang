// latex.go
// Functionality to parse strings into latex
//
// Author: Lukas Gröning
// Date: 22.02.2025

package main

import (
	"strings"
)

type LaTeXConverter struct {}

func (l LaTeXConverter) convertToLatex(input string) string {
	// TODO 
	// implementation not correct yet

	// Replace "/" with "\frac{}{}" for fractions
	input = strings.ReplaceAll(input, "/", "\\frac{")
	input = strings.ReplaceAll(input, ")", "}{") + "}"

	// Replace "*" with "\cdot" for multiplication
	input = strings.ReplaceAll(input, "*", " \\cdot ")

	// Replace "sqrt" with "\sqrt{}" for square root
	input = strings.ReplaceAll(input, "sqrt", "\\sqrt{")

	// Add the surrounding LaTeX display math brackets
	return "\\[" + input + "\\]"
}