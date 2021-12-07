package main

import (
	"aoc-2021/utils"
	"fmt"
)

func main() {
	input := utils.ParseInt(utils.ReadLinesSplit("./input", ','), 10)
	median := utils.Median(input)
	mean := utils.Mean(input)

	fmt.Printf("Part 1: %d\nPart 2: %d", calculateFuel(input, median, true), calculateFuel(input, mean, false))
}

func calculateFuel(numbers []int, target int, constantFuelRate bool) (fuelCount int) {
	for _, number := range numbers {
		fuelCost := 1
		for number != target {
			if number < target {
				number++
			} else {
				number--
			}
			fuelCount += fuelCost
			if !constantFuelRate {
				fuelCost++
			}
		}
	}
	return
}
