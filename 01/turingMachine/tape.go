package turingmachine

type tape map[int]information

var Tape tape = tape{
	0: one,
	1: zero,
}
