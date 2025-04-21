package main

import (
	"fmt"

	turingmachine "github.com/yamaga-shu/functional-go/01/turingMachine"
)

func main() {
	tm := turingmachine.TuringMachine{}

	// result will represent "3" in the binary system
	result := tm.Calculate(
		turingmachine.AddOne,
		turingmachine.Two,
		turingmachine.InitState,
		turingmachine.EndState,
	)

	fmt.Println(result)
	// Output:
	// map[-1:B 0:1 1:1 2:B]
}
