package turingmachine

type tape map[int]information

// Two represents "2" in the binary system
var Two tape = tape{
	0: one,
	1: zero,
}
