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

	zeroes := make([]int, 12)
	total := 0
	for s.Scan() {
		for i, b := range s.Text() {
			if b == '0' {
				zeroes[i]++
			}
		}

		total++
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	gammaStr, epsilonStr := "", ""
	for _, num := range zeroes {
		if num < total/2 {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}

	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("power consumption:", gamma*epsilon)
}

func part2(f *os.File) {
	s := bufio.NewScanner(f)

	var data []string
	for s.Scan() {
		data = append(data, s.Text())
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	oxygenStr := filter(data, 0, true)
	oxygen, err := strconv.ParseInt(oxygenStr, 2, 64)
	if err != nil {
		panic(err)
	}

	co2Str := filter(data, 0, false)
	co2, err := strconv.ParseInt(co2Str, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("life support rating:", oxygen*co2, oxygen, co2)
}

func filter(data []string, idx int, most bool) string {
	switch len(data) {
	case 1:
		return data[0]
	case 2:
		switch {
		case most && data[0][idx] == '1':
			return data[0]
		case !most && data[0][idx] == '0':
			return data[0]
		default:
			return data[1]
		}
	}

	var zeroes, ones []string
	for _, num := range data {
		if num[idx] == '0' {
			zeroes = append(zeroes, num)
		} else {
			ones = append(ones, num)
		}
	}

	switch {
	case most && len(zeroes) > len(ones):
		return filter(zeroes, idx+1, most)
	case most:
		return filter(ones, idx+1, most)
	case len(zeroes) > len(ones):
		return filter(ones, idx+1, most)
	}

	return filter(zeroes, idx+1, most)
}