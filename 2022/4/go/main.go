// https://adventofcode.com/2022/day/4
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contentBytes), "\r\n")

	var fullyOverlapCount int
	for _, line := range lines {
		elfPair := strings.Split(line, ",")
		firstElf, secondElf := parseRange(elfPair[0]), parseRange(elfPair[1])
		if fullyOverlaps(firstElf, secondElf) || fullyOverlaps(secondElf, firstElf) {
			fullyOverlapCount++
		}
	}

	if fullyOverlapCount != 305 {
		panic(fullyOverlapCount)
	}
	fmt.Println(fullyOverlapCount)
}

func parseRange(rangeStr string) [2]int {
	numStrs := strings.Split(rangeStr, "-")
	a, err := strconv.Atoi(numStrs[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(numStrs[1])
	if err != nil {
		panic(err)
	}
	return [2]int{a, b}
}

func fullyOverlaps(range1, range2 [2]int) bool {
	return range1[0] <= range2[0] && range1[1] >= range2[1]
}

func overlaps(range1, range2 [2]int) bool {
	return range1[0] <= range2[0] && range1[1] >= range2[0] ||
		range2[0] <= range1[0] && range2[1] >= range1[0] ||
		range1[0] <= range2[1] && range1[1] >= range2[1] ||
		range2[0] <= range1[1] && range2[1] >= range1[1]
}

func part2() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contentBytes), "\r\n")

	var overlapCount int
	for _, line := range lines {
		elfPair := strings.Split(line, ",")
		firstElf, secondElf := parseRange(elfPair[0]), parseRange(elfPair[1])
		if overlaps(firstElf, secondElf) {
			overlapCount++
		}
	}

	if overlapCount != 811 {
		panic(overlapCount)
	}
	fmt.Println(overlapCount)
}
