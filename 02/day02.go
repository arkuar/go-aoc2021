package main

import (
	"aoc-2021/utils"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./input")
	p1, p2 := solution(lines)
	fmt.Printf("Part 1: %d\nPart 2: %d", p1, p2)
}

func solution(lines []string) (int, int) {
	var x, y1, y2, aim int
	for _, l := range lines {
		var (
			command string
			steps   int
		)

		fmt.Sscanf(l, "%s %d", &command, &steps)

		switch command {
		case "forward":
			x += steps
			y2 += aim * steps
		case "down":
			y1 += steps
			aim += steps
		case "up":
			y1 -= steps
			aim -= steps
		}
	}

	return x * y1, x * y2
}
