package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type program struct {
	id        int
	reg       map[string]int
	last      int
	commands  map[int][]string
	output    chan int
	input     chan int
	sendcount int
	mulcount  int
}

func (p program) valorreg(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return p.reg[s]
	}
	return val
}

func (p *program) cmd(parts []string) (int, bool) {
	switch parts[0] {
	case "snd":
		val := p.valorreg(parts[1])
		fmt.Printf("%d Sending: %d\n", p.id, val)
		p.last = val
		p.output <- val
		p.sendcount++
	case "set":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] = val
		fmt.Printf("%d Set %s to %d\n", p.id, parts[1], val)
	case "add":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] += val
		fmt.Printf("%d Increase %s by %d [%d]\n", p.id, parts[1], val, p.reg[parts[1]])
	case "sub":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] -= val
		fmt.Printf("%d Increase %s by %d [%d]\n", p.id, parts[1], val, p.reg[parts[1]])
	case "mul":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] *= val
		fmt.Printf("%d Multiply %s by %d [%d]\n", p.id, parts[1], val, p.reg[parts[1]])
		p.mulcount++
	case "rcv":
		c1 := make(chan bool)
		go func() {
			time.Sleep(time.Second * 1)
			c1 <- true
		}()
		select {
		case rcv := <-p.input:
			fmt.Printf("%d Received %d into %s\n", p.id, rcv, parts[1])
			p.reg[parts[1]] = rcv
		case <-c1:
			fmt.Printf("%d Aborting on wait\n", p.id)
			return 0, true
		}
	case "jgz":
		cond := p.valorreg(parts[1])
		jump := p.valorreg(parts[2])
		if cond > 0 {
			fmt.Printf("%d Jumping by %d\n", p.id, jump)
			return jump, false
		}
	case "jnz":
		cond := p.valorreg(parts[1])
		jump := p.valorreg(parts[2])
		if cond != 0 {
			fmt.Printf("%d Jumping by %d\n", p.id, jump)
			return jump, false
		}
	case "mod":
		regval := p.valorreg(parts[1])
		val := p.valorreg(parts[2])
		p.reg[parts[1]] = regval % val
		fmt.Printf("%d Modulo %s by %d [%d]\n", p.id, parts[1], val, p.reg[parts[1]])

	}
	return 1, false
}

func (p *program) runprogram(progreader io.Reader, id int, output chan int, input chan int) {
	p.reg["p"] = id
	p.id = id
	p.commands = make(map[int][]string)
	p.output = output
	p.input = input
	scanner := bufio.NewScanner(progreader)
	var i int
	for scanner.Scan() {
		p.commands[i] = strings.Split(scanner.Text(), " ")
		i++
	}
	var length = i
	fmt.Printf("%d Loaded %d instructions\n", id, length)
	pc := 0
	for {
		if pc < 0 || pc >= length {
			fmt.Printf("%d Stepping out of program, stop\n", p.id)
			break
		}
		jump, stop := p.cmd(p.commands[pc])
		if stop == true {
			return
		}
		pc += jump
	}
}

func main() {
	//input0, err := os.Open("input")
	//if err != nil {
	//	panic(err)
	//}
	input1, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	//p0 := program{reg: make(map[string]int)}
	p1 := program{reg: make(map[string]int)}
	fromzero := make(chan int, 1000)
	fromone := make(chan int, 1000)
	//go p0.runprogram(input0, 0, fromzero, fromone)
	p1.reg["a"] = 0
	p1.runprogram(input1, 1, fromone, fromzero)
	fmt.Println(p1.mulcount)
}
