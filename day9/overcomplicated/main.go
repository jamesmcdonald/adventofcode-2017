package main

import (
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	parse(string(input))
}
