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

	result1 := getMirrorsTotal(patterns, 0)
	fmt.Println("part 1 result:", result1)

	result2 := getMirrorsTotal(patterns, 1)
	fmt.Println("part 2 result:", result2)
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

// getMirrorsTotal gets the final, combined score of all the patterns.
func getMirrorsTotal(patterns [][]string, smudgeCount int) int {
	var result int
	for _, pattern := range patterns {
		row, col := findMirror(pattern, smudgeCount)
		if col > 0 {
			result += col
		} else {
			result += row * 100
		}
	}

	return result
}

// findMirror returns the row or column of the mirror in a pattern with smudgeCount
// smudges.
func findMirror(pattern []string, smudgeCount int) (row, col int) {
	// search for 2 identical rows
	for y := 0; y < len(pattern)-1; y++ {
		diffCount := mirrorRowDiff(pattern, y)
		if diffCount == smudgeCount {
			// fmt.Printf("found mirror between rows %d and %d\n", y+1, y+2)
			return y + 1, 0
		}
	}

	// search for 2 identical columns
	for x := 0; x < len(pattern[0])-1; x++ {
		diffCount := mirrorColDiff(pattern, x)
		if diffCount == smudgeCount {
			// fmt.Printf("found mirror between cols %d and %d\n", x+1, x+2)
			return 0, x + 1
		}
	}

	panic("No mirror found")
}

// mirrorRowDiff counts up to two differences in a pattern assuming a mirror exists at
// row y.
func mirrorRowDiff(pattern []string, y int) (diffCount int) {
	diffCount = rowsDiff(pattern, y, y+1)
	if diffCount > 1 {
		return diffCount
	}

	y1 := y
	y2 := y + 1
	for {
		y1--
		y2++
		if y1 < 0 || y2 >= len(pattern) {
			return diffCount
		}
		diffCount += rowsDiff(pattern, y1, y2)
		if diffCount > 1 {
			return diffCount
		}
	}
}

// rowsDiff counts up to two differences between two rows.
func rowsDiff(pattern []string, y1, y2 int) (diffCount int) {
	for x, ch1 := range pattern[y1] {
		if ch1 != rune(pattern[y2][x]) {
			if diffCount == 0 {
				diffCount = 1
			} else {
				return 2
			}
		}
	}

	return diffCount
}

// mirrorColDiff counts up to two differences in a pattern assuming a mirror exists at
// column x.
func mirrorColDiff(pattern []string, x int) (diffCount int) {
	diffCount = colsDiff(pattern, x, x+1)
	if diffCount > 1 {
		return diffCount
	}

	x1 := x
	x2 := x + 1
	for {
		x1--
		x2++
		if x1 < 0 || x2 >= len(pattern[0]) {
			return diffCount
		}
		diffCount += colsDiff(pattern, x1, x2)
		if diffCount > 1 {
			return diffCount
		}
	}
}

// colsDiff counts up to two differences between two columns.
func colsDiff(pattern []string, x1, x2 int) (diffCount int) {
	for _, row := range pattern {
		if row[x1] != row[x2] {
			if diffCount == 0 {
				diffCount = 1
			} else {
				return 2
			}
		}
	}

	return diffCount
}
