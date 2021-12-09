package main

import (
	"aoc-2021/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

var visited = map[string]bool{}

func main() {
	input := parseInput(utils.ReadLines("./input"))
	p1, p2 := solution(input)
	fmt.Printf("Part 1: %d\nPart 2: %d", p1, p2)
}

func parseInput(lines []string) [][]int {
	heightMap := make([][]int, 0)
	paddingRow := make([]int, 0)
	for i := 0; i < len(lines[0])+2; i++ {
		paddingRow = append(paddingRow, math.MaxInt)
	}

	heightMap = append(heightMap, paddingRow)

	for _, l := range lines {
		lineNumbers := make([]int, 0)
		lineNumbers = append(lineNumbers, math.MaxInt)
		for _, n := range strings.Split(l, "") {
			lineNumbers = append(lineNumbers, utils.ConvertToInt(n))
		}
		lineNumbers = append(lineNumbers, math.MaxInt)
		heightMap = append(heightMap, lineNumbers)
	}
	heightMap = append(heightMap, paddingRow)

	return heightMap
}

func solution(lines [][]int) (int, int) {
	var lowPoints []int
	var basins []int
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[1])-1; x++ {
			n := lines[y][x]
			if n < lines[y+1][x] && n < lines[y-1][x] && n < lines[y][x+1] && n < lines[y][x-1] {
				lowPoints = append(lowPoints, n+1)
				basins = append(basins, countBasinSize(lines, y, x, 1))
			}
		}
	}

	sort.Ints(basins)

	return utils.SumSlice(lowPoints), utils.MulSlice(basins[len(basins)-3:])
}

func countBasinSize(lines [][]int, y, x, count int) int {
	if lines[y][x] == 9 {
		return count - 1
	}

	pos := fmt.Sprintf("%d,%d", y, x)
	if _, isVisited := visited[pos]; isVisited {
		return count - 1
	} else {
		visited[pos] = true
	}

	if y+1 < len(lines)-1 {
		count = countBasinSize(lines, y+1, x, count+1)
	}
	if y-1 > 0 {
		count = countBasinSize(lines, y-1, x, count+1)
	}
	if x+1 < len(lines[1])-1 {
		count = countBasinSize(lines, y, x+1, count+1)
	}
	if x-1 > 0 {
		count = countBasinSize(lines, y, x-1, count+1)
	}

	return count
}
