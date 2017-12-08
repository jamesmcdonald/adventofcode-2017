package main

import (
	"fmt"
	"github.com/jamesmcdonald/adventofcode-2017/day8/cpu"
	"os"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	c := cpu.NewCPU()
	c.LoadFromReader(input)
	max := c.Run()
	topreg := 0
	for _, reg := range c.Registers {
		if reg > topreg {
			topreg = reg
		}
	}
	fmt.Println(topreg, max)
}
