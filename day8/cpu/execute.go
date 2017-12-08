package cpu

func (c *CPU) evaluateCondition(cond condition) bool {
	switch cond.op {
	case eq:
		return c.Registers[cond.register] == cond.value
	case ne:
		return c.Registers[cond.register] != cond.value
	case gt:
		return c.Registers[cond.register] > cond.value
	case lt:
		return c.Registers[cond.register] < cond.value
	case ge:
		return c.Registers[cond.register] >= cond.value
	case le:
		return c.Registers[cond.register] <= cond.value
	default:
		panic(cond.op)
	}
}

// Execute a single instruction
func (c *CPU) Execute(i instruction) int {
	result := c.evaluateCondition(i.condition)
	if !result {
		return 0
	}
	switch i.op {
	case inc:
		c.Registers[i.register] += i.value
	case dec:
		c.Registers[i.register] -= i.value
	default:
		panic(i.op)
	}
	return c.Registers[i.register]
}

// Run the program loaded into the CPU returning the max register value
func (c *CPU) Run() int {
	max := 0
	for _, i := range c.Instructions {
		reg := c.Execute(i)
		if reg > max {
			max = reg
		}
	}
	return max
}
