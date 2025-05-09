package main

import (
	"fmt"

	. "github.com/yb981/wwi24sea-konzepte-golang/edsl/functions-edsl"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("-------------------------------------- LIVE DEMO SHOWCASE -----------------------------------\n\n")
	fmt.Printf("----------------------------------------- 1st FUNCTION --------------------------------------\n\n")
	f := Func{Mult{Const{2}, Var{}}}
	fmt.Printf("First Function: f(x) = 2 * x\n\n")

	fmt.Printf("f(9) = %v\n", f.Eval(9))
	fmt.Println("Latex Output: ", f.Latex())
	fmt.Printf("Derivative Function: %v\n\n", f.Derive())

	fmt.Printf("----------------------------------------- 2nd FUNCTION --------------------------------------\n\n")

	f = Func{Add{Var{}, Mult{Var{}, Var{}}}}
	fmt.Printf("First Function: f(x) = x + x * x\n\n")

	fmt.Printf("f(9) = %v\n", f.Eval(9))
	fmt.Println("Latex Output: ", f.Latex())
	fmt.Printf("Derivative Function: %v\n", f.Derive())
	fmt.Printf("Value of derivative Function f(9): %v\n\n", f.Derive().Eval(9))

	fmt.Printf("----------------------------------------- 3rd FUNCTION --------------------------------------\n\n")
	f = Func{Pow{Var{}, Const{10}}}
	fmt.Printf("First Function: f(x) = x ^ 10\n\n")

	fmt.Printf("f(10) = %v\n", f.Eval(10))
	fmt.Println("Latex Output: ", f.Latex())
	fmt.Printf("1st derivative Function: %v\n", f.Derive())
	fmt.Printf("2nd derivative Function: %v\n", f.Derive().Derive())
	fmt.Printf("3rd derivative Function: %v\n", f.Derive().Derive().Derive())
	fmt.Printf("Value of 3rd derivative: f(1) = %v\n\n", f.Derive().Derive().Derive().Eval(1))

	fmt.Printf("----------------------------------------- 4th FUNCTION --------------------------------------\n\n")
	f = Func{Div{Var{}, Add{Const{10}, Const{20}}}}
	fmt.Printf("First Function: f(x) =  x / (10 + 20)\n\n")

	fmt.Printf("f(10) = %v\n", f.Eval(30))
	fmt.Println("Latex Output: ", f.Latex())
	fmt.Printf("1st derivative Function: %v\n", f.Derive())
	fmt.Printf("2nd derivative Function: %v\n", f.Derive().Derive())
	fmt.Printf("Value of 1st derivative: f(1000000) = %v\n\n", f.Derive().Eval(1000000))

	fmt.Printf("----------------------------------------- 5th FUNCTION --------------------------------------\n\n")
	f = Func{Div{Var{}, Var{}}}
	fmt.Printf("First Function: f(x) =  x / x\n\n")

	fmt.Printf("f(30) = %v\n", f.Eval(30))
	fmt.Println("Latex Output: ", f.Latex())
	fmt.Printf("1st derivative Function: %v\n", f.Derive())
	fmt.Printf("Value of 1st derivative: f(1000000) = %v\n\n", f.Derive().Eval(1000000))

	fmt.Printf("----------------------------------------- 6th FUNCTION --------------------------------------\n\n")
	f = Func{Sqrt{Var{}}}
	fmt.Printf("First Function: f(x) = \\sqrt{x}\n\n")

	fmt.Printf("f(9) = %v\n", f.Eval(9))
	fmt.Println("Latex Output: ", f.Latex())
	fmt.Printf("1st derivative Function: %v\n", f.Derive())
	fmt.Printf("Value of 1st derivative: f(9) = %v\n\n", f.Derive().Eval(9))

	fmt.Printf("----------------------------------------- 7th FUNCTION --------------------------------------\n\n")
	fmt.Println("Live demo! :D")
	c := Func{Div{Var{}, Const{10}}}
	fmt.Println(c.Derive().Latex())
}
