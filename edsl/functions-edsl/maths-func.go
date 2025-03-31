package main

import (
	"fmt"
	"math"
)

// Interface, which all Expression types implement

type Expression interface {
	eval(num float64) float64
	derive() Expression
	latex() string
}

// Mathematical Function value: Starting Point for eDSL

type Func struct {
	fn Expression
}

func (f Func) String() string {
	return fmt.Sprintf(f.latex()[3 : len(f.latex())-3])
}

func (f Func) eval(num float64) float64 {
	return f.fn.eval(num)
}

func (f Func) derive() Expression {
	return Func{f.fn.derive()}
}

func (f Func) latex() string {
	return fmt.Sprintf("\\( f(x) = %s \\)", f.fn.latex())
}

// Used for Variable x

type Var struct{}

func (v Var) eval(num float64) float64 {
	return num
}

func (v Var) derive() Expression {
	return Const{1}
}

func (v Var) latex() string {
	return "x"
}

// Used for Constant Value c

type Const struct {
	val float64
}

func (c Const) eval(num float64) float64 {
	return c.val
}

func (c Const) derive() Expression {
	return Const{0}
}

func (c Const) latex() string {
	return fmt.Sprintf("%g", c.val)
}

/*
*	Begin of Operator List
*/

// Addition

type Add struct {
	left, right Expression
}

func (a Add) eval(num float64) float64 {
	return a.left.eval(num) + a.right.eval(num)
}

func (a Add) derive() Expression {

	// Simplification logic

	if a.left.derive().latex() == "0" {
		return a.right.derive()
	}
	if a.right.derive().latex() == "0" {
		return a.left.derive()
	}

	return Add{a.left.derive(), a.right.derive()}
}

func (a Add) latex() string {
	return fmt.Sprintf("%s + %s", a.left.latex(), a.right.latex())
}

// Subtraction

type Sub struct {
	left, right Expression
}

func (s Sub) eval(num float64) float64 {
	return s.left.eval(num) - s.right.eval(num)
}

func (s Sub) derive() Expression {

	//Simplification Logic

	if s.right.derive().latex() == "0" {
		return s.left.derive()
	}

	return Sub{s.left.derive(), s.right.derive()}
}

func (s Sub) latex() string {
	return fmt.Sprintf("%s - %s", s.left.latex(), s.right.latex())
}

// Multiplication

type Mult struct {
	left, right Expression
}

func (m Mult) eval(num float64) float64 {
	return m.left.eval(num) * m.right.eval(num)
}

func (m Mult) derive() Expression {

	// Simplification Logic

	if m.left.derive().eval(1) > 0 && m.right.derive().eval(1) > 0 {
		return Add{
			checkRedundancyMult(Mult{m.left.derive(), m.right}), 
			checkRedundancyMult(Mult{m.left, m.right.derive()}),
		}
	} else if m.left.derive().eval(1) > 0 {
		return checkRedundancyMult(Mult{m.left.derive(), m.right})
	} else if m.right.derive().eval(1) > 0 {
		return checkRedundancyMult(Mult{m.left, m.right.derive()})
	} else {
		return Const{0}
	}
}

func checkRedundancyMult(m Mult) Expression {
	if m.left.latex() == "1" {
		return m.right
	} else if m.right.latex() == "1" {
		return m.left
	}
	return m
}

func (m Mult) latex() string {
	return fmt.Sprintf("%s \\cdot %s", m.left.latex(), m.right.latex())
}

// Division

type Div struct {
	left, right Expression
}

func (d Div) eval(num float64) float64 {
	return d.left.eval(num) / d.right.eval(num)
}

func (d Div) derive() Expression {
	if d.left.derive().eval(0) != 0 && d.right.derive().eval(0) != 0 {
		return Div{Sub{Mult{d.left.derive(), d.right}, Mult{d.left.derive(), d.right}}, Pow{d.right, Const{2}}}
	} else if d.left.derive().eval(0) != 0 {
		return Div{d.left.derive(), d.right}
	} else if d.right.derive().eval(0) != 0 {
		return Sub{Const{0}, Div{d.left, Pow{Var{}, Const{2}}}}
	}
	return Const{0}
}

func (d Div) latex() string {
	return fmt.Sprintf("\\frac{%s}{%s}", d.left.latex(), d.right.latex())
}

// Power

type Pow struct {
	val Expression
	exp Const
}

func (p Pow) eval(num float64) float64 {
	return math.Pow(p.val.eval(num), p.exp.eval(num))
}

func (p Pow) derive() Expression {
	if p.val.latex() == "x" {
		return Mult{p.exp, Pow{p.val, Const{p.exp.val - 1}}}
	} else {
		return Const{0}
	}
}

func (p Pow) latex() string {
	return fmt.Sprintf("%s ^ %s", p.val.latex(), p.exp.latex())
}

// Root

type Sqrt struct {
	val Expression
}

func (s Sqrt) eval(num float64) float64 {
	return math.Sqrt(s.val.eval(num))
}

func (s Sqrt) derive() Expression {
	if s.val.eval(0) == 0 {
		return Div{Const{1}, Mult{Const{2}, s}}
	}
	return Const{0}
}

func (s Sqrt) latex() string {
	return fmt.Sprintf("\\sqrt{%s}", string(s.val.latex()[0]))
}

/*
*	End of Operator List
*/