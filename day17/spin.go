package main

import "fmt"

const step int = 337

type bufType struct {
	buffer []int
	pos    int
}

func (b *bufType) add(n int) {
	b.pos = (b.pos + step) % len(b.buffer)
	b.buffer = append(b.buffer, 0)
	b.pos++
	copy(b.buffer[b.pos+1:], b.buffer[b.pos:])
	b.buffer[b.pos] = n
}

func followzero(n int) int {
	afterz := 0
	pos := 0
	for i := 1; i <= n; i++ {
		pos = (pos + step) % i
		if pos == 0 {
			afterz = i
		}
		pos++
	}
	return afterz
}

func main() {
	b := bufType{buffer: []int{0}}
	for i := 1; i <= 2017; i++ {
		b.add(i)
	}
	fmt.Println(b.buffer[b.pos+1])
	b = bufType{buffer: []int{0}}
	fmt.Println(followzero(50000000))
}
