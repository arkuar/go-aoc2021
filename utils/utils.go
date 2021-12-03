package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ParseInt(lines []string, base int) (result []int) {
	for _, l := range lines {
		if i, err := strconv.ParseInt(l, base, 0); err != nil {
			log.Fatal(err)
		} else {
			result = append(result, int(i))
		}
	}
	return
}

// ReadLines reads a file and returns the lines
func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// ReadIntLines reads a file and returns the lines as integers
func ReadIntLines(filename string) []int {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines []int
	for sc.Scan() {
		lines = append(lines, ConvertToInt(sc.Text()))
	}
	return lines
}

// ReadLinesSplit reads a file using a user defined split function and returns the lines
func ReadLinesSplit(filename string, split bufio.SplitFunc) []string {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(split)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// ReadFile reads a file and returns it as a single string
func ReadFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	check(err)

	return string(buf)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
