package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type vector struct {
	x int
	y int
	z int
}

func (v vector) add(a vector) vector {
	v.x += a.x
	v.y += a.y
	v.z += a.z
	return v
}

type particle struct {
	p    vector
	v    vector
	a    vector
	dead bool
}

func parsevec(v string) vector {
	vs := v[3 : len(v)-1]
	coords := strings.Split(vs, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(coords[2])
	if err != nil {
		panic(err)
	}
	return vector{x, y, z}
}

func parse(l string) particle {
	parts := strings.Split(l, ", ")
	return particle{parsevec(parts[0]), parsevec(parts[1]), parsevec(parts[2]), false}
}

func (p particle) animate() particle {
	p.v = p.v.add(p.a)
	p.p = p.p.add(p.v)
	return p
}

func (p particle) collide(q particle) bool {
	if p.p.x == q.p.x && p.p.y == q.p.y && p.p.z == q.p.z && !q.dead {
		return true
	}
	return false
}

func (p particle) distance() int {
	x := p.p.x
	y := p.p.y
	z := p.p.z
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}
	return x + y + z
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	particles := make([]particle, len(lines))
	for i, l := range lines {
		particles[i] = parse(l)
	}
	for i := 0; i < 100000; i++ {
		for j, p := range particles {
			particles[j] = p.animate()
		}
	}
	min := particles[0].distance()
	minp := 0
	for i, p := range particles[1:] {
		distance := p.distance()
		if distance < min {
			min = distance
			minp = i
		}
	}
	fmt.Println(minp, min)

	// reload particles
	particles = make([]particle, len(lines))
	for i, l := range lines {
		particles[i] = parse(l)
	}

	for i := 0; i < 1000; i++ {
		for j, p := range particles {
			if p.dead {
				continue
			}
			particles[j] = p.animate()
		}
		for j, p := range particles {
			if p.dead {
				continue
			}
			for k, q := range particles {
				if j != k && p.collide(q) {
					fmt.Println("Collision", p, q)
					particles[j].dead = true
					particles[k].dead = true
				}
			}
		}
	}
	count := 0
	for _, p := range particles {
		if !p.dead {
			count++
		}
	}
	fmt.Println(count)
}
