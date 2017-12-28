package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type direction int

const (
	up    direction = 0
	right           = 1
	down            = 2
	left            = 3
)

func (d direction) turnleft() direction {
	return (d + 4 - 1) % 4
}

func (d direction) turnright() direction {
	return (d + 1) % 4
}

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}
	return "extradimensional hypertravel"
}

type node struct {
	infected bool
	start    bool
}

type loc struct {
	x int
	y int
}

type memory struct {
	nodes      [][]node
	loc        loc
	vdir       direction
	infections int
}

func loadmem(input []string) memory {
	n := make([][]node, len(input))
	for i := range n {
		n[i] = make([]node, len(input))
	}
	for i, row := range input {
		for j, char := range row {
			if char == '#' {
				n[i][j].infected = true
				n[i][j].start = true
			}
		}
	}
	l := loc{len(input) / 2, len(input) / 2}
	v := up
	return memory{n, l, v, 0}
}

func (m *memory) extendnodes() {
	oldsize := len(m.nodes)
	newnodes := make([][]node, 3*oldsize)
	for i := range newnodes {
		newnodes[i] = make([]node, 3*oldsize)
	}
	for i, row := range m.nodes {
		for j, thenode := range row {
			newnodes[i+oldsize][j+oldsize] = thenode
		}
	}
	m.loc.x += oldsize
	m.loc.y += oldsize
	m.nodes = newnodes
}

func (m *memory) move(d direction) {
	//	fmt.Println("Moving", d)
	switch d {
	case up:
		if m.loc.y == 0 {
			m.extendnodes()
		}
		m.loc.y--
	case down:
		if m.loc.y == len(m.nodes)-1 {
			m.extendnodes()
		}
		m.loc.y++
	case left:
		if m.loc.x == 0 {
			m.extendnodes()
		}
		m.loc.x--
	case right:
		if m.loc.x == len(m.nodes)-1 {
			m.extendnodes()
		}
		m.loc.x++
	}
}

func (m memory) String() string {
	var b bytes.Buffer
	for i, row := range m.nodes {
		for j, node := range row {
			var c rune
			if node.infected {
				c = '#'
			} else {
				c = '.'
			}
			if i == m.loc.y && j == m.loc.x {
				fmt.Fprintf(&b, "[%c]", c)
			} else {
				fmt.Fprintf(&b, " %c ", c)
			}
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func (m *memory) virophage() {
	fmt.Println("Facing", m.vdir)
	if m.nodes[m.loc.y][m.loc.x].infected {
		m.vdir = m.vdir.turnright()
		m.nodes[m.loc.y][m.loc.x].infected = false
		fmt.Println("Turn right and CLEAN")
	} else {
		m.vdir = m.vdir.turnleft()
		m.nodes[m.loc.y][m.loc.x].infected = true
		fmt.Println("Turn left and INFECT")
		m.infections++
	}
	fmt.Println("Move", m.vdir)
	m.move(m.vdir)
}

func main() {
	inbytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(inbytes)), "\n")
	m := loadmem(input)
	fmt.Println(m)
	for i := 0; i < 10000; i++ {
		//		time.Sleep(time.Millisecond * 10)
		m.virophage()
	}
	fmt.Println(m.infections)
}
