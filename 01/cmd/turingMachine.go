package main

import (
	"fmt"

	turingmachine "github.com/functional-go/01/turingMachine"
)

func main() {
	tm := turingmachine.TuringMachine{}

	result := tm.Calculate(
		turingmachine.Program,
		turingmachine.Tape,
		turingmachine.InitState,
		turingmachine.EndState,
	)

	fmt.Println(result)
	// Output:
	// map[-1:B 0:1 1:1 2:B]
}
