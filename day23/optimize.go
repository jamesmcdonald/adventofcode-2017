package main

func main() {
	var a, b, c, d, e, f, g, h int
	a = 1
	b = 57
	c = b
	if a != 0 {
		goto lab1
	}
	if 1 != 0 {
		goto lab2
	}
lab1:
	b = b * 100
	b += 100000
	c = b
	c += 17000
lab2:
	f = 1
	d = 2
lab5:
	e = 2
lab4:
	g = d
	g *= e
	g -= b
	if g != 0 {
		goto lab3
	}
	f = 0
lab3:
	e++
	g = e
	g -= b
	if g != 0 {
		goto lab4
	}
	d++
	g = d
	g -= b
	if g != 0 {
		goto lab5
	}
	if f != 0 {
		goto lab6
	}
	h++
lab6:
	g = b
	g -= c
	if g != 0 {
		goto lab7
	}
	if 1 != 0 {
		goto lab8
	}
lab7:
	b += 17
	if 1 != 0 {
		goto lab2
	}
lab8:
}
