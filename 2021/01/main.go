package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//part1(f)
	part2(f)
}

func part1(f *os.File) {
	s := bufio.NewScanner(f)

	times := 0
	prev := -1

	for i := 0; s.Scan(); i++ {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		if i > 0 && num > prev {
			times++

		}

		prev = num
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("number of times a depth measurement increases:", times)
}

func part2(f *os.File) {
	s := bufio.NewScanner(f)

	times := 0
	buf := [3]int{}

	for i := 0; s.Scan(); i++ {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		if i < 3 {
			buf[i] = num

			continue
		}

		if num > buf[0] {
			times++
		}

		buf[0], buf[1], buf[2] = buf[1], buf[2], num
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("number of times the sum of measurements in this sliding window increases:", times)
}
