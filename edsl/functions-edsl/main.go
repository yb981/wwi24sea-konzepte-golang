package main

import (
	"fmt"
)

func main() {
	a := Func{Add{Const{10}, Sqr{Var{}}}}
	b := a.derive()
	c := b.derive()
	fmt.Print(a, "\n")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(b.latex())
}
