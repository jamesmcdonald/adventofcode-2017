package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rowchecksum1(row []int) int {
	max := row[0]
	min := row[0]
	for _, i := range row[1:] {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	return max - min
}

func rowchecksum2(row []int) int {
	for i, a := range row {
		for _, b := range row[i+1:] {
			if a%b == 0 {
				return a / b
			} else if b%a == 0 {
				return b / a
			}
		}
	}
	panic(row)
}

func checksum(rowcs func(row []int) int) int {
	sum := 0
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var numbers []int
		parts := strings.Split(scanner.Text(), "\t")
		for _, p := range parts {
			val, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, val)
		}
		sum += rowcs(numbers)
	}
	return sum
}

func main() {
	fmt.Println(checksum(rowchecksum1))
	fmt.Println(checksum(rowchecksum2))
}
