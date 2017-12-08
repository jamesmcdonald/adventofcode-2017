package cpu

func (cpu *CPU) load(input <-chan instruction) {
	for i := range input {
		cpu.Instructions = append(cpu.Instructions, i)
	}
}
