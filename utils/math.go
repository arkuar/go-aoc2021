package utils

import (
	"strconv"
)

// ConvertToInt converts a string to integer
func ConvertToInt(s string) int {
	value, err := strconv.Atoi(s)
	Check(err)
	return value
}

// SumSlice sums the contents of the slice
func SumSlice(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
