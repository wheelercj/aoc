package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	destStart int
	srcStart  int
	length    int
}

func main() {
	seeds, allMaps := parseInput("input.txt")
	part1(seeds, allMaps)
	part2(seeds, allMaps)
}

func parseInput(fileName string) ([]int, [][]Map) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	seedsAndCategoryStrs := strings.Split(string(bytes), "\r\n\r\n")
	categoryCount := len(seedsAndCategoryStrs) - 1

	seedsStr := seedsAndCategoryStrs[0]
	seedsStrs := strings.Split(seedsStr, " ")
	seeds := make([]int, len(seedsStrs)-1)
	for i := 0; i < len(seeds); i++ {
		seeds[i] = mustParseInt(seedsStrs[i+1])
	}

	allMaps := make([][]Map, categoryCount)
	for i := range allMaps {
		allMaps[i] = getCategoryMap(seedsAndCategoryStrs[i+1])
	}

	return seeds, allMaps
}

func getCategoryMap(categoryStr string) []Map {
	categoryParts := strings.Split(strings.Trim(categoryStr, "\r\n"), ":\r\n")
	lines := strings.Split(categoryParts[1], "\r\n")
	categoryMap := make([]Map, len(lines))
	for i, line := range lines {
		numStrs := strings.Split(line, " ")
		destStart := mustParseInt(numStrs[0])
		srcStart := mustParseInt(numStrs[1])
		length := mustParseInt(numStrs[2])
		categoryMap[i] = Map{
			destStart: destStart,
			srcStart:  srcStart,
			length:    length,
		}
	}
	return categoryMap
}

func getLocation(id int, allMaps [][]Map) int {
	for i := range allMaps {
		id = getCategoryValue(id, allMaps[i])
	}
	return id
}

func getCategoryValue(id int, categoryMaps []Map) int {
	for _, categoryMap := range categoryMaps {
		destStart := categoryMap.destStart
		srcStart := categoryMap.srcStart
		length := categoryMap.length
		if id >= srcStart && id < srcStart+length {
			diff := destStart - srcStart
			return id + diff
		}
	}
	return id
}

func mustParseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func part1(seeds []int, allMaps [][]Map) {
	lowestLocationNum := getLocation(seeds[0], allMaps)
	for i := 1; i < len(seeds); i++ {
		locationNum := getLocation(seeds[i], allMaps)
		if locationNum < lowestLocationNum {
			lowestLocationNum = locationNum
		}
	}

	fmt.Println("lowest location number:", lowestLocationNum)
}

func part2(seedRanges []int, allMaps [][]Map) {
	lowestLocationNum := getLowestLocationNum(seedRanges[0], seedRanges[1], allMaps)
	for i := 2; i < len(seedRanges); i += 2 {
		locationNum := getLowestLocationNum(seedRanges[i], seedRanges[i+1], allMaps)
		if locationNum < lowestLocationNum {
			lowestLocationNum = locationNum
		}
	}

	fmt.Println("lowest location number:", lowestLocationNum)
}

func getLowestLocationNum(seedRangeStart, rangeLength int, allMaps [][]Map) int {
	lowestLocationNum := getLocation(seedRangeStart, allMaps)
	for seed := seedRangeStart; seed < seedRangeStart+rangeLength; seed++ {
		locationNum := getLocation(seed, allMaps)
		if locationNum < lowestLocationNum {
			lowestLocationNum = locationNum
		}
	}
	return lowestLocationNum
}
