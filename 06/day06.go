package main

import (
	"aoc-2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ParseInt(strings.Split(utils.ReadFile("./input"), ","), 10)
	initialState := initFish(input)
	fmt.Printf("Part 1: %d\n", countFish(initialState, 80))
	fmt.Printf("Part 2: %d", countFish(initialState, 256))
}

func initFish(input []int) []int {
	initialState := make([]int, 9)
	for _, i := range input {
		initialState[i]++
	}
	return initialState
}

func countFish(fish []int, days int) (count int) {
	for i := 0; i < days; i++ {
		after := make([]int, 9)
		for j, f := range fish {
			if j == 0 {
				after[6] += f // Update old ones
				after[8] += f // Create new ones
			} else {
				after[j-1] += f
			}
		}
		fish = after
	}

	for _, f := range fish {
		count += f
	}
	return
}
