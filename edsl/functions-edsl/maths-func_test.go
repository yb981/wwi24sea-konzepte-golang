package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	f := Func{Var{}}
	if f.latex() != "\\( f(x) = x \\)" {
		t.Errorf("Expected \\( f(x) = x \\), got %s", f.latex())
	}
	if fmt.Sprintf("%v", f) != "f(x) = x" {
		t.Errorf("Expected f(x) = x, got %s", f)
	}
	if f.derive().eval(1) != 1 {
		t.Errorf("Expected 0, got %v", f.derive().eval(1))
	}
}

func TestConst(t *testing.T) {
	c := Const{5}
	if c.eval(0) != 5 {
		t.Errorf("Expected 5, got %f", c.eval(0))
	}
	if c.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", c.derive().eval(0))
	}
}

func TestVar(t *testing.T) {
	v := Var{}
	if v.eval(3) != 3 {
		t.Errorf("Expected 3, got %f", v.eval(3))
	}
	if v.derive().eval(0) != 1 {
		t.Errorf("Expected derivative 1, got %f", v.derive().eval(0))
	}
}

func TestAdd(t *testing.T) {
	add := Add{Const{2}, Const{3}}
	if add.eval(0) != 5 {
		t.Errorf("Expected 5, got %f", add.eval(0))
	}
	if add.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", add.derive().eval(0))
	}
	if add.latex() != "2 + 3" {
		t.Errorf("Expected 2 + 3, got %v", add.latex())
	}

	add = Add{Var{}, Const{20}}
	if add.derive().eval(1) != 1 {
		t.Errorf("Expected 1, got %s", add.latex())
	}

	add = Add{Var{}, Var{}}
	if add.derive().eval(1) != 2 {
		t.Errorf("Expected 2, got %v", add.derive().eval(1))
	}
}

func TestSub(t *testing.T) {
	sub := Sub{Const{5}, Const{2}}
	if sub.eval(0) != 3 {
		t.Errorf("Expected 3, got %f", sub.eval(0))
	}
	if sub.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", sub.derive().eval(0))
	}

	sub = Sub{Var{}, Const{10}}
	if sub.derive().latex() != "1" {
		t.Errorf("Expected 1, got %s", sub.derive().latex())
	}

	sub = Sub{Var{}, Var{}}
	if sub.derive().eval(1) != 0 {
		t.Errorf("Expected 0, got %v", sub.derive().eval(1))
	}
	if sub.latex() != "x - x" {
		t.Errorf("Expected x - x, got %v", sub.latex())
	}
}

func TestMult(t *testing.T) {
	mult := Mult{Const{3}, Const{4}}
	if mult.eval(0) != 12 {
		t.Errorf("Expected 12, got %f", mult.eval(0))
	}
	if mult.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", mult.derive().eval(0))
	}

	mult = Mult{Var{}, Var{}}
	if mult.derive().latex() != "x + x" {
		t.Errorf("Expected x + x, got %v", mult.derive().latex())
	}
	if mult.latex() != "x \\cdot x" {
		t.Errorf("Expected x \\cdot x, got %v", mult.latex())
	}

	mult = Mult{Var{}, Const{5}}
	if mult.derive().latex() != "5" {
		t.Errorf("Expected 5, got %v", mult.derive().latex())
	}

	mult = Mult{Const{5}, Var{}}
	if mult.derive().latex() != "5" {
		t.Errorf("Expected 5, got %v", mult.derive().latex())
	}

	mult = Mult{Mult{Const{2}, Var{}}, Mult{Const{2}, Var{}}}
	if mult.derive().latex() != "2 \\cdot 2 \\cdot x + 2 \\cdot x \\cdot 2" {
		t.Errorf("Expected a, got %v", mult.derive().latex())
	}
}

func TestDiv(t *testing.T) {
	div := Div{Const{10}, Const{2}}
	if div.eval(0) != 5 {
		t.Errorf("Expected 5, got %f", div.eval(0))
	}
	if div.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", div.derive().eval(0))
	}
	if div.latex() != "\\frac{10}{2}" {
		t.Errorf("Expected \\frac{10}{2}, got %v", div.latex())
	}
	div = Div{Var{}, Var{}}
	l := "\\frac{1 \\cdot x - 1 \\cdot x}{x ^ 2}"
	if div.derive().latex() != l {
		t.Errorf("Expected %v, got %v", l, div.derive().latex())
	}

	div = Div{Const{10}, Var{}}
	l = "0 - \\frac{10}{x ^ 2}"
	if div.derive().latex() != l {
		t.Errorf("Expected %v, got %v", l, div.derive().latex())
	}

	div = Div{Var{}, Const{10}}
	l = "\\frac{1}{10}"
	if div.derive().latex() != l {
		t.Errorf("Expected %v, got %v", l, div.derive().latex())
	}
}

func TestPow(t *testing.T) {
	pow := Pow{Const{2}, Const{3}}
	if pow.eval(0) != 8 {
		t.Errorf("Expected 8, got %f", pow.eval(0))
	}
	if pow.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", pow.derive().eval(0))
	}

	pow = Pow{Var{}, Const{10}}
	if pow.eval(0) != 0 {
		t.Errorf("Expected 0, got %f", pow.eval(0))
	}
	if pow.derive().eval(1) != 10 {
		t.Errorf("Expected 10, got %f", pow.eval(1))
	}
	if pow.latex() != "x ^ 10" {
		t.Errorf("Expected x ^ 10, got %s", pow.latex())
	}
}

func TestSqrt(t *testing.T) {
	sqr := Sqrt{Const{9}}
	if sqr.eval(0) != 3 {
		t.Errorf("Expected 3, got %f", sqr.eval(0))
	}
	if sqr.derive().eval(0) != 0 {
		t.Errorf("Expected derivative 0, got %f", sqr.derive().eval(0))
	}
	if sqr.latex() != "\\sqrt{9}" {
		t.Errorf("Expected \\sqrt{9}, got %v", sqr.latex())
	}
	sqr = Sqrt{Var{}}
	if sqr.derive().eval(1) != 0.5 {
		t.Errorf("Expected 1, got %v", sqr.derive().eval(1))
	}
}
