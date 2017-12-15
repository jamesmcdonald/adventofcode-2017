package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jamesmcdonald/adventofcode-2017/day10/knot"
)

func main() {
	instring, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input1, err := knot.ParseInput(string(instring))
	if err != nil {
		panic(err)
	}
	input2 := knot.ParseASCII(string(instring))
	result1 := knot.Knot(input1, 1)
	fmt.Println(result1[0] * result1[1])

	result2 := knot.Knot(input2, 64)
	knot.PrintHash(knot.Hash(result2))
}
