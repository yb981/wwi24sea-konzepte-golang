// calculator_test.go
// test for RPN Calculator
//
// Author: Lukas Gröning, Ajun Anpalakan
// Date: 22.02.2025

package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/yb981/wwi24sea-konzepte-golang/datastructures"
)

func setupCalculator() *calculator {
	return &calculator{
		numberStack: datastructures.Stack[float64]{},
		history:     datastructures.Stack[string]{},
		latex:       datastructures.Stack[string]{},
	}
}

func assertTopFloat(t *testing.T, stack datastructures.Stack[float64], expected float64, msg string) {
	t.Helper()
	val, err := stack.Peek()
	if err != nil || val != expected {
		t.Errorf("%s: expected %v, got %v", msg, expected, val)
	}
	_, _ = stack.Pop() // remove top for next test if needed
}

func assertTopString(t *testing.T, stack datastructures.Stack[string], expected string, msg string) {
	t.Helper()
	val, err := stack.Peek()
	if err != nil || val != expected {
		t.Errorf("%s: expected %s, got %s", msg, expected, val)
	}
}

func TestBinaryOperations(t *testing.T) {
	calc := setupCalculator()

	calc.numberStack.Push(4)
	calc.numberStack.Push(2)
	calc.performBinaryOperation("+")
	assertTopFloat(t, calc.numberStack, 6, "4 + 2 sollte 6 sein")

	calc.numberStack.Push(10)
	calc.numberStack.Push(2)
	calc.performBinaryOperation("/")
	assertTopFloat(t, calc.numberStack, 5, "10 / 2 sollte 5 sein")

	calc.numberStack.Push(3)
	calc.numberStack.Push(3)
	calc.performBinaryOperation("^")
	assertTopFloat(t, calc.numberStack, 27, "3 ^ 3 sollte 27 sein")
}

func TestFactorial(t *testing.T) {
	calc := setupCalculator()
	if calc.factorial(0) != 1 || calc.factorial(5) != 120 || calc.factorial(1) != 1 {
		t.Error("Factorial function is incorrect")
	}
}

func TestSumOperation(t *testing.T) {
	calc := setupCalculator()
	calc.numberStack.Push(1)
	calc.numberStack.Push(2)
	calc.numberStack.Push(3)
	calc.performMultiOperation("++")
	assertTopFloat(t, calc.numberStack, 6, "1 + 2 + 3 sollte 6 sein")
}

func TestProductOperation(t *testing.T) {
	calc := setupCalculator()
	calc.numberStack.Push(2)
	calc.numberStack.Push(3)
	calc.numberStack.Push(4)
	calc.performMultiOperation("**")
	assertTopFloat(t, calc.numberStack, 24, "2 * 3 * 4 sollte 24 sein")
}

func TestHandleWrongInput(t *testing.T) {
	calc := setupCalculator()
	calc.handleNumberInput("abc")
	if calc.numberStack.Size() != 0 {
		t.Error("Ungültige Eingabe sollte nicht auf den Stack gelangen")
	}
}

func TestIntegration(t *testing.T) {
	calc := setupCalculator()
	calc.handleNumberInput("2")
	calc.handleNumberInput("3")
	calc.checkInput("+")
	assertTopFloat(t, calc.numberStack, 5, "2 + 3 sollte 5 sein")

	calc.handleNumberInput("4")
	calc.checkInput("sqrt")
	assertTopFloat(t, calc.numberStack, 2, "sqrt(4) sollte 2 sein")
}

func TestInfixNotationOutput(t *testing.T) {
	calc := setupCalculator()
	calc.handleNumberInput("3")
	calc.handleNumberInput("5")
	calc.performBinaryOperation("+")
	assertTopString(t, calc.history, "(3 + 5)", "Erwartete Infix-Notation 1")

	calc.handleNumberInput("2")
	calc.performBinaryOperation("*")
	assertTopString(t, calc.history, "((3 + 5) * 2)", "Erwartete Infix-Notation 2")

	calc.handleNumberInput("2")
	calc.performBinaryOperation("^")
	assertTopString(t, calc.history, "(((3 + 5) * 2) ^ 2)", "Erwartete Infix-Notation 3")
}

func TestComplexCalculationWithFactorial(t *testing.T) {
	calc := setupCalculator()

	calc.handleNumberInput("3")
	calc.performUnaryOperation("!")
	calc.handleNumberInput("2")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("+")
	res, err := calc.numberStack.Pop()
	if err != nil || res != 8 {
		t.Errorf("3! + 2! sollte 8 sein, aber war %v (err: %v)", res, err)
	}

	calc.handleNumberInput("5")
	calc.performUnaryOperation("!")
	calc.handleNumberInput("4")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("-")
	calc.handleNumberInput("2")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("*")
	res, err = calc.numberStack.Pop()
	if err != nil || res != 192 {
		t.Errorf("(5! - 4!) * 2! sollte 192 sein, aber war %v (err: %v)", res, err)
	}

	calc.handleNumberInput("5")
	calc.performUnaryOperation("!")
	calc.handleNumberInput("4")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("-")
	calc.handleNumberInput("3")
	calc.performUnaryOperation("!")
	calc.handleNumberInput("2")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("+")
	calc.performBinaryOperation("*")
	res, err = calc.numberStack.Pop()
	if err != nil || res != 768 {
		t.Errorf("(5! - 4!) * (3! + 2!) sollte 768 sein, aber war %v (err: %v)", res, err)
	}
}

