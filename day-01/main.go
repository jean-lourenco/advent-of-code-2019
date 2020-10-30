package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	fmt.Println(part1(scanLines(f)))
	f.Seek(0, 0)
	fmt.Println(part2(scanLines(f)))
}

func scanLines(r io.Reader) <-chan int {
	c := make(chan int)
	go func() {
		scan := bufio.NewScanner(r)

		for scan.Scan() {
			mass, _ := strconv.Atoi(scan.Text())
			c <- mass
		}
		close(c)
	}()
	return c
}

func part1(ch <-chan int) int {
	sum := 0
	for mass := range ch {
		sum += calcFuelRequirement(mass)
	}

	return sum
}

func part2(ch <-chan int) int {
	sum := 0
	for mass := range ch {
		cur := calcFuelRequirement(mass)
		sum += cur

		for {
			cur = calcFuelRequirement(cur)
			if cur < 1 {
				break
			}
			sum += cur
		}
	}

	return sum
}

func calcFuelRequirement(mass int) int {
	return mass/3 - 2
}
