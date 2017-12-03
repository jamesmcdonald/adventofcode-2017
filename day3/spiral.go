package main

import (
	"fmt"
)

type direction int

const (
	north = iota
	east
	south
	west
)

func (d direction) String() string {
	switch d {
	case north:
		return "north"
	case east:
		return "east"
	case west:
		return "west"
	default:
		return "south"
	}
}

func spiral(loc int) (x int, y int) {
	i := 1
	var direction direction = east

	if loc == 1 {
		return 0, 0
	}
	for pos := 2; pos <= loc; pos++ {
		switch direction {
		case east:
			x++

			// Increase spiral size at the bottom left corner
			if x == i && y == -i {
				i++
				fmt.Println("Increase size to", i)
			}

			if x == i {
				direction = north
				fmt.Printf("Go %s at (%d,%d)\n", direction, x, y)
			}
		case north:
			y++
			if y == i {
				direction = west
				fmt.Printf("Go %s at (%d,%d)\n", direction, x, y)
			}
		case west:
			x--
			if x == -i {
				direction = south
				fmt.Printf("Go %s at (%d,%d)\n", direction, x, y)
			}
		case south:
			y--
			if y == -i {
				direction = east
				fmt.Printf("Go %s at (%d,%d)\n", direction, x, y)
			}
		}
		fmt.Printf("%7d (%2d,%2d) %s %d\n", pos, x, y, direction, i)
	}
	return
}

func main() {
	x, y := spiral(277678)
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	fmt.Println(x + y)
}
