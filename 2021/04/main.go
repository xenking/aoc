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

	s := bufio.NewScanner(f)

	s.Scan()

	var nums []int
	nn := strings.Split(s.Text(), ",")
	for _, n := range nn {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	var boards [][25]int
	board := [25]int{}
	j := 0

	for s.Scan() {
		row := s.Text()
		if row == "" {
			continue
		}

		fields := strings.Fields(row)
		if len(fields) != 5 {
			panic("invalid input")
		}

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}

			board[j] = num
			j++
		}

		if j == 25 {
			j = 0
			boards = append(boards, board)
		}
	}

	//part1(nums, boards)
	part2(nums, boards)
}

func part1(nums []int, boards [][25]int) {
	num, numSum, idx := getWinningBoard(nums, boards)
	if num == -1 {
		panic("no winning board")
	}

	boardSum := 0
	for _, n := range boards[idx] {
		boardSum += n
	}
	sum := boardSum - numSum

	fmt.Println("winning final score:", num*sum)
}

func getWinningBoard(nums []int, boards [][25]int) (int, int, int) {
	rows := make([][5]int, len(boards))
	cols := make([][5]int, len(boards))
	boardSum := make([]int, len(boards))
	for _, num := range nums {
		for k, b := range boards {
			for i, n := range b {
				if n != num {
					continue
				}

				rows[k][i/5]++
				cols[k][i%5]++
				boardSum[k] += n

				if rows[k][i/5] == 5 || cols[k][i%5] == 5 {
					return num, boardSum[k], k
				}
			}
		}
	}

	return -1, -1, 0
}

func part2(nums []int, boards [][25]int) {
	num, numSum, board := getLosingBoard(nums, boards)
	if num == -1 {
		panic("no losing board")
	}

	boardSum := 0
	for _, n := range board {
		boardSum += n
	}
	sum := boardSum - numSum

	fmt.Println("losing final score:", num*sum)
}

func getLosingBoard(nums []int, boards [][25]int) (int, int, [25]int) {
	rows := make([][5]int, len(boards))
	cols := make([][5]int, len(boards))
	boardSum := make([]int, len(boards))
	excluded := make(map[int]struct{})
	for _, num := range nums {
		for k, b := range boards {
			if _, ok := excluded[k]; ok {
				continue
			}

			for i, n := range b {
				if n != num {
					continue
				}

				rows[k][i/5]++
				cols[k][i%5]++
				boardSum[k] += n

				if rows[k][i/5] == 5 || cols[k][i%5] == 5 {

					excluded[k] = struct{}{}
					if len(excluded) == len(boards) {
						return num, boardSum[k], b
					}

					break
				}
			}
		}
	}

	return -1, -1, [25]int{}
}
