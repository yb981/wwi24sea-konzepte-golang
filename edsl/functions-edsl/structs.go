package main

import "math"

// 1. Ansatz: Mit Interfaces & Structs

type Expression interface {
	eval(num float64) float64
	derive() Expression
	latex() string
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
	return ""
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
	return ""
}

// Addition

type Add struct {
	left, right Expression
}

func (a Add) eval(num float64) float64 {
	return a.left.eval(num) + a.right.eval(num)
}

func (a Add) derive() Expression {
	return Add{a.left.derive(), a.right.derive()}
}

func (a Add) latex() string {
	return ""
}

// Subtract

type Sub struct {
	left, right Expression
}

func (s Sub) eval(num float64) float64 {
	return s.left.eval(num) - s.right.eval(num)
}

func (s Sub) derive() Expression {
	return Sub{s.left.derive(), s.right.derive()}
}

func (s Sub) latex() string {
	return ""
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
	return ""
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
	return ""
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
	return ""
}

// Root

type Sqr struct {
	val Expression
}

func (s Sqr) eval(num float64) float64 {
	return math.Sqrt(s.val.eval(num))
}

func (s Sqr) derive() Expression {
	return Const{1}
}

func (s Sqr) latex() string {
	return ""
}
