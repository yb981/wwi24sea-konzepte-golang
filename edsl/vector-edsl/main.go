package main

import(
    "fmt"
)

func main() {
    a := svg{ width: 200, height: 200 }
    fmt.Println(a.toSVG())
    fmt.Println("a\vc")
}