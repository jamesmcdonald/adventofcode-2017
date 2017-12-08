package cpu

type registers map[string]int

// CPU represents a whole CPU with state and instructions
type CPU struct {
	Registers    registers
	Instructions []instruction
}

// NewCPU allocated a new CPU and returns it
func NewCPU() CPU {
	return CPU{
		Registers: make(registers),
	}
}

type op int

const (
	dec op = iota
	inc
)

type compare int

const (
	eq compare = iota
	ne
	lt
	gt
	le
	ge
)

type condition struct {
	register string
	op       compare
	value    int
}

type instruction struct {
	register  string
	op        op
	value     int
	condition condition
}
