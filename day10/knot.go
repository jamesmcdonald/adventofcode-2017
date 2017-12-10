package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseinput(instring string) ([]int, error) {
	instrings := strings.Split(strings.TrimSpace(instring), ",")
	input := make([]int, len(instrings))
	for i, s := range instrings {
		var err error
		input[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}
	return input, nil
}

func parseascii(instring string) []int {
	instring = strings.TrimSpace(instring)
	input := make([]int, len(instring))
	for i, c := range instring {
		input[i] = int(c)
	}
	for _, val := range []int{17, 31, 73, 47, 23} {
		input = append(input, val)
	}
	return input
}

func knot(sequence []int, rounds int) [256]int {
	var ring [256]int
	for i := range ring {
		ring[i] = i
	}

	skip := 0
	pos := 0
	for r := 0; r < rounds; r++ {
		for _, s := range sequence {
			for i, j := 0, s-1; i < j; i, j = i+1, j-1 {
				ring[(i+pos)%256], ring[(j+pos)%256] =
					ring[(j+pos)%256], ring[(i+pos)%256]
			}
			pos += s + skip
			pos = pos % 256
			skip++
		}
	}

	return ring
}

func hash(ring [256]int) [16]int {
	var result [16]int
	for i := 0; i < 16; i++ {
		result[i] = ring[i*16]
		for j := 1; j < 16; j++ {
			result[i] ^= ring[i*16+j]
		}
	}
	return result
}

func printhash(h [16]int) {
	for _, v := range h {
		fmt.Printf("%02x", v)
	}
	fmt.Println()
}

func main() {
	instring, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input1, err := parseinput(string(instring))
	if err != nil {
		panic(err)
	}
	input2 := parseascii(string(instring))
	result1 := knot(input1, 1)
	fmt.Println(result1[0] * result1[1])

	result2 := knot(input2, 64)
	printhash(hash(result2))
}
