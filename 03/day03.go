package main

import (
	"aoc-2021/utils"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./input")
	bits := utils.ParseInt(lines, 2)

	fmt.Printf("Part 1: %d\n", part1(bits, len(lines[0])))
	fmt.Printf("Part 2: %d", part2(bits, len(lines[0])))
}

func part1(bits []int, bitLength int) int {
	var gamma, epsilon int
	for i := 0; i < bitLength; i++ {
		common := mostCommon(bits, i)
		gamma |= common << i
		epsilon |= (common ^ 1) << i
	}
	return gamma * epsilon
}

func part2(bits []int, bitLength int) int {
	ox := bits
	for i := bitLength - 1; i >= 0; i-- {
		common := mostCommon(ox, i)
		ox = matchingBits(common, i, ox)
		if len(ox) == 1 {
			break
		}
	}

	co2 := bits
	for i := bitLength - 1; i >= 0; i-- {
		leastCommon := mostCommon(co2, i) ^ 1
		co2 = matchingBits(leastCommon, i, co2)
		if len(co2) == 1 {
			break
		}
	}
	return ox[0] * co2[0]
}

func matchingBits(common int, idx int, lines []int) (result []int) {
	for _, bits := range lines {
		if (bits>>idx)&1 == common {
			result = append(result, bits)
		}
	}
	return
}

// Find most common bit at index
func mostCommon(lines []int, idx int) (common int) {
	var count int
	for _, bits := range lines {
		count += (bits >> idx) & 1
	}
	if count >= (len(lines)+1)/2 {
		common = 1
	}
	return
}
