package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type image []string

type rule struct {
	match   []int
	pattern image
}

type rulebook []rule

func (im image) checksum(offset int, rotate int) int {
	var flip, rot bool
	var checksum int
	if rotate&1 == 1 {
		// flip the sense of bits
		flip = true
	}
	if rotate&2 == 2 {
		// rotate (columns, not rows)
		rot = true
	}
	for i := 0; i < len(im); i++ {
		if rot {
			if im[i][offset] == '#' {
				if flip {
					checksum += 1 << uint(i)
				} else {
					checksum += 1 << uint(len(im)-1-i)
				}
			}
		} else {
			if im[offset][i] == '#' {
				if flip {
					checksum += 1 << uint(i)
				} else {
					checksum += 1 << uint(len(im)-1-i)
				}
			}
		}
	}
	return checksum
}

func (im image) match(sums []int) bool {
	if len(sums) != len(im) {
		return false
	}
	for r := 0; r < 4; r++ {
		match := true
		for i, sum := range sums {
			if im.checksum(i, r) != sum {
				match = false
				break
			}
		}
		if match {
			return true
		}
		// test again backwards, because bleh
		match = true
		for i, sum := range sums {
			if im.checksum(len(im)-1-i, r) != sum {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func (rb rulebook) match(i image) image {
	for _, r := range rb {
		if len(r.match) != len(i) {
			continue
		}
		if i.match(r.match) {
			fmt.Println(r.match, "=>", r.pattern)
			return r.pattern
		}
	}
	return image{}
}

func readbook(input []string) rulebook {
	var rb rulebook
	for _, line := range input {
		parts := strings.Split(line, " => ")
		pattern := strings.Split(parts[1], "/")
		rows := strings.Split(parts[0], "/")
		ints := make([]int, len(rows))
		for i, row := range rows {
			for _, c := range row {
				ints[i] <<= 1
				if c == '#' {
					ints[i]++
				}
			}
		}
		rb = append(rb, rule{ints, pattern})
	}
	return rb
}

func smoosh(bits [][]image) image {
	size := len(bits[0]) * len(bits[0][0])
	fmt.Println("New size", size)
	img := make(image, size)
	for i := 0; i < len(bits); i++ {
		for j := 0; j < len(bits[i]); j++ {
			for k := 0; k < len(bits[i][j]); k++ {
				img[j*len(bits[i][j])+k] += bits[i][j][k]
			}
		}
	}
	return img
}

func (im image) iterate(rb rulebook) image {
	var newimage [][]image
	if len(im)%2 == 0 {
		newimage = make([][]image, len(im)/2)
		for i := range newimage {
			newimage[i] = make([]image, len(im)/2)
			for j := range newimage[i] {
				newimage[i][j] = make(image, 2)
			}
		}
		for i, row := range im {
			for j := range row {
				if j%2 == 0 {
					newimage[i/2][j/2][i%2] = row[j : j+2]
				}
			}
		}
	} else {
		newimage = make([][]image, len(im)/3)
		for i := range newimage {
			newimage[i] = make([]image, len(im)/3)
			for j := range newimage[i] {
				newimage[i][j] = make(image, 3)
			}
		}
		for i, row := range im {
			for j := range row {
				if j%3 == 0 {
					newimage[i/3][j/3][i%3] = row[j : j+3]
				}
			}
		}
	}
	for i, row := range newimage {
		for j, img := range row {
			newimage[i][j] = rb.match(img)
		}
	}
	return smoosh(newimage)
}

func (im image) count() int {
	count := 0
	for _, row := range im {
		for _, c := range row {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	rb := readbook(strings.Split(strings.TrimSpace(string(input)), "\n"))
	fmt.Println(rb)
	img := image{".#.", "..#", "###"}
	for i := 0; i < 5; i++ {
		img = img.iterate(rb)
		fmt.Println(img)
		fmt.Println(i+1, img.count())
	}
	img = image{".#.", "..#", "###"}
	for i := 0; i < 18; i++ {
		img = img.iterate(rb)
		fmt.Println(i+1, img.count())
	}
}
