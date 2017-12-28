package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bridge [][2]int

type componentStore map[int][]int

func (br bridge) inbridge(a int, b int) bool {
	for _, c := range br {
		if c[0] == a && c[1] == b || c[1] == a && c[0] == b {
			return true
		}
	}
	return false
}

func (br bridge) strength() int {
	str := 0
	for _, v := range br {
		str += v[0] + v[1]
	}
	return str
}

func (cs componentStore) bridgeGenerator(br bridge, output chan<- bridge) {
	last := br[len(br)-1][1]
	//fmt.Println(br, "last is", last)
	for _, c := range cs[last] {
		if !br.inbridge(last, c) {
			sub := make(bridge, len(br)+1)
			copy(sub, br)
			sub[len(sub)-1] = [2]int{last, c}
			//fmt.Printf("matched %d for last %d in %v\n", c, last, br)
			output <- sub
			cs.bridgeGenerator(sub, output)
		}
	}
}

func (cs componentStore) generateBridges(output chan<- bridge) {
	b := bridge{[2]int{0, 0}}
	cs.bridgeGenerator(b, output)
	close(output)
}

func parseInput(input []string) componentStore {
	cs := make(componentStore)
	for _, s := range input {
		parts := strings.Split(s, "/")
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		cs[a] = append(cs[a], b)
		cs[b] = append(cs[b], a)
	}
	return cs
}

func main() {
	inbytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(inbytes)), "\n")
	cs := parseInput(input)
	fmt.Println(cs)
	bc := make(chan bridge)
	go cs.generateBridges(bc)
	max := 0
	maxlen := 0
	for br := range bc {
		strength := br.strength()
		if len(br) > maxlen {
			maxlen = len(br)
			max = strength
		} else if len(br) == maxlen && strength > max {
			fmt.Println("Bump max to", strength, "at maxlen", maxlen, br)
			max = strength
		}
	}
	fmt.Println(max)
}
