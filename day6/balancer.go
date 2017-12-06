package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type block int

type memory []block

func loadmem(r io.Reader) memory {
	var m memory
	s := bufio.NewScanner(r)
	for s.Scan() {
		for _, val := range strings.Split(s.Text(), "\t") {
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			m = append(m, block(v))
		}
	}
	return m
}

func (m memory) String() string {
	conv := make([]string, len(m))
	for i, v := range m {
		conv[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(conv, ",")
}

func (m *memory) balance() {
	size := len(*m)
	var blocks block
	blocks, rbindex := 0, 0

	for i, v := range *m {
		if v > blocks {
			blocks = v
			rbindex = i
		}
	}

	(*m)[rbindex] = 0
	for i := (rbindex + 1) % size; blocks > 0; i = (i + 1) % size {
		(*m)[i]++
		blocks--
	}

}

type history map[string]memory

func (h *history) load(m memory) {
	if *h == nil {
		*h = make(history)
	}
	(*h)[m.String()] = m
}

// Run the balancer in a loop until it repeats
// Return the number of steps
func (m *memory) runbalancer() (count int) {
	h := make(history)
	for h[m.String()] == nil {
		h.load(*m)
		//fmt.Println(count, m)
		m.balance()
		count++
	}
	return
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	m := loadmem(input)
	firstcount := m.runbalancer()
	repeatcount := m.runbalancer()

	fmt.Println("Repeat detected after", firstcount, "iterations")
	fmt.Println("Reiterated after", repeatcount, "iterations")
}
