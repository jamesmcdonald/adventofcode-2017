package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	garbage := false
	skip := false
	depth := 0
	sum := 0
	gc := 0
	for _, c := range input {
		if skip {
			skip = false
			continue
		}
		if garbage {
			if c == '!' {
				skip = true
			} else if c == '>' {
				garbage = false
			} else {
				gc++
			}
			continue
		}
		switch c {
		case '{':
			depth++
			sum += depth
		case '}':
			depth--
		case '<':
			garbage = true
		}
	}
	fmt.Println(depth, sum, gc)
}
