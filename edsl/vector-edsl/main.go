package main

import (
	"fmt"
)

func main() {
	a := svg{ width: 200, height: 200,
		content: []Element{
			rect{width: 200, height: 200, fill:"lime"},
			rect{width: 100, height: 100, fill:"blue"},
			circle{r: 45, cx: 50, cy: 50, fill: "blue"},
			line{0, 0, 200, 200, "stroke:coral;stroke-width:14"},
			text{x: 75, y: 100, fill: "white", content: "Wassup!!!!"},
		},
	}
	a.saveSVG("shapess")
	fmt.Println(a.toSVG())
	fmt.Println(a)
}
