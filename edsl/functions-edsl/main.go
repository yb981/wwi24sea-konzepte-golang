package main

import(
	"fmt"
)

func main(){
	a := Sub{Add{Var{}, Const{3}}, Mult{Const{5}, Const{4}}}
	b := Div{Const{5}, Const{2}}
	fmt.Println(a.eval(10))
	fmt.Println(b.eval(0))
}