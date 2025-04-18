package turingmachine

// state represents machine's state
type state int

const (
	q0 state = iota
	q1
	q2
	q3
	q4
)

// move represents next position of machine, only backward or forward
type move int

const (
	backward move = iota
	forward
)

// information represents available information of this turing-machine
type information string

const (
	zero information = "0"
	one  information = "1"
	back information = "B"
)

// impertrative would be passed to turing-machine
type imperative struct {
	write information
	move  move
	next  state
}
