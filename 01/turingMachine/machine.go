package turingmachine

// turingmachine is a model of Turing macine
type TuringMachine struct {
	position int
}

func (tm *TuringMachine) Calculate(p program, t tape, initState state, endState state) tape {
	var state state = initState
	var instr instruction

	// loop while state will be end state
	for state != endState {
		// read current information
		cell, ok := t[tm.position]
		if !ok {
			cell = blank
		}

		// get current instruction from state & cell
		instr = p[state][cell]

		// write to the tape
		t[tm.position] = instr.write
		// move head
		if instr.move == forward {
			tm.position++
		} else if instr.move == backward {
			tm.position--
		}
		// transit next state
		state = instr.next
	}

	return t
}
