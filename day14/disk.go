package main

import (
	"fmt"
	"strconv"

	"github.com/jamesmcdonald/adventofcode-2017/day10/knot"
)

const input string = "ugkiagan"

func countHashBits(hash [16]int) int {
	bits := 0
	for _, i := range hash {
		for j := i; j != 0; j >>= 1 {
			bits += j & 1
		}
	}
	return bits
}

type block [128]bool

type diskspace [128]block

type disk struct {
	diskspace
	regionmap [128][128]int
}

func (b block) String() string {
	var buffer [128]rune
	for i := 0; i < 128; i++ {
		if b[i] {
			buffer[i] = '#'
		} else {
			buffer[i] = '.'
		}
	}
	return string(buffer[:])
}

func (d disk) String() string {
	var buffer [128 * 129]rune
	for i := 0; i < 128; i++ {
		buffer[(i+1)*129-1] = '\n'
		for j := 0; j < 128; j++ {
			if d.diskspace[i][j] {
				if d.regionmap[i][j] != 0 {
					s := strconv.Itoa(d.regionmap[i][j])
					buffer[i*128+i+j] = rune(s[len(s)-1])
				} else {
					buffer[i*128+i+j] = '#'
				}
			} else {
				if d.regionmap[i][j] != 0 {
					buffer[i*128+i+j] = '!'
				} else {
					buffer[i*128+i+j] = '.'
				}
			}
		}
	}
	return string(buffer[:])
}

func (d *disk) markgroup(i int, j int, group int) {
	d.regionmap[i][j] = group
	if j > 0 && d.diskspace[i][j-1] && d.regionmap[i][j-1] == 0 {
		d.markgroup(i, j-1, group)
	}
	if j < 127 && d.diskspace[i][j+1] && d.regionmap[i][j+1] == 0 {
		d.markgroup(i, j+1, group)
	}
	if i > 0 && d.diskspace[i-1][j] && d.regionmap[i-1][j] == 0 {
		d.markgroup(i-1, j, group)
	}
	if i < 127 && d.diskspace[i+1][j] && d.regionmap[i+1][j] == 0 {
		d.markgroup(i+1, j, group)
	}
}

func (d *disk) itermarkgroup(i int, j int, group int) {
	type coord struct {
		x int
		y int
	}
	queue := []coord{{i, j}}
	add := func(c coord) {
		for _, q := range queue {
			if q == c {
				return
			}
		}
		queue = append(queue, c)
	}
	for len(queue) > 0 {
		i, j := queue[0].x, queue[0].y
		d.regionmap[i][j] = group
		if j > 0 && d.diskspace[i][j-1] && d.regionmap[i][j-1] == 0 {
			add(coord{i, j - 1})
		}
		if j < 127 && d.diskspace[i][j+1] && d.regionmap[i][j+1] == 0 {
			add(coord{i, j + 1})
		}
		if i > 0 && d.diskspace[i-1][j] && d.regionmap[i-1][j] == 0 {
			add(coord{i - 1, j})
		}
		if i < 127 && d.diskspace[i+1][j] && d.regionmap[i+1][j] == 0 {
			add(coord{i + 1, j})
		}
		queue = queue[1:]
	}

}

func (d *disk) chanmarkgroup(i int, j int, group int) {
	type coord struct {
		x int
		y int
	}
	queue := make(chan coord, 20)
	queue <- coord{i, j}
	for {
		select {
		case c := <-queue:
			i, j := c.x, c.y
			if d.regionmap[i][j] != 0 {
				continue
			}
			d.regionmap[i][j] = group
			if j > 0 && d.diskspace[i][j-1] && d.regionmap[i][j-1] == 0 {
				queue <- coord{i, j - 1}
			}
			if j < 127 && d.diskspace[i][j+1] && d.regionmap[i][j+1] == 0 {
				queue <- coord{i, j + 1}
			}
			if i > 0 && d.diskspace[i-1][j] && d.regionmap[i-1][j] == 0 {
				queue <- coord{i - 1, j}
			}
			if i < 127 && d.diskspace[i+1][j] && d.regionmap[i+1][j] == 0 {
				queue <- coord{i + 1, j}
			}
		default:
			close(queue)
			return
		}
	}
}

func (d *disk) countRegions() int {
	var groups int
	gid := 1
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if d.regionmap[i][j] != 0 {
				continue
			}
			if d.diskspace[i][j] {
				groups++
				d.chanmarkgroup(i, j, gid)
				gid++
			}
		}
	}
	fmt.Println(gid)
	return groups
}

func hashtoblock(hash [16]int) block {
	var output [128]bool
	for i := 0; i < 16; i++ {
		for j := 0; j < 8; j++ {
			if hash[i]&1 == 1 {
				output[i*8+j] = true
			}
			hash[i] >>= 1
		}
	}
	return output
}

func main() {
	bits := 0
	var d disk
	for i := 0; i < 128; i++ {
		str := fmt.Sprintf("%s-%d", input, i)
		hash := knot.Hash(knot.Knot(knot.ParseASCII(str), 64))
		bits += countHashBits(hash)
		d.diskspace[i] = hashtoblock(hash)
	}
	fmt.Println(bits)
	fmt.Println(d.countRegions())
	fmt.Println(d)
}
