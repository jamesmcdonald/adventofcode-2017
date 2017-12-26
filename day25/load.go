package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type loadstate int

const (
	begin loadstate = iota
	checksum
	blank
	getstate
	current
	write
	move
	new
)

func (m *Machine) load(input io.Reader) {
	scanner := bufio.NewScanner(input)
	ls := begin

	var sc [2]statechange
	var st rune
	cond := 1

	for scanner.Scan() {
		text := scanner.Text()
		switch ls {
		case begin:
			m.start = rune(text[len(text)-2])
			ls = checksum
		case checksum:
			parts := strings.Split(text, " ")
			var err error
			m.steps, err = strconv.Atoi(parts[5])
			if err != nil {
				panic(err)
			}
			ls = blank
		case blank:
			ls = getstate
		case getstate:
			st = rune(text[len(text)-2])
			ls = current
		case current:
			if cond == 1 {
				cond = 0
			} else {
				cond = 1
			}
			ls = write
		case write:
			sc[cond].value = int(text[len(text)-2] - '0')
			ls = move
		case move:
			if text[len(text)-3] == 'h' {
				sc[cond].dir = right
			} else {
				sc[cond].dir = left
			}
			ls = new
		case new:
			sc[cond].next = rune(text[len(text)-2])
			if cond == 1 {
				m.stategen(st, sc)
				ls = blank
			} else {
				ls = current
			}
		}
	}
}
