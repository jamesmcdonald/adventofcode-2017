package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type program struct {
	reg      map[string]int64
	last     int64
	commands map[int64][]string
}

func (p program) valorreg(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return reg[s]
	}
	return val
}

func (p program) cmd(parts []string) int64 {
	switch parts[0] {
	case "snd":
		val := p.valorreg(parts[1])
		fmt.Printf("Playing sound: %d\n", val)
		p.last = val
	case "set":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] = val
		fmt.Printf("Set %s to %d\n", parts[1], val)
	case "add":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] += val
		fmt.Printf("Increase %s by %d [%d]\n", parts[1], val, p.reg[parts[1]])
	case "mul":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] *= val
		fmt.Printf("Multiply %s by %d [%d]\n", parts[1], val, p.reg[parts[1]])
	case "rcv":
		val := p.valorreg(parts[1])
		if val != 0 {
			fmt.Printf("Received last frequency: %d\n", p.last)
		}
	case "jgz":
		cond := p.valorreg(parts[1])
		jump := p.valorreg(parts[2])
		if cond > 0 {
			fmt.Printf("Jumping by %d\n", jump)
			return jump
		}
	case "mod":
		regval := p.valorreg(parts[1])
		val := p.valorreg(parts[2])
		p.reg[parts[1]] = regval % val
		fmt.Printf("Modulo %s by %d [%d]\n", parts[1], val, p.reg[parts[1]])

	}
	return 1
}

func (p *program) runprogram(progreader io.Reader, id int, output chan int64, input chan int64) {
	program := make(map[int64][]string)
	scanner := bufio.NewScanner(progreader)
	var i int64
	for scanner.Scan() {
		program[i] = strings.Split(scanner.Text(), " ")
		i++
	}
	var length = i
	fmt.Printf("%d Loaded %d instructions\n", id, length)
	pc := int64(0)
	for {
		if pc < 0 || pc >= length {
			break
		}
		jump := cmd(program[pc])
		pc += jump
	}

}

func main() {
	reg = make(map[string]int64)
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fromzero := make(chan int64)
	fromone := make(chan int64)
	go runprogram(input, 0, fromzero, fromone)
	runprogram(input, 1, fromone, fromzero)
}
