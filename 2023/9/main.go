package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sequences := parseInput("input.txt")
	part1(sequences)
}

func parseInput(fileName string) [][]int {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")
	sequences := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		numsStrs := strings.Split(lines[i], " ")
		sequences[i] = make([]int, len(numsStrs))
		for j := range sequences[i] {
			sequences[i][j] = mustParseInt(numsStrs[j])
		}
	}
	return sequences
}

func mustParseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

// getNextDiffValue finds the next value in the sequence that is the sequences of
// differences of s.
func getNextDiffValue(s []int) int {
	diffs := make([]int, len(s)-1)
	allZeroes := true
	for j := 0; j < len(diffs); j++ {
		diffs[j] = s[j+1] - s[j]
		if diffs[j] != 0 {
			allZeroes = false
		}
	}
	if allZeroes {
		return 0
	} else {
		return diffs[len(diffs)-1] + getNextDiffValue(diffs)
	}
}

func part1(sequences [][]int) {
	var sum int
	for _, s := range sequences {
		sum += s[len(s)-1] + getNextDiffValue(s)
	}

	fmt.Println("part 1 result:", sum)
}
