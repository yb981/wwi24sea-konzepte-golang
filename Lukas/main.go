// main.go
// Main function to start the calculator programm.
//
// Author: Lukas Gröning
// Date: 22.02.2025

package main

import (
	"fmt"
)

func main() {
	calculator := calculator{}

	calculator.printWelcomeMessage()

	for {
		calculator.numberStack.Print()

		input := getInput()

		calculator.checkInput(input)
	}
}

func getInput() string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Eingabefehler")
	}
	return input
}