func captureOutput(f func()) string {
	// Speichere das aktuelle stdout
	old := os.Stdout

	// Pipe erstellen
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Funktion ausführen, deren Ausgabe wir abfangen wollen
	f()

	// Pipe schließen und os.Stdout wiederherstellen
	w.Close()
	os.Stdout = old

	// Ausgabe lesen
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	return buf.String()
}

func TestCheckInput_Commands(t *testing.T) {
	c := setupCalculator()
	// Test help
	output := captureOutput(func() {
		c.checkInput("help")
	})
	if !strings.Contains(output, "Welcome to the RPN Calculator") {
		t.Error("help Befehl gab falsche Ausgabe")
	}

	// Test latex
	c.latex.Push("x^2")
	output = captureOutput(func() {
		c.checkInput("latex")
	})
	if !strings.Contains(output, "\\[x^2\\]") {
		t.Error("latex Befehl gab falsche Ausgabe")
	}
}

func TestPerformMultiOperation_TooFewItems(t *testing.T) {
	c := setupCalculator()
	output := captureOutput(func() {
		c.checkInput("++")
	})
	if !strings.Contains(output, "at least 2 numbers") {
		t.Error("multi-op mit <2 items soll einen Fehler ergeben")
	}
}

func TestHandleNumberInput(t *testing.T) {
	c := setupCalculator()
	c.handleNumberInput("42.5")
	if c.numberStack.Size() != 1 {
		t.Error("handleNumberInput hat nichts auf den Stack gelegt")
	}
	top, err := c.numberStack.Peek()
	if err != nil || top != 42.5 {
		t.Errorf("Top des Stacks sollte 42.5 sein, war aber %v (err: %v)", top, err)
	}
}

func TestRestoreState(t *testing.T) {
	c := setupCalculator()
	c.restoreState("term", "latex", 5.0)

	num, err1 := c.numberStack.Peek()
	hist, err2 := c.history.Peek()
	latex, err3 := c.latex.Peek()

	if err1 != nil || num != 5.0 {
		t.Errorf("Zahl Stack: Erwartet 5.0, aber war %v (err: %v)", num, err1)
	}
	if err2 != nil || hist != "term" {
		t.Errorf("History Stack: Erwartet 'term', aber war %v (err: %v)", hist, err2)
	}
	if err3 != nil || latex != "latex" {
		t.Errorf("Latex Stack: Erwartet 'latex', aber war %v (err: %v)", latex, err3)
	}
}
func TestUnaryOperations_Abs(t *testing.T) {
	c := setupCalculator()
	c.numberStack.Push(-5)
	c.history.Push("-5")
	c.latex.Push("-5")
	output := captureOutput(func() {
		c.performUnaryOperation("abs")
	})
	if !strings.Contains(output, "abs(-5) = 5") {
		t.Error("abs funktionert nicht richtig")
	}
}

func TestUnaryOperations_Sqrt_Negative(t *testing.T) {
	c := setupCalculator()
	c.numberStack.Push(-4)
	c.history.Push("-4")
	c.latex.Push("-4")
	output := captureOutput(func() {
		c.performUnaryOperation("sqrt")
	})
	if !strings.Contains(output, "not defined for negative numbers") {
		t.Error("sqrt error handling für negative Nummer fehlgeschlagen")
	}
}

func TestUnaryOperations_Log_Normal(t *testing.T) {
	c := setupCalculator()
	c.numberStack.Push(10)
	c.history.Push("10")
	c.latex.Push("10")
	output := captureOutput(func() {
		c.performUnaryOperation("log")
	})
	if !strings.Contains(output, "log(10)") {
		t.Error("log Berechnung falsch")
	}
}

func TestUnaryOperations_Log_Error(t *testing.T) {
	c := setupCalculator()
	c.numberStack.Push(0)
	c.history.Push("0")
	c.latex.Push("0")
	output := captureOutput(func() {
		c.performUnaryOperation("log")
	})
	if !strings.Contains(output, "not defined for zero or negative numbers") {
		t.Error("log error handling falsch")
	}
}

func TestUnaryOperations_Factorial_Error(t *testing.T) {
	c := setupCalculator()
	c.numberStack.Push(-2)
	c.history.Push("-2")
	c.latex.Push("-2")
	output := captureOutput(func() {
		c.performUnaryOperation("!")
	})
	if !strings.Contains(output, "Factorial is not defined for negative numbers") {
		t.Error("factorial error handling falsch")
	}
}

func TestBinaryOperation_NoNumbers(t *testing.T) {
	c := setupCalculator()
	output := captureOutput(func() {
		c.performBinaryOperation("+")
	})
	if !strings.Contains(output, "at least 2 numbers") {
		t.Error("binary operation sollte ohne Nummern nicht funktionieren")
	}
}

func TestUnaryOperation_NoNumbers(t *testing.T) {
	c := setupCalculator()
	output := captureOutput(func() {
		c.performUnaryOperation("abs")
	})
	if !strings.Contains(output, "at least 1 number") {
		t.Error("unary operation sollte ohne Nummern nicht funktionieren")
	}
}
