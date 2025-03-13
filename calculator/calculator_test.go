package main

import (
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
