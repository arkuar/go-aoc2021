package main

import (
	"aoc-2021/utils"
	"fmt"
	"sort"
)

var errorScores = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var completionScores = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
var symbols = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

func main() {
	lines := utils.ReadLines("./input")
	p1, p2 := findScores(lines)
	fmt.Printf("Part 1: %d\nPart 2: %d", p1, p2)
}

func findScores(lines []string) (errorScore, middleScore int) {
	var stack []rune
	var incompleteScores []int
	for _, line := range lines {
		for i, token := range line {
			// Check if token is an opening or closing symbol
			if _, isOpening := symbols[token]; isOpening {
				stack = append(stack, token)
			} else {
				// Closing symbol has to have a matching opening at the top of the stack
				// otherwise line is invalid
				top := stack[len(stack)-1]
				if expected, ok := symbols[top]; ok && expected == token {
					stack = stack[:len(stack)-1]
				} else {
					errorScore += errorScores[token]
					stack = make([]rune, 0)
					break
				}
			}
			// Incomplete line, calculate the score of completing it
			if len(line)-1 == i {
				total := 0
				for len(stack) != 0 {
					total = (total * 5) + completionScores[symbols[stack[len(stack)-1]]]
					stack = stack[:len(stack)-1]
				}
				incompleteScores = append(incompleteScores, total)
			}
		}
	}

	sort.Ints(incompleteScores)
	middleScore = incompleteScores[len(incompleteScores)/2]

	return
}
