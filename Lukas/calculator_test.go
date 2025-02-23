package main

import (
	"testing"
)

// Test für Stack-Funktionen
func TestStackOperations(t *testing.T) {
	stack := Stack[float64]{}

	stack.Push(3.0)
	stack.Push(5.0)

	if stack.Pop() != 5.0 {
		t.Error("Pop() sollte 5.0 zurückgeben")
	}

	if stack.Top() != 3.0 {
		t.Error("Top() sollte 3.0 zurückgeben")
	}

	stack.Pop()
	if len(stack) != 0 {
		t.Error("Stack sollte nach Pop() leer sein")
	}
}

// Test für binäre Operationen (+, -, *, /, ^)
func TestBinaryOperations(t *testing.T) {
	calc := calculator{}

	// Für diese Tests ist nur das Ergebnis wichtig, daher direkten Push in den numberStack
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

// Test für Fakultät (!)
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

// Test für Summenoperation (++ = Summe aller Stack-Elemente)
func TestSumOperation(t *testing.T) {
	calc := calculator{}

	// Hier werden direkt Zahlen in den numberStack gepusht, da der History-String nicht relevant ist
	calc.numberStack.Push(1)
	calc.numberStack.Push(2)
	calc.numberStack.Push(3)
	calc.performSumOperation()
	if calc.numberStack.Pop() != 6 {
		t.Error("1 + 2 + 3 sollte 6 sein")
	}
}

// Test für Multiplikation aller Stack-Elemente (**)
func TestProductOperation(t *testing.T) {
	calc := calculator{}

	calc.numberStack.Push(2)
	calc.numberStack.Push(3)
	calc.numberStack.Push(4)
	calc.performProductOperation()
	if calc.numberStack.Pop() != 24 {
		t.Error("2 * 3 * 4 sollte 24 sein")
	}
}

// Test für falsche Eingaben
func TestHandleWrongInput(t *testing.T) {
	calc := calculator{}
	calc.handleNumberInput("abc")
	if len(calc.numberStack) != 0 {
		t.Error("Ungültige Eingabe sollte nicht auf den Stack gelangen")
	}
}

// Integrationstest: mehrere Operationen zusammen testen
func TestIntegration(t *testing.T) {
	calc := calculator{}

	// 2 3 +
	calc.handleNumberInput("2")
	calc.handleNumberInput("3")
	calc.checkInput("+")
	if calc.numberStack.Pop() != 5 {
		t.Error("2 + 3 sollte 5 sein")
	}

	// 4 sqrt (Ergebnis sollte 2 sein)
	calc.handleNumberInput("4")
	calc.checkInput("sqrt")
	if calc.numberStack.Pop() != 2 {
		t.Error("sqrt(4) sollte 2 sein")
	}
}

// Test für die korrekte Infix-Notation in der History
func TestInfixNotationOutput(t *testing.T) {
	calc := calculator{}

	// Test für eine binäre Operation (Addition)
	calc.handleNumberInput("3")
	calc.handleNumberInput("5")
	calc.performBinaryOperation("+")
	expected := "(3 + 5)"
	if calc.history.Top() != expected {
		t.Errorf("Erwartete Infix-Notation: %s, aber erhalten: %s", expected, calc.history.Top())
	}

	// Test für eine komplexere Rechnung: ((3 + 5) * 2)
	calc.handleNumberInput("2")
	calc.performBinaryOperation("*")
	expected = "((3 + 5) * 2)"
	if calc.history.Top() != expected {
		t.Errorf("Erwartete Infix-Notation: %s, aber erhalten: %s", expected, calc.history.Top())
	}

	// Test für Exponentiation: (((3 + 5) * 2) ^ 2)
	calc.handleNumberInput("2")
	calc.performBinaryOperation("^")
	expected = "(((3 + 5) * 2) ^ 2)"
	if calc.history.Top() != expected {
		t.Errorf("Erwartete Infix-Notation: %s, aber erhalten: %s", expected, calc.history.Top())
	}
}
