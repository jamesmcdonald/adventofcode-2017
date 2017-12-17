package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type programlist []rune

func (p *programlist) spin(n int) {
	size := len(*p)
	n = size - n
	*p = programlist(append((*p)[n:], (*p)[:n]...))
}

func (p *programlist) exchange(a int, b int) {
	(*p)[a], (*p)[b] = (*p)[b], (*p)[a]
}

func (p *programlist) swap(a rune, b rune) {
	apos := -1
	bpos := -1
	for i, r := range *p {
		switch r {
		case a:
			if bpos >= 0 {
				(*p)[i], (*p)[bpos] = (*p)[bpos], (*p)[i]
				break
			}
			apos = i
		case b:
			if apos >= 0 {
				(*p)[i], (*p)[apos] = (*p)[apos], (*p)[i]
				break
			}
			bpos = i
		}
	}
}

func (p programlist) String() string {
	return string(p)
}

func atoi(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func (p *programlist) dance(steps []string) {
	for _, step := range steps {
		switch step[0] {
		case 's':
			p.spin(atoi(step[1:]))
		case 'x':
			nums := strings.Split(step[1:], "/")
			p.exchange(atoi(nums[0]), atoi(nums[1]))
		case 'p':
			chars := strings.Split(step[1:], "/")
			p.swap(rune(chars[0][0]), rune(chars[1][0]))
		default:
			panic(step)
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	steps := strings.Split(strings.TrimSpace(string(input)), ",")
	p := programlist("abcdefghijklmnop")
	p.dance(steps)
	fmt.Println(p)
	p = programlist("abcdefghijklmnop")
	for i := 1; i < 1000000000%30+1; i++ {
		p.dance(steps)
	}
	fmt.Println(p)
}
