package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestBinaryOperations(t *testing.T) {
	calc := calculator{}

	calc.numberStack.Push(4)
	calc.numberStack.Push(2)
	calc.performBinaryOperation("+")
	if calc.numberStack.Pop() != 6 {
		t.Error("4 + 2 sollte 6 sein")
	}

	calc.numberStack.Push(10)
	calc.numberStack.Push(2)
	calc.performBinaryOperation("/")
	if calc.numberStack.Pop() != 5 {
		t.Error("10 / 2 sollte 5 sein")
	}

	calc.numberStack.Push(3)
	calc.numberStack.Push(3)
	calc.performBinaryOperation("^")
	if calc.numberStack.Pop() != 27 {
		t.Error("3 ^ 3 sollte 27 sein")
	}
}

func TestFactorial(t *testing.T) {
	calc := calculator{}

	if calc.factorial(0) != 1 {
		t.Error("0! sollte 1 sein")
	}

	if calc.factorial(5) != 120 {
		t.Error("5! sollte 120 sein")
	}

	if calc.factorial(1) != 1 {
		t.Error("1! sollte 1 sein")
	}
}

func TestSumOperation(t *testing.T) {
	calc := calculator{}

	calc.numberStack.Push(1)
	calc.numberStack.Push(2)
	calc.numberStack.Push(3)
	calc.performMultiOperation("++")
	if calc.numberStack.Pop() != 6 {
		t.Error("1 + 2 + 3 sollte 6 sein")
	}
}

func TestProductOperation(t *testing.T) {
	calc := calculator{}

	calc.numberStack.Push(2)
	calc.numberStack.Push(3)
	calc.numberStack.Push(4)
	calc.performMultiOperation("**")
	if calc.numberStack.Pop() != 24 {
		t.Error("2 * 3 * 4 sollte 24 sein")
	}
}

func TestHandleWrongInput(t *testing.T) {
	calc := calculator{}
	calc.handleNumberInput("abc")
	if len(calc.numberStack) != 0 {
		t.Error("Ungültige Eingabe sollte nicht auf den Stack gelangen")
	}
}

func TestIntegration(t *testing.T) {
	calc := calculator{}

	calc.handleNumberInput("2")
	calc.handleNumberInput("3")
	calc.checkInput("+")
	if calc.numberStack.Pop() != 5 {
		t.Error("2 + 3 sollte 5 sein")
	}

	calc.handleNumberInput("4")
	calc.checkInput("sqrt")
	if calc.numberStack.Pop() != 2 {
		t.Error("sqrt(4) sollte 2 sein")
	}
}

func TestInfixNotationOutput(t *testing.T) {
	calc := calculator{}

	calc.handleNumberInput("3")
	calc.handleNumberInput("5")
	calc.performBinaryOperation("+")
	expected := "(3 + 5)"
	if calc.history.Top() != expected {
		t.Errorf("Erwartete Infix-Notation: %s, aber erhalten: %s", expected, calc.history.Top())
	}

	calc.handleNumberInput("2")
	calc.performBinaryOperation("*")
	expected = "((3 + 5) * 2)"
	if calc.history.Top() != expected {
		t.Errorf("Erwartete Infix-Notation: %s, aber erhalten: %s", expected, calc.history.Top())
	}

	calc.handleNumberInput("2")
	calc.performBinaryOperation("^")
	expected = "(((3 + 5) * 2) ^ 2)"
	if calc.history.Top() != expected {
		t.Errorf("Erwartete Infix-Notation: %s, aber erhalten: %s", expected, calc.history.Top())
	}
}

func TestComplexCalculationWithFactorial(t *testing.T) {
	calc := calculator{}

	calc.handleNumberInput("3")
	calc.performUnaryOperation("!")
	calc.handleNumberInput("2")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("+")
	if calc.numberStack.Pop() != 8 {
		t.Error("3! + 2! sollte 8 sein")
	}

	calc.handleNumberInput("5")
	calc.performUnaryOperation("!")
	calc.handleNumberInput("4")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("-")
	calc.handleNumberInput("2")
	calc.performUnaryOperation("!")
	calc.performBinaryOperation("*")
	if calc.numberStack.Pop() != 192 {
		t.Error("(5! - 4!) * 2! sollte 192 sein")
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
	if calc.numberStack.Pop() != 768 {
		t.Error("(5! - 4!) * (3! + 2!) sollte 768 sein")
	}
}

func setupCalculator() *calculator {
	return &calculator{
		numberStack: Stack[float64]{},
		history:     Stack[string]{},
		latex:       Stack[string]{},
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
		t.Error("help command did not produce expected output")
	}

	// Test latex
	c.latex.Push("x^2")
	output = captureOutput(func() {
		c.checkInput("latex")
	})
	if !strings.Contains(output, "\\[x^2\\]") {
		t.Error("latex command did not print expected LaTeX output")
	}
}

func TestPerformMultiOperation_TooFewItems(t *testing.T) {
	c := setupCalculator()
	output := captureOutput(func() {
		c.checkInput("++")
	})
	if !strings.Contains(output, "at least 2 numbers") {
		t.Error("multi-op with <2 items should produce error")
	}
}

func TestHandleNumberInput(t *testing.T) {
	c := setupCalculator()
	c.handleNumberInput("42.5")
	if len(c.numberStack) != 1 || c.numberStack.Top() != 42.5 {
		t.Error("handleNumberInput did not push correct number")
	}
}

func TestRestoreState(t *testing.T) {
	c := setupCalculator()
	c.restoreState("term", "latex", 5.0)
	if c.numberStack.Top() != 5.0 || c.history.Top() != "term" || c.latex.Top() != "latex" {
		t.Error("restoreState did not restore correct values")
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
		t.Error("abs did not calculate correctly")
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
		t.Error("sqrt error handling failed for negative number")
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
		t.Error("log normal calculation failed")
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
		t.Error("log error handling failed")
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
		t.Error("factorial error handling failed")
	}
}

func TestBinaryOperation_NoNumbers(t *testing.T) {
	c := setupCalculator()
	output := captureOutput(func() {
		c.performBinaryOperation("+")
	})
	if !strings.Contains(output, "at least 2 numbers") {
		t.Error("binary operation should fail with no numbers")
	}
}
