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
	zero information = "0"
	one  information = "1"
	back information = "B"
)

// imperative instructs the machine on how to behave based on the current head information
type imperative struct {
	write information // a value to be written on the tape
	move  move        // transition of the head
	next  state       // next state
}

type program map[state]map[information]imperative

// programs is passed to the Turing machine
var programs = program{
	q0: {
		one:  {write: one, move: forward, next: q0},
		zero: {write: zero, move: forward, next: q0},
		back: {write: back, move: backward, next: q1},
	},
	q1: {
		one:  {write: one, move: forward, next: q2},
		zero: {write: zero, move: backward, next: q0},
		back: {write: back, move: backward, next: q3},
	},
	q2: {
		one:  {write: one, move: backward, next: q2},
		zero: {write: zero, move: backward, next: q2},
		back: {write: back, move: forward, next: q4},
	},
	q3: {
		one:  {write: one, move: forward, next: q4},
		zero: {write: zero, move: forward, next: q4},
		back: {write: back, move: forward, next: q4},
	},
}
