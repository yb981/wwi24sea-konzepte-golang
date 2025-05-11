// maths-func.go
//
// Framework für mathematische Ausdrücke, das die Evaluation, Ableitung und LaTeX-Darstellung verschiedener algebraischer Operationen ermöglicht.
//
// Author: Kevin Becker
// Date: 01.04.2025
package maths

import (
	"fmt"
	"math"
)

// Interface, which all Expression types implement

type Expression interface {
	Eval(num float64) float64
	Derive() Expression
	Latex() string
}

// Mathematical Function value: Starting Point for eDSL

type Func struct {
	Fn Expression
}

func (f Func) String() string {
	return fmt.Sprintf(f.Latex()[3 : len(f.Latex())-3])
}

func (f Func) Eval(num float64) float64 {
	return f.Fn.Eval(num)
}

func (f Func) Derive() Expression {
	return Func{Fn: f.Fn.Derive()}
}

func (f Func) Latex() string {
	return fmt.Sprintf("\\( f(x) = %s \\)", f.Fn.Latex())
}

// Used for Variable x

type Var struct{}

func (v Var) Eval(num float64) float64 {
	return num
}

func (v Var) Derive() Expression {
	return Const{Val: 1}
}

func (v Var) Latex() string {
	return "x"
}

// Used for Constant Value c

type Const struct {
	Val float64
}

func (c Const) Eval(num float64) float64 {
	return c.Val
}

func (c Const) Derive() Expression {
	return Const{Val: 0}
}

func (c Const) Latex() string {
	return fmt.Sprintf("%g", c.Val)
}

/*
* Begin of Operator List
 */

// Addition

type Add struct {
	Left, Right Expression
}

func (a Add) Eval(num float64) float64 {
	return a.Left.Eval(num) + a.Right.Eval(num)
}

func (a Add) Derive() Expression {
	if a.Left.Derive().Latex() == "0" {
		return a.Right.Derive()
	}
	if a.Right.Derive().Latex() == "0" {
		return a.Left.Derive()
	}
	return Add{Left: a.Left.Derive(), Right: a.Right.Derive()}
}

func (a Add) Latex() string {
	return fmt.Sprintf("%s + %s", a.Left.Latex(), a.Right.Latex())
}

// Subtraction

type Sub struct {
	Left, Right Expression
}

func (s Sub) Eval(num float64) float64 {
	return s.Left.Eval(num) - s.Right.Eval(num)
}

func (s Sub) Derive() Expression {
	if s.Right.Derive().Latex() == "0" {
		return s.Left.Derive()
	}
	return Sub{Left: s.Left.Derive(), Right: s.Right.Derive()}
}

func (s Sub) Latex() string {
	return fmt.Sprintf("%s - %s", s.Left.Latex(), s.Right.Latex())
}

// Multiplication

type Mult struct {
	Left, Right Expression
}

func (m Mult) Eval(num float64) float64 {
	return m.Left.Eval(num) * m.Right.Eval(num)
}

func (m Mult) Derive() Expression {
	if m.Left.Derive().Eval(1) > 0 && m.Right.Derive().Eval(1) > 0 {
		return Add{
			Left:  checkRedundancyMult(Mult{Left: m.Left.Derive(), Right: m.Right}),
			Right: checkRedundancyMult(Mult{Left: m.Left, Right: m.Right.Derive()}),
		}
	} else if m.Left.Derive().Eval(1) > 0 {
		return checkRedundancyMult(Mult{Left: m.Left.Derive(), Right: m.Right})
	} else if m.Right.Derive().Eval(1) > 0 {
		return checkRedundancyMult(Mult{Left: m.Left, Right: m.Right.Derive()})
	} else {
		return Const{Val: 0}
	}
}

func checkRedundancyMult(m Mult) Expression {
	if m.Left.Latex() == "1" {
		return m.Right
	} else if m.Right.Latex() == "1" {
		return m.Left
	}
	return m
}

func (m Mult) Latex() string {
	return fmt.Sprintf("%s \\cdot %s", m.Left.Latex(), m.Right.Latex())
}

// Division

type Div struct {
	Left, Right Expression
}

func (d Div) Eval(num float64) float64 {
	return d.Left.Eval(num) / d.Right.Eval(num)
}

func (d Div) Derive() Expression {
	if d.Left.Derive().Eval(0) != 0 && d.Right.Derive().Eval(0) != 0 {
		return Div{Left: Sub{Left: Mult{Left: d.Left.Derive(), Right: d.Right}, Right: Mult{Left: d.Left.Derive(), Right: d.Right}}, Right: Pow{Val: d.Right, Exp: Const{Val: 2}}}
	} else if d.Left.Derive().Eval(0) != 0 {
		return Div{Left: d.Left.Derive(), Right: d.Right}
	} else if d.Right.Derive().Eval(0) != 0 {
		return Sub{Left: Const{Val: 0}, Right: Div{Left: d.Left, Right: Pow{Val: Var{}, Exp: Const{Val: 2}}}}
	}
	return Const{Val: 0}
}

func (d Div) Latex() string {
	return fmt.Sprintf("\\frac{%s}{%s}", d.Left.Latex(), d.Right.Latex())
}

// Power

type Pow struct {
	Val Expression
	Exp Const
}

func (p Pow) Eval(num float64) float64 {
	return math.Pow(p.Val.Eval(num), p.Exp.Eval(num))
}

func (p Pow) Derive() Expression {
	if p.Val.Latex() == "x" {
		return Mult{Left: p.Exp, Right: Pow{Val: p.Val, Exp: Const{Val: p.Exp.Val - 1}}}
	} else {
		return Const{Val: 0}
	}
}

func (p Pow) Latex() string {
	return fmt.Sprintf("%s ^ %s", p.Val.Latex(), p.Exp.Latex())
}

// Root

type Sqrt struct {
	Val Expression
}

func (s Sqrt) Eval(num float64) float64 {
	return math.Sqrt(s.Val.Eval(num))
}

func (s Sqrt) Derive() Expression {
	if s.Val.Eval(0) == 0 {
		return Div{Left: Const{Val: 1}, Right: Mult{Left: Const{Val: 2}, Right: s}}
	}
	return Const{Val: 0}
}

func (s Sqrt) Latex() string {
	return fmt.Sprintf("\\sqrt{%s}", string(s.Val.Latex()[0]))
}

/*
* End of Operator List
 */
