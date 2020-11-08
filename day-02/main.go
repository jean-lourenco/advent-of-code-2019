package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fmt.Println(part1(scan(file)))
	file.Seek(0, 0)
	fmt.Println(part2(scan(file)))
	// result1 := 5866714
	// result2 := 5208
}

func scan(r io.Reader) <-chan int {
	c := make(chan int)
	go func() {
		scan := bufio.NewScanner(r)
		scan.Scan()
		strs := strings.Split(scan.Text(), ",")
		for _, str := range strs {
			i, _ := strconv.Atoi(str)
			c <- i
		}
		close(c)
	}()

	return c
}

func part1(c <-chan int) int {
	l := make([]int, 0)
	for v := range c {
		l = append(l, v)
	}

	return compute(l, 12, 2)
}

func part2(c <-chan int) int {
	src := make([]int, 0)
	for v := range c {
		src = append(src, v)
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			l := make([]int, len(src))
			copy(l, src)
			result := compute(l, noun, verb)
			if result == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return 0
}

func compute(l []int, noun int, verb int) int {
	l[1] = noun
	l[2] = verb

	for i := 0; i < len(l); i = i + 4 {
		cur := l[i]

		if cur == 99 {
			return l[0]
		} else if cur == 1 {
			l[l[i+3]] = l[l[i+1]] + l[l[i+2]]
		} else {
			l[l[i+3]] = l[l[i+1]] * l[l[i+2]]
		}
	}

	return 0
}
