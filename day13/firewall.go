package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rules map[int]int

func severity(rules map[int]int, delay int) int {
	severity := 0
	for fdepth, frange := range rules {
		fdepth += delay
		if fdepth == 1 {
			if frange == 1 {
				severity++
			}
		} else if fdepth%((frange-1)*2) == 0 {
			severity += fdepth * frange
		}
	}
	return severity
}

func main() {
	rules := make(map[int]int)
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		if len(parts) != 2 {
			panic(parts)
		}
		fdepth, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(parts)
		}
		frange, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(parts)
		}
		rules[fdepth] = frange
	}

	fmt.Println(severity(rules, 0))
	for d := 0; ; d++ {
		if severity(rules, d) == 0 {
			fmt.Println("Delay", d)
			break
		}
	}
}
