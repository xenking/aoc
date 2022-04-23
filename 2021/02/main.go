package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	depth, pos := 0, 0

	for i := 0; s.Scan(); i++ {
		fields := strings.Fields(s.Text())
		if len(fields) != 2 {
			panic("invalid input")
		}

		value, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		switch fields[0] {
		case "forward":
			pos += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("planned course:", depth*pos)
}

func part2(f *os.File) {
	s := bufio.NewScanner(f)

	depth, pos, aim := 0, 0, 0

	for i := 0; s.Scan(); i++ {
		fields := strings.Fields(s.Text())
		if len(fields) != 2 {
			panic("invalid input")
		}

		value, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		switch fields[0] {
		case "forward":
			pos += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("planned course:", depth*pos)
}
