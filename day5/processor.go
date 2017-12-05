package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type processor struct {
	pos    int
	count  int
	memory []int
}

func (p *processor) Append(n int) {
	p.memory = append(p.memory, n)
}

func (p *processor) Load(source <-chan int) {
	for n := range source {
		p.Append(n)
	}
}

func (p *processor) Run() int {
	for p.pos < len(p.memory) {
		p.memory[p.pos]++
		p.pos += p.memory[p.pos] - 1
		p.count++
	}
	return p.count
}

func (p *processor) Run2() int {
	for p.pos < len(p.memory) {
		oldpos := p.pos
		p.pos += p.memory[p.pos]
		offset := p.pos - oldpos
		if offset >= 3 {
			p.memory[oldpos]--
		} else {
			p.memory[oldpos]++
		}
		p.count++
	}
	return p.count
}

func loadintlines(r io.Reader, dest chan<- int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		dest <- val
	}
	close(dest)
}

func tee(source <-chan int, dest1 chan<- int, dest2 chan<- int) {
	for val := range source {
		dest1 <- val
		dest2 <- val
	}
	close(dest1)
	close(dest2)
}

func main() {
	var p processor
	var p2 processor
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := make(chan int)
	loader1 := make(chan int)
	loader2 := make(chan int)

	go p.Load(loader1)
	go p2.Load(loader2)

	go loadintlines(input, reader)
	tee(reader, loader1, loader2)

	fmt.Println(p.Run())
	fmt.Println(p2.Run2())
}
