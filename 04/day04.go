package main

import (
	"aoc-2021/utils"
	"errors"
	"fmt"
	"strings"
)

func main() {
	input := strings.Split(utils.ReadFile("./input"), "\n\n")
	numbers := utils.ParseInt(strings.Split(input[0], ","), 10)

	boards := buildBoards(input[1:])
	winner, number, err := findBoard(boards, numbers, true)
	utils.Check(err)
	fmt.Printf("Part 1: %d\n", calculateScore(winner, number))

	last, number, err := findBoard(boards, numbers, false)
	utils.Check(err)
	fmt.Printf("Part 2: %d", calculateScore(last, number))
}

func calculateScore(board [][]int, number int) int {
	var unmarked int
	for _, row := range board {
		for _, num := range row {
			if num != -1 {
				unmarked += num
			}
		}
	}
	return unmarked * number
}

func buildBoards(input []string) (boards [][][]int) {
	var board [][]int
	for _, l := range input {
		for _, row := range strings.Split(l, "\n") {
			board = append(board, utils.ParseInt(strings.Fields(row), 10))
		}
		boards = append(boards, board)
		board = make([][]int, 0)
	}
	return
}

func findBoard(boards [][][]int, numbers []int, winning bool) ([][]int, int, error) {
	winners := make(map[int]bool)
	for _, number := range numbers {
		for id, board := range boards {
			markNumber(board, number)
			if isWinner(board) {
				if winning {
					return board, number, nil
				}
				winners[id] = true
				if len(winners) == len(boards) && !winning {
					return board, number, nil
				}
			}
		}
	}
	return nil, -1, errors.New("no winners")
}

func markNumber(board [][]int, number int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if board[y][x] == number {
				board[y][x] = -1
			}
		}
	}
}

func isWinner(board [][]int) bool {
	for i := 0; i < 5; i++ {
		rowMatch := 0
		colMatch := 0
		for j := 0; j < 5; j++ {
			if board[i][j] == -1 {
				rowMatch++
			}
			if board[j][i] == -1 {
				colMatch++
			}
		}
		if colMatch == 5 || rowMatch == 5 {
			return true
		}
	}

	return false
}
