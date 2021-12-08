package main

import (
	"aoc-2021/utils"
	"fmt"
	"strings"
)

var uniqueDigits = map[int]int{2: 1, 4: 4, 3: 7, 7: 8}

func main() {
	lines := utils.ReadLines("./input")
	inputs, outputs := parseLines(lines)
	fmt.Printf("Part 1: %d\n", uniqueOutputDigits(outputs))
	fmt.Printf("Part 2: %d", outputSum(inputs, outputs))
}

func parseLines(lines []string) (inputs []string, outputs []string) {
	for _, l := range lines {
		splitLines := strings.Split(l, " | ")
		inputs = append(inputs, splitLines[0])
		outputs = append(outputs, splitLines[1])
	}
	return
}

func uniqueOutputDigits(outputs []string) (count int) {
	for _, output := range outputs {
		for _, signal := range strings.Fields(output) {
			if _, ok := uniqueDigits[len(signal)]; ok {
				count++
			}
		}
	}
	return
}

func outputSum(inputs []string, outputs []string) (sum int) {
	for i := 0; i < len(inputs); i++ {
		inputEntry, outputEntry := strings.Fields(inputs[i]), strings.Fields(outputs[i])
		sum += deduceOutput(inputEntry, outputEntry)
	}
	return
}

func deduceOutput(input []string, output []string) (result int) {
	deduced := make(map[int]string, 10)
	// Deduce unique digits
	for _, inputDigit := range input {
		if val, ok := uniqueDigits[len(inputDigit)]; ok {
			deduced[val] = utils.SortString(inputDigit)
		}
	}

	// Deduce the rest
	for _, inputSignal := range input {
		sortedSignal := utils.SortString(inputSignal)
		if len(inputSignal) == 5 {
			if similarity := compareSignals(sortedSignal, deduced[1]); similarity == 2 {
				deduced[3] = sortedSignal
			} else if similarity := compareSignals(sortedSignal, deduced[4]); similarity == 3 {
				deduced[5] = sortedSignal
			} else {
				deduced[2] = sortedSignal
			}
		} else if len(inputSignal) == 6 {
			if similarity := compareSignals(sortedSignal, deduced[4]); similarity == 4 {
				deduced[9] = sortedSignal
			} else if similarity := compareSignals(sortedSignal, deduced[1]); similarity == 2 {
				deduced[0] = sortedSignal
			} else {
				deduced[6] = sortedSignal
			}
		}
	}

	// Decode output
	for _, segment := range output {
		sorted := utils.SortString(segment)
		result *= 10
		for value, digit := range deduced {
			if sorted == digit {
				result += value
			}
		}
	}
	return result
}

// Check the similarity of the signals
func compareSignals(signal, comparison string) (similarity int) {
	for _, num := range strings.Split(signal, "") {
		if found := strings.Index(comparison, num); found != -1 {
			similarity++
		}
	}
	return
}
