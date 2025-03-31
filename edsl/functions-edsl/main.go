package main

import (
	"fmt"
)

func main() {
	a := Mult{Var{}, Var{}}
	fmt.Println(a.latex())
	fmt.Println(a.derive().latex())
}
