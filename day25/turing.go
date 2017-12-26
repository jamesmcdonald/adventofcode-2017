package main

type tapeType map[int]int

// Machine represents a Turing Machine
type Machine struct {
	start  rune
	steps  int
	cursor int
	tape   tapeType
	sm     statemachine
}

type direction int

func (d direction) String() string {
	switch d {
	case left:
		return "left"
	case right:
		return "right"
	}
	return "STRANGE"
}

const (
	left  direction = -1
	right direction = 1
)

// NewMachine inits a machine
func NewMachine() Machine {
	var m Machine
	m.tape = make(tapeType)
	m.sm = make(statemachine)
	return m
}

func (m *Machine) move(dir direction) {
	m.cursor += int(dir)
}

func (m *Machine) write(i int) {
	m.tape[m.cursor] = i
}

func (m *Machine) read() int {
	return m.tape[m.cursor]
}

func (m Machine) checksum() int {
	checksum := 0
	for _, v := range m.tape {
		if v == 1 {
			checksum++
		}
	}
	return checksum
}

type stateFn func() rune

type statemachine map[rune]stateFn

type statechange struct {
	value int
	dir   direction
	next  rune
}

func (m *Machine) stategen(r rune, states [2]statechange) {
	m.sm[r] = func() rune {
		val := m.read()
		m.write(states[val].value)
		m.move(states[val].dir)
		//		fmt.Printf("%c: r %d w %d m %s p %d n %c\n", r, val,
		//			states[val].value, states[val].dir, m.cursor, states[val].next)
		return states[val].next
	}
}

func (m *Machine) run() int {
	state := m.start
	for i := 0; i < m.steps; i++ {
		//fmt.Printf("%08d ", i)
		state = m.sm[state]()
	}
	return m.checksum()
}
