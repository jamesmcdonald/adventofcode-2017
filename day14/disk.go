package main

import (
	"fmt"

	"github.com/jamesmcdonald/adventofcode-2017/day10/knot"
)

const input string = "ugkiagan"

func countHashBits(hash [16]int) int {
	bits := 0
	for _, i := range hash {
		for j := i; j != 0; j >>= 1 {
			bits += j & 1
		}
	}
	return bits
}

func main() {
	bits := 0
	for i := 0; i < 128; i++ {
		str := fmt.Sprintf("%s-%d", input, i)
		hash := knot.Hash(knot.Knot(knot.ParseASCII(str), 64))
		bits += countHashBits(hash)
	}
	fmt.Println(bits)
}
