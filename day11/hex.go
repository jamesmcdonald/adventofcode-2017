package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type hexloc struct {
	x        int
	y        int
	furthest int
}

type direction int

const (
	n direction = iota
	s
	ne
	nw
	se
	sw
)

var dirname = map[string]direction{
	"n":  n,
	"s":  s,
	"ne": ne,
	"nw": nw,
	"se": se,
	"sw": sw,
}

func (h *hexloc) move(d direction) {
	switch d {
	case n:
		h.y += 2
	case s:
		h.y -= 2
	case ne:
		h.x++
		h.y++
	case nw:
		h.x--
		h.y++
	case se:
		h.x++
		h.y--
	case sw:
		h.x--
		h.y--
	}
	distance := h.distance()
	if distance > h.furthest {
		h.furthest = distance
	}
}

func (h hexloc) distance() int {
	result := 0
	if h.x < 0 {
		h.x = -h.x
	}
	if h.y < 0 {
		h.y = -h.y
	}
	result += h.x
	h.y -= h.x
	result += h.y / 2
	return result
}

func (h *hexloc) walk(path []direction) {
	for _, d := range path {
		h.move(d)
	}
}

func main() {
	instring, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	instrings := strings.Split(strings.TrimSpace(string(instring)), ",")
	input := make([]direction, len(instrings))
	for i, s := range instrings {
		input[i] = dirname[s]
	}
	var h hexloc
	h.walk(input)
	fmt.Println(h.distance(), h.furthest)
}
