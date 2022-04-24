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

	part1(f)
	//part2(f)
}

func part1(f *os.File) {
	s := bufio.NewScanner(f)
	replacer := strings.NewReplacer(" -> ", ",")
	var area [][]int
	for s.Scan() {
		coords := strings.Split(replacer.Replace(s.Text()), ",")
		if len(coords) != 4 {
			panic("invalid input")
		}

		xs, ys, xe, ye := atoi(coords[0]), atoi(coords[1]), atoi(coords[2]), atoi(coords[3])
		if xs != xe && ys != ye {
			continue
		}

		maxX, maxY := max(xs, xe), max(ys, ye)
		if len(area) < maxX {
			for i := len(area); i <= maxX+1; i++ {
				area = append(area, make([]int, maxY+1))
			}
		}

		for x := min(xs, xe); x <= maxX; x++ {
			for y := min(ys, ye); y <= maxY; y++ {
				if len(area[x]) <= maxY {
					area[x] = append(area[x], make([]int, maxY-len(area[x])+1)...)
				}

				area[x][y]++
			}
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	overlaps := 0
	for _, row := range area {
		for _, count := range row {
			if count >= 2 {
				overlaps++
			}
		}
	}

	fmt.Println("points do at least two lines overlap:", overlaps)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
