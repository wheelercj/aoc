package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func getLoc(id int, allMaps [][]Map) int {
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
	lowestLocNum := getLoc(seeds[0], allMaps)
	for i := 1; i < len(seeds); i++ {
		locNum := getLoc(seeds[i], allMaps)
		if locNum < lowestLocNum {
			lowestLocNum = locNum
		}
	}

	fmt.Println("part 1 result:", lowestLocNum)
}

func part2(seedRanges []int, allMaps [][]Map) {
	start := time.Now()

	locsCh := make(chan int, len(seedRanges)/2)
	for i := 0; i < len(seedRanges); i += 2 {
		go getLowestLocNum(seedRanges[i], seedRanges[i+1], allMaps, locsCh)
	}

	lowestLocNum := <-locsCh
	for i := 0; i < len(seedRanges)/2-1; i++ {
		locNum := <-locsCh
		if locNum < lowestLocNum {
			lowestLocNum = locNum
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("part 2 result: %d (took %v)", lowestLocNum, elapsed)
}

func getLowestLocNum(seedRangeStart, rangeLength int, allMaps [][]Map, locsCh chan int) {
	lowestLocNum := getLoc(seedRangeStart, allMaps)
	for seed := seedRangeStart; seed < seedRangeStart+rangeLength; seed++ {
		locNum := getLoc(seed, allMaps)
		if locNum < lowestLocNum {
			lowestLocNum = locNum
		}
	}
	locsCh <- lowestLocNum
}
