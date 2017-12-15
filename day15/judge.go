package main

import "fmt"

func generator(start int, factor int, divisor int) (chan int, chan bool) {
	previous := start
	output := make(chan int)
	control := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-control:
				close(output)
				return
			default:
			}
			next := previous * factor % 2147483647
			for next%divisor != 0 {
				next = next * factor % 2147483647
			}
			output <- next
			previous = next
		}
	}()
	return output, control
}

func generatorA(divisor int) (chan int, chan bool) {
	return generator(618, 16807, divisor)
}

func generatorB(divisor int) (chan int, chan bool) {
	return generator(814, 48271, divisor)
}

func judge(iterations int, in1 chan int, s1 chan bool, in2 chan int, s2 chan bool) int {
	count := 0
	for i := 0; i < iterations; i++ {
		if i == iterations-1 {
			s1 <- true
			s2 <- true
		}

		i1 := <-in1
		i2 := <-in2
		if i1&0xffff == i2&0xffff {
			count++
		}
	}
	return count
}

func main() {
	g1, s1 := generatorA(1)
	g2, s2 := generatorB(1)
	fmt.Println(judge(40000000, g1, s1, g2, s2))
	g1, s1 = generatorA(4)
	g2, s2 = generatorB(8)
	fmt.Println(judge(5000000, g1, s1, g2, s2))

}
