package main

import (
	"fmt"
)

func main() {
	a := Func{Add{Const{5.23}, Var{}}}
	b := a.derive()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b.latex())
	fmt.Println(a.eval(2))
}
