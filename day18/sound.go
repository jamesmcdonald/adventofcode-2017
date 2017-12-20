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
	reg       map[string]int64
	last      int64
	commands  map[int64][]string
	output    chan int64
	input     chan int64
	sendcount int
}

func (p program) valorreg(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return p.reg[s]
	}
	return val
}

func (p *program) cmd(parts []string) (int64, bool) {
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
	case "mul":
		val := p.valorreg(parts[2])
		p.reg[parts[1]] *= val
		fmt.Printf("%d Multiply %s by %d [%d]\n", p.id, parts[1], val, p.reg[parts[1]])
	case "rcv":
		val := p.valorreg(parts[1])
		if val != 0 {
			c1 := make(chan bool)
			go func() {
				time.Sleep(time.Second * 1)
				c1 <- true
			}()
			select {
			case rcv := <-p.input:
				fmt.Printf("%d Received last frequency: %d\n", p.id, rcv)
			case <-c1:
				fmt.Printf("%d Aborting on wait\n", p.id)
				return 0, true
			}
		}
	case "jgz":
		cond := p.valorreg(parts[1])
		jump := p.valorreg(parts[2])
		if cond > 0 {
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

func (p *program) runprogram(progreader io.Reader, id int, output chan int64, input chan int64) {
	p.id = id
	p.commands = make(map[int64][]string)
	p.output = output
	p.input = input
	scanner := bufio.NewScanner(progreader)
	var i int64
	for scanner.Scan() {
		p.commands[i] = strings.Split(scanner.Text(), " ")
		i++
	}
	var length = i
	fmt.Printf("%d Loaded %d instructions\n", id, length)
	pc := int64(0)
	for {
		if pc < 0 || pc >= length {
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
	input0, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	input1, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	p0 := program{reg: make(map[string]int64)}
	p1 := program{reg: make(map[string]int64)}
	fromzero := make(chan int64, 1000)
	fromone := make(chan int64, 1000)
	go p0.runprogram(input0, 0, fromzero, fromone)
	p1.runprogram(input1, 1, fromone, fromzero)
	fmt.Println(p0.sendcount, p1.sendcount)
}
