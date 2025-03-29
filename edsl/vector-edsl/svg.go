package main

import (
	"fmt"
	"os"
)

type Element interface {
	toSVG() string
}

//1. Ansatz: Interfaces & Receiver Funktionen

type svg struct {
	width, height int
	content       []Element
}

func (s svg) String() string {
	return fmt.Sprintf("SVG Format at adress %p", &s)
}

func (s svg) saveSVG() {
	err := os.WriteFile("shapes.svg", []byte(s.toSVG()), 0644)
	if err != nil {
		fmt.Errorf("Writing the file was not successful.")
	}
}

func (s svg) toSVG() string {
	res := ""
	for _, v := range s.content {
		res += v.toSVG() + "\n"
	}
	return fmt.Sprintf(`<svg width="%v" height="%v">
%s</svg>`, s.width, s.height, res)
}

type rect struct {
	width  int
	height int
	fill   string
}

func (r rect) toSVG() string {
	return fmt.Sprintf(`	<rect width="%v" height="%v" fill="%v" />`,
		r.width, r.height, r.fill)
}

type circle struct {
	r, cx, cy int
	fill      string
}

func (c circle) toSVG() string {
	return fmt.Sprintf(`	<circle r="%v" cx="%v" cy="%v" fill="%v" />`,
		c.r, c.cx, c.cy, c.fill)
}

type line struct {
	x1, y1, x2, y2 int
	style          string
}

func (l line) toSVG() string {
	return fmt.Sprintf(`	<line x1="%v" y1="%v" x2="%v" y2="%v" style="%s" />`,
		l.x1, l.y1, l.x2, l.y2, l.style)
}

type text struct {
	x, y, dx, dy, rotate int
	fill, content        string
}

func (t text) toSVG() string {
	return fmt.Sprintf(`	<text x="%v" y="%v" fill="%v">%s</text>`,
		t.x, t.y, t.fill, t.content)
}
