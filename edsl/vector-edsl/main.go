package main

import (
	"fmt"
)

func main() {
	a := svg{ width: 200, height: 200,
		content: []Element{
			rect{200, 200, "red"},
			rect{100, 100, "green"},
			circle{r: 45, cx: 50, cy: 50, fill: "blue"},
			line{0, 0, 200, 200, "stroke:coral;stroke-width:14"},
			text{x: 75, y: 100, fill: "white", content: "Wassup!!!!"},
		},
	}
	a.saveSVG()
	fmt.Println(a.toSVG())
	fmt.Println(a)
}
