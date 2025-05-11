// main.go
// Startet das eDSL Vector Demo
//
// Author: Kevin Becker
// Date: 01.04.2025
package main

import (
	"fmt"

	. "github.com/yb981/wwi24sea-konzepte-golang/edsl/vector-edsl"
)

func main() {
	fmt.Printf("-------------------------------------- LIVE DEMO SHOWCASE -----------------------------------\n\n")

	fmt.Printf("------------------------------------- FIRST SVG SCHOWCASE -----------------------------------\n\n")

	a := Svg{Width: 200, Height: 200,
		Content: []Element{
			Rect{Width: 200, Height: 200, Fill: "lime"},
			Rect{Width: 100, Height: 100, Fill: "blue"},
			Circle{R: 45, Cx: 50, Cy: 50, Fill: "blue"},
			Line{0, 0, 200, 200, "stroke:coral;stroke-Width:14"},
			Text{X: 75, Y: 100, Fill: "white", Content: "Wassup!!!!", Size: 20},
		},
	}

	a.SaveSVG("shapes")
	fmt.Println(a.ToSVG())

	fmt.Printf("------------------------------------ SECOND SVG SHOWCASE -----------------------------------\n\n")

	b := Svg{200, 200, []Element{
		Rect{Width: 200, Height: 200},
		Circle{R: 100, Fill: "yellow"},
		Ellipse{Rx: 20, Ry: 20, Style: "Fill:lime"},
	}}

	fmt.Println(b.ToSVG())
	b.SaveSVG("shapes2")

	fmt.Printf("------------------------------------ THIRD SVG SHOWCASE -----------------------------------\n\n")

	c := Svg{1000, 1000, []Element{
		Rect{Width: 1000, Height: 1000, Fill: "purple"},
		Rect{Width: 900, Height: 900, X: 50, Y: 50, Fill: "teal"},
		Circle{R: 500, Cx: 500, Cy: 500},
		Rect{Width: 900, Height: 300, X: 50, Y: 350, Fill: "white"},
		Text{X: 250, Y: 525, Content: "SVG is cool!", Fill: "black", Size: 100},
	}}

	fmt.Println(c.ToSVG())
	c.SaveSVG("shapes3")
}
