package maths

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	f := Func{Var{}}
	if f.Latex() != "\\( f(x) = x \\)" {
		t.Errorf("Expected \\( f(x) = x \\), got %s", f.Latex())
	}
	if fmt.Sprintf("%v", f) != "f(x) = x" {
		t.Errorf("Expected f(x) = x, got %s", f)
	}
	if f.Derive().Eval(1) != 1 {
		t.Errorf("Expected 0, got %v", f.Derive().Eval(1))
	}
}

func TestConst(t *testing.T) {
	c := Const{5}
	if c.Eval(0) != 5 {
		t.Errorf("Expected 5, got %f", c.Eval(0))
	}
	if c.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", c.Derive().Eval(0))
	}
}

func TestVar(t *testing.T) {
	v := Var{}
	if v.Eval(3) != 3 {
		t.Errorf("Expected 3, got %f", v.Eval(3))
	}
	if v.Derive().Eval(0) != 1 {
		t.Errorf("Expected derivative 1, got %f", v.Derive().Eval(0))
	}
}

func TestAdd(t *testing.T) {
	add := Add{Const{2}, Const{3}}
	if add.Eval(0) != 5 {
		t.Errorf("Expected 5, got %f", add.Eval(0))
	}
	if add.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", add.Derive().Eval(0))
	}
	if add.Latex() != "2 + 3" {
		t.Errorf("Expected 2 + 3, got %v", add.Latex())
	}

	add = Add{Var{}, Const{20}}
	if add.Derive().Eval(1) != 1 {
		t.Errorf("Expected 1, got %s", add.Latex())
	}

	add = Add{Var{}, Var{}}
	if add.Derive().Eval(1) != 2 {
		t.Errorf("Expected 2, got %v", add.Derive().Eval(1))
	}
}

func TestSub(t *testing.T) {
	sub := Sub{Const{5}, Const{2}}
	if sub.Eval(0) != 3 {
		t.Errorf("Expected 3, got %f", sub.Eval(0))
	}
	if sub.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", sub.Derive().Eval(0))
	}

	sub = Sub{Var{}, Const{10}}
	if sub.Derive().Latex() != "1" {
		t.Errorf("Expected 1, got %s", sub.Derive().Latex())
	}

	sub = Sub{Var{}, Var{}}
	if sub.Derive().Eval(1) != 0 {
		t.Errorf("Expected 0, got %v", sub.Derive().Eval(1))
	}
	if sub.Latex() != "x - x" {
		t.Errorf("Expected x - x, got %v", sub.Latex())
	}
}

func TestMult(t *testing.T) {
	mult := Mult{Const{3}, Const{4}}
	if mult.Eval(0) != 12 {
		t.Errorf("Expected 12, got %f", mult.Eval(0))
	}
	if mult.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", mult.Derive().Eval(0))
	}

	mult = Mult{Var{}, Var{}}
	if mult.Derive().Latex() != "x + x" {
		t.Errorf("Expected x + x, got %v", mult.Derive().Latex())
	}
	if mult.Latex() != "x \\cdot x" {
		t.Errorf("Expected x \\cdot x, got %v", mult.Latex())
	}

	mult = Mult{Var{}, Const{5}}
	if mult.Derive().Latex() != "5" {
		t.Errorf("Expected 5, got %v", mult.Derive().Latex())
	}

	mult = Mult{Const{5}, Var{}}
	if mult.Derive().Latex() != "5" {
		t.Errorf("Expected 5, got %v", mult.Derive().Latex())
	}

	mult = Mult{Mult{Const{2}, Var{}}, Mult{Const{2}, Var{}}}
	if mult.Derive().Latex() != "2 \\cdot 2 \\cdot x + 2 \\cdot x \\cdot 2" {
		t.Errorf("Expected a, got %v", mult.Derive().Latex())
	}
}

func TestDiv(t *testing.T) {
	div := Div{Const{10}, Const{2}}
	if div.Eval(0) != 5 {
		t.Errorf("Expected 5, got %f", div.Eval(0))
	}
	if div.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", div.Derive().Eval(0))
	}
	if div.Latex() != "\\frac{10}{2}" {
		t.Errorf("Expected \\frac{10}{2}, got %v", div.Latex())
	}
	div = Div{Var{}, Var{}}
	l := "\\frac{1 \\cdot x - 1 \\cdot x}{x ^ 2}"
	if div.Derive().Latex() != l {
		t.Errorf("Expected %v, got %v", l, div.Derive().Latex())
	}

	div = Div{Const{10}, Var{}}
	l = "0 - \\frac{10}{x ^ 2}"
	if div.Derive().Latex() != l {
		t.Errorf("Expected %v, got %v", l, div.Derive().Latex())
	}

	div = Div{Var{}, Const{10}}
	l = "\\frac{1}{10}"
	if div.Derive().Latex() != l {
		t.Errorf("Expected %v, got %v", l, div.Derive().Latex())
	}
}

func TestPow(t *testing.T) {
	pow := Pow{Const{2}, Const{3}}
	if pow.Eval(0) != 8 {
		t.Errorf("Expected 8, got %f", pow.Eval(0))
	}
	if pow.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", pow.Derive().Eval(0))
	}

	pow = Pow{Var{}, Const{10}}
	if pow.Eval(0) != 0 {
		t.Errorf("Expected 0, got %f", pow.Eval(0))
	}
	if pow.Derive().Eval(1) != 10 {
		t.Errorf("Expected 10, got %f", pow.Eval(1))
	}
	if pow.Latex() != "x ^ 10" {
		t.Errorf("Expected x ^ 10, got %s", pow.Latex())
	}
}

func TestSqrt(t *testing.T) {
	sqr := Sqrt{Const{9}}
	if sqr.Eval(0) != 3 {
		t.Errorf("Expected 3, got %f", sqr.Eval(0))
	}
	if sqr.Derive().Eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", sqr.Derive().Eval(0))
	}
	if sqr.Latex() != "\\sqrt{9}" {
		t.Errorf("Expected \\sqrt{9}, got %v", sqr.Latex())
	}
	sqr = Sqrt{Var{}}
	if sqr.Derive().Eval(1) != 0.5 {
		t.Errorf("Expected 1, got %v", sqr.Derive().Eval(1))
	}
}
