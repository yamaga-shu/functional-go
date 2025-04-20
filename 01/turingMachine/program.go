package turingmachine

// state represents the machine's state
type state int

const (
	q0 state = iota
	q1
	q2
	q3
	q4
)

// move represents the next position of the machine, either backward or forward
type move int

const (
	backward move = iota
	forward
)

// information represents the available information of this Turing machine
type information string

const (
	zero  information = "0"
	one   information = "1"
	blank information = "B"
)

// instruction instructs the machine on how to behave based on the current head information
type instruction struct {
	write information // a value to be written on the tape
	move  move        // transition of the head
	next  state       // next state
}

type program map[state]map[information]instruction

// AddOne is passed to the Turing machine.
// This program defines the logic to add 1 to a given binary number.
var AddOne = program{
	q0: {
		one:   {write: one, move: forward, next: q0},
		zero:  {write: zero, move: forward, next: q0},
		blank: {write: blank, move: backward, next: q1},
	},
	q1: {
		one:   {write: zero, move: backward, next: q1},
		zero:  {write: one, move: backward, next: q2},
		blank: {write: one, move: backward, next: q3},
	},
	q2: {
		one:   {write: one, move: backward, next: q2},
		zero:  {write: zero, move: backward, next: q2},
		blank: {write: blank, move: forward, next: q4},
	},
	q3: {
		one:   {write: one, move: forward, next: q4},
		zero:  {write: zero, move: forward, next: q4},
		blank: {write: blank, move: forward, next: q4},
	},
}

var InitState state = q0
var EndState state = q4
