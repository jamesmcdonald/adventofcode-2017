package cpu

import (
	"bufio"
	"fmt"
	"io"
)

var ops map[string]op = map[string]op{
	"inc": inc,
	"dec": dec,
}
var compareops map[string]compare = map[string]compare{
	"==": eq,
	"!=": ne,
	"<":  lt,
	">":  gt,
	"<=": le,
	">=": ge,
}

// Parser parses lines of input and emits instructions
func Parser(r io.Reader, output chan<- instruction) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		var reg, op, creg, cop string
		var val, cval int
		_, err := fmt.Sscanf(text, "%s %s %d if %s %s %d", &reg, &op, &val, &creg, &cop, &cval)
		if err != nil {
			panic(err)
		}
		c := condition{
			register: creg,
			value:    cval,
			op:       compareops[cop],
		}
		i := instruction{
			register:  reg,
			value:     val,
			op:        ops[op],
			condition: c,
		}

		output <- i
	}
	close(output)
}

// LoadFromReader loads from a reader into the CPU
func (c *CPU) LoadFromReader(r io.Reader) {
	// PRESS PLAY ON TAPE
	tape := make(chan instruction)
	go Parser(r, tape)
	c.load(tape)
}
