package main

import (
	"aoc-2021/utils"
	"fmt"
)

func main() {
	measurements := utils.ReadIntLines("./input")
	part1 := countIncreases(measurements)
	part2 := countSumIncreases(measurements)
	fmt.Printf("Part 1: %d\nPart 2: %d", part1, part2)
}

func countIncreases(measurements []int) (count int) {
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			count++
		}
	}
	return
}

func countSumIncreases(measurements []int) (count int) {
	for i := 2; i < len(measurements)-1; i++ {
		if utils.SumSlice(measurements[i-2:i+1]) < utils.SumSlice(measurements[i-1:i+2]) {
			count++
		}
	}
	return
}
