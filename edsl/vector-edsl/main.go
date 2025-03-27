package main

import(
    "fmt"
)

func main() {
    a := svg{ width: 200, height: 200,
        content: []Element{
            rect{200, 200, "red"},
            rect{100, 100, "green"},
            circle{r: 45, fill: "blue"},
            line{0, 0, 200, 200, "stroke:coral;stroke-width:14"},
        },
    }
    a.saveSVG()
    fmt.Println(a.toSVG())
    fmt.Println(a)
}