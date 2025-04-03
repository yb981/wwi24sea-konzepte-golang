package main

import (
	"fmt"
)

func main() {
	fmt.Printf("-------------------------------------- LIVE DEMO SHOWCASE -----------------------------------\n\n")

	fmt.Printf("------------------------------------- FIRST SVG SCHOWCASE -----------------------------------\n\n")

	a := svg{width: 200, height: 200,
		content: []Element{
			rect{width: 200, height: 200, fill: "lime"},
			rect{width: 100, height: 100, fill: "blue"},
			circle{r: 45, cx: 50, cy: 50, fill: "blue"},
			line{0, 0, 200, 200, "stroke:coral;stroke-width:14"},
			text{x: 75, y: 100, fill: "white", content: "Wassup!!!!", size: 20},
		},
	}

	a.saveSVG("shapes")
	fmt.Println(a.toSVG())

	fmt.Printf("------------------------------------ SECOND SVG SHOWCASE -----------------------------------\n\n")

	b := svg{200, 200, []Element{
		rect{width: 200, height: 200},
		circle{r: 100, fill: "yellow"},
		ellipse{rx: 20, ry: 20, style: "fill:lime"},
	}}

	fmt.Println(b.toSVG())
	b.saveSVG("shapes2")

	fmt.Printf("------------------------------------ THIRD SVG SHOWCASE -----------------------------------\n\n")

	c := svg{1000, 1000, []Element{
		rect{width: 1000, height: 1000, fill: "purple"},
		rect{width: 900, height: 900, x: 50, y: 50, fill: "teal"},
		circle{r: 500, cx: 500, cy: 500},
		rect{width: 900, height: 300, x: 50, y: 350, fill: "white"},
		text{x: 250, y: 525, content: "SVG is cool!", fill: "black", size: 100},
	}}

	fmt.Println(c.toSVG())
	c.saveSVG("shapes3")
}
