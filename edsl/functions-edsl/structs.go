package main

import (
	"math"
	"fmt"
)

// 1. Ansatz: Mit Interfaces & Structs

type Expression interface {
	eval(num float64) float64
	derive() Expression
	latex() string
}

// Function
type Func struct{
	fn Expression
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

// Variable
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

// Constant

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
	return fmt.Sprintf("%.2f", c.val)
}

// Addition

type Add struct {
	left, right Expression
}

func (a Add) eval(num float64) float64 {
	return a.left.eval(num) + a.right.eval(num)
}

func (a Add) derive() Expression {
	if a.left.derive().latex() == "0.00" {
		return a.right.derive()
	}
	if a.right.derive().latex() == "0.00" {
		return a.left.derive()
	}
	return Add{a.left.derive(), a.right.derive()}
}

func (a Add) latex() string {
	return fmt.Sprintf("%s + %s", a.left.latex(), a.right.latex())
}

// Subtract

type Sub struct {
	left, right Expression
}

func (s Sub) eval(num float64) float64 {
	return s.left.eval(num) - s.right.eval(num)
}

func (s Sub) derive() Expression {
	if s.right.derive().latex() == "0.00" {
		return s.left.derive()
	}
	return Sub{s.left.derive(), s.right.derive()}
}

func (s Sub) latex() string {
	return fmt.Sprintf("%s - %s", s.left.latex(), s.right.latex())
}
// Multiply

type Mult struct {
	left, right Expression
}

func (m Mult) eval(num float64) float64 {
	return m.left.eval(num) * m.right.eval(num)
}

func (m Mult) derive() Expression {
	if m.left.derive().eval(1) > 0 && m.right.derive().eval(1) > 0 {
		return Add{m.left.derive(), m.right.derive()}
	} else if m.left.derive().eval(1) > 0 {
		return Mult{m.left.derive(), m.right}
	} else if m.right.derive().eval(1) > 0 {
		return Mult{m.left, m.right.derive()}
	} else {
		return Const{0}
	}
}

func (m Mult) latex() string {
	return fmt.Sprintf("%s * %s", m.left.latex(), m.right.latex())
}

// Divide

type Div struct {
	left, right Expression
}

func (d Div) eval(num float64) float64 {
	return d.left.eval(num) / d.right.eval(num)
}

func (d Div) derive() Expression{
	return Const{1}
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
	return Const{1}
}

func (p Pow) latex() string {
	return fmt.Sprintf("%s ^ %s", p.val.latex(), string(p.exp.latex()[0]))
}

// Root

type Sqr struct {
	val Expression
}

func (s Sqr) eval(num float64) float64 {
	return math.Sqrt(s.val.eval(num))
}

func (s Sqr) derive() Expression {
	if s.val.eval(0) == 0 {
		return Div{Const{1}, Mult{Const{2}, s}}
	}
	return Const{0}
}

func (s Sqr) latex() string {
	return fmt.Sprintf("\\sqrt{%s}", string(s.val.latex()[0]))
}
