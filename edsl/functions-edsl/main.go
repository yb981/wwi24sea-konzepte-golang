package main

import(
	"fmt"
)

func main(){
	a := Func{Add{Mult{Var{}, Const{6}}, Var{}}}
	b := a.derive()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b.latex())
	fmt.Println(a.eval(2))
}