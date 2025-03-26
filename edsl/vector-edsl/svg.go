package main

import "fmt"

type Element interface{
	toSVG() string
}

//1. Ansatz: Interfaces & Receiver Funktionen

type svg struct{
	width, height int
	content []Element
}

func (s svg) toSVG() string {
	return fmt.Sprintf(`<svg>
	%s
</svg>`, s.content)
}

type rect struct{
	width, height int
}

func (r rect) toSVG() string {
	return fmt.Sprintf(``)
}