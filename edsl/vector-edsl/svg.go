// svg.go
//
// SVG-Generierungspaket zur Erzeugung von SVG-Dateien mit verschiedenen SVG-Elemen
//
// Author: Kevin Becker, Lukas Grönning

package svg

import (
	"fmt"
	"os"
)

// Interface which all SVG Elements implement
type Element interface {
	ToSVG() string
}

// Entry Point for SVG-Generation: SVG Struct
type Svg struct {
	Width, Height int
	Content       []Element
}

func (s Svg) SaveSVG(name string) {
	fileName := fmt.Sprintf("%v.svg", name)
	os.WriteFile(fileName, []byte(s.ToSVG()), 0644)
}

func (s Svg) ToSVG() string {
	res := ""
	for _, v := range s.Content {
		res += v.ToSVG() + "\n"
	}
	return fmt.Sprintf(`<svg width="%v" height="%v" xmlns="http://www.w3.org/2000/svg">
%s</svg>`, s.Width, s.Height, res)
}

// Rectangle Element
type Rect struct {
	Width, Height, X, Y, Rx, Ry int
	Fill                        string
}

func (r Rect) ToSVG() string {
	return fmt.Sprintf(`	<rect width="%v" height="%v" x="%v" y="%v" rx="%v" ry="%v" fill="%v" />`,
		r.Width, r.Height, r.X, r.Y, r.Rx, r.Ry, r.Fill)
}

// Circle Element
type Circle struct {
	R, Cx, Cy int
	Fill      string
}

func (c Circle) ToSVG() string {
	return fmt.Sprintf(`	<circle r="%v" cx="%v" cy="%v" fill="%v" />`,
		c.R, c.Cx, c.Cy, c.Fill)
}

// Line Element
type Line struct {
	X1, Y1, X2, Y2 int
	Style          string
}

func (l Line) ToSVG() string {
	return fmt.Sprintf(`	<line x1="%v" y1="%v" x2="%v" y2="%v" style="%s" />`,
		l.X1, l.Y1, l.X2, l.Y2, l.Style)
}

// Text Element
type Text struct {
	X, Y, Size    int
	Fill, Content string
}

func (t Text) ToSVG() string {
	return fmt.Sprintf(`	<text x="%v" y="%v" fill="%v" font-size="%v">%s</text>`,
		t.X, t.Y, t.Fill, t.Size, t.Content)
}

// Ellipse Element
type Ellipse struct {
	Rx, Ry, Cx, Cy int
	Style          string
}

func (e Ellipse) ToSVG() string {
	return fmt.Sprintf(`	<ellipse rx="%v" ry="%v" cx="%v" cy="%v" style="%v" />`,
		e.Rx, e.Ry, e.Cx, e.Cy, e.Style)
}
