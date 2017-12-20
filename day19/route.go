package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type vector struct {
	name string
	x    int
	y    int
}

func (v vector) add(a vector) vector {
	return vector{v.name, v.x + a.x, v.y + a.y}
}

var (
	north = vector{"north", 0, -1}
	east  = vector{"east", 1, 0}
	south = vector{"south", 0, 1}
	west  = vector{"west", -1, 0}
)

func walk(grid []string, sx int) string {
	var result string
	loc := vector{"loc", sx, 0}
	dir := south
	steps := 0
	for stop := false; !stop; {
		switch grid[loc.y][loc.x] {
		case '|', '-':
			loc = loc.add(dir)
			steps++
		case '+':
			switch dir {
			case north, south:
				eloc := loc.add(east)
				switch grid[eloc.y][eloc.x] {
				case ' ', '\n':
					dir = west
				default:
					dir = east
				}
				loc = loc.add(dir)
			case east, west:
				eloc := loc.add(north)
				switch grid[eloc.y][eloc.x] {
				case ' ', '\n':
					dir = south
				default:
					dir = north
				}
				loc = loc.add(dir)
			}
			steps++
		case ' ', '\n':
			stop = true
			fmt.Println()
			fmt.Println(steps)
		default:
			fmt.Printf("%c", grid[loc.y][loc.x])
			loc = loc.add(dir)
			steps++
		}
	}
	return result
}

func main() {
	ibytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(ibytes), "\n")
	for i, c := range input[0] {
		if c == '|' {
			walk(input, i)
			break
		}
	}
}
