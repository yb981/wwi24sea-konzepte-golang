package main

import(
	"fmt"
)

func main(){
	a := Func{Sub{Const{10}, Var{}}}
	fmt.Println(a.latex())
	b := a.derive()
	fmt.Println(b.latex())
	fmt.Println(b.eval(9))
}