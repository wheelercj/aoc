package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "./input.txt"
	fmt.Println("using input from file", fileName)
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	patterns := parseInput(string(content))

	result := part1(patterns)
	fmt.Println("part 1 result:", result)
}

func parseInput(content string) [][]string {
	content = strings.TrimSpace(content)
	patternStrs := strings.Split(content, "\r\n\r\n")

	var patterns [][]string
	for _, patternStr := range patternStrs {
		patterns = append(patterns, strings.Split(patternStr, "\r\n"))
	}

	return patterns
}

func part1(patterns [][]string) int {
	var result int
	for _, pattern := range patterns {
		row, col := findMirror(pattern)
		if col > 0 {
			result += col
		} else {
			result += row * 100
		}
	}

	return result
}

func findMirror(pattern []string) (row, col int) {
	// search for 2 identical rows
	for y := 0; y < len(pattern)-1; y++ {
		if isMirrorAtRow(pattern, y) {
			fmt.Printf("found mirror between rows %d and %d\n", y+1, y+2)
			return y + 1, 0
		}
	}

	// search for 2 identical columns
	for x := 0; x < len(pattern[0])-1; x++ {
		if isMirrorAtCol(pattern, x) {
			fmt.Printf("found mirror between cols %d and %d\n", x+1, x+2)
			return 0, x + 1
		}
	}

	panic("No mirror found")
}

func isMirrorAtRow(pattern []string, y int) bool {
	if !rowsMatch(pattern, y, y+1) {
		return false
	}

	y1 := y
	y2 := y + 1
	for {
		y1--
		y2++
		if y1 < 0 || y2 >= len(pattern) {
			return true
		}
		if !rowsMatch(pattern, y1, y2) {
			return false
		}
	}
}

func rowsMatch(pattern []string, y1, y2 int) bool {
	for x, ch1 := range pattern[y1] {
		if ch1 != rune(pattern[y2][x]) {
			return false
		}
	}

	return true
}

func isMirrorAtCol(pattern []string, x int) bool {
	if !colsMatch(pattern, x, x+1) {
		return false
	}

	x1 := x
	x2 := x + 1
	for {
		x1--
		x2++
		if x1 < 0 || x2 >= len(pattern[0]) {
			return true
		}
		if !colsMatch(pattern, x1, x2) {
			return false
		}
	}
}

func colsMatch(pattern []string, x1, x2 int) bool {
	for _, row := range pattern {
		if row[x1] != row[x2] {
			return false
		}
	}

	return true
}
