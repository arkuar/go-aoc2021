package main

import (
	"aoc-2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadLines("./input")
	fmt.Printf("Part 1: %d\n", countOverlaps(input, false))
	fmt.Printf("Part 2: %d", countOverlaps(input, true))
}

func countOverlaps(entries []string, checkDiagonals bool) (overlaps int) {
	var diagram [999][999]int
	for _, entry := range entries {
		var x1, y1, x2, y2 int
		fmt.Sscanf(entry, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		sx, sy := step(x1, x2), step(y1, y2)

		if sx*sy != 0 && !checkDiagonals {
			continue
		}

		var steps int
		if sx != 0 {
			steps = utils.Abs(x1 - x2)
		} else {
			steps = utils.Abs(y1 - y2)
		}

		for i := 0; i <= steps; i++ {
			diagram[y1][x1]++
			if diagram[y1][x1] == 2 {
				overlaps++
			}
			y1 += sy
			x1 += sx
		}
	}
	return
}

func step(a, b int) int {
	if b < a {
		return -1
	}
	if b > a {
		return 1
	}
	return 0
}
