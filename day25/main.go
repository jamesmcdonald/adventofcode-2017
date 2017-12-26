package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	m := NewMachine()
	m.load(input)
	fmt.Println(m.run())
}
