package main

import(
	"fmt"
)

func main(){
	a := Sub{Add{Pow{Var{}, Const{2}}, Mult{Const{2}, Var{}}}, Const{3}}
	fmt.Println(a.eval(5), a)
}