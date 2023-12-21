package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	destStart int
	destEnd   int
	srcStart  int
	srcEnd    int
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
	seedsAndCatStrs := strings.Split(string(bytes), "\r\n\r\n")
	catCount := len(seedsAndCatStrs) - 1

	seedsStr := seedsAndCatStrs[0]
	seedsStrs := strings.Split(seedsStr, " ")
	seeds := make([]int, len(seedsStrs)-1)
	for i := 0; i < len(seeds); i++ {
		seeds[i] = mustParseInt(seedsStrs[i+1])
	}

	allMaps := make([][]Map, catCount)
	for i := range allMaps {
		allMaps[i] = getCatMaps(seedsAndCatStrs[i+1])
	}

	return seeds, allMaps
}

func getCatMaps(catStr string) []Map {
	catParts := strings.Split(strings.Trim(catStr, "\r\n"), ":\r\n")
	lines := strings.Split(catParts[1], "\r\n")
	catMaps := make([]Map, len(lines))
	for i, line := range lines {
		numStrs := strings.Split(line, " ")
		destStart := mustParseInt(numStrs[0])
		srcStart := mustParseInt(numStrs[1])
		length := mustParseInt(numStrs[2])
		catMaps[i] = Map{
			destStart: destStart,
			destEnd:   destStart + length - 1,
			srcStart:  srcStart,
			srcEnd:    srcStart + length - 1,
			length:    length,
		}
	}
	return catMaps
}

func getLoc(seedId int, allMaps [][]Map) int {
	id := seedId
	for i := range allMaps {
		id = getCatValue(id, allMaps[i])
	}
	return id
}

func getCatValue(catKey int, catMaps []Map) int {
	for _, catMap := range catMaps {
		destStart := catMap.destStart
		srcStart := catMap.srcStart
		length := catMap.length
		if catKey >= srcStart && catKey < srcStart+length {
			diff := destStart - srcStart
			return catKey + diff
		}
	}
	return catKey
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
	lowestLocNum := -1
	startCats := make([]int, len(seedRanges)/2)

seedLoop:
	for len(seedRanges) > 0 {
		start, length := seedRanges[0], seedRanges[1]
		end := start + length - 1
		seedRanges = seedRanges[2:]

		startCat := startCats[0]
		startCats = startCats[1:]

		for i := startCat; i < len(allMaps); i++ {
			catMaps := allMaps[i]
			for _, catMap := range catMaps {
				if start >= catMap.srcStart && end <= catMap.srcEnd {
					// The seed range is entirely within the category map's range.
					start = start - catMap.srcStart + catMap.destStart
					end = end - catMap.srcStart + catMap.destStart
					break // go to the next category
				} else if start < catMap.srcStart && end >= catMap.srcStart && end <= catMap.srcEnd {
					// The seed range is partially within the category map's range,
					// extending further to the left (and maybe also the right).
					// Split the seed range into two smaller ranges, one of which is
					// entirely within the category map's range.
					newSeed1Start := start
					newSeed1Length := catMap.srcStart - start
					newSeed2Start := catMap.srcStart
					newSeed2Length := end - catMap.srcStart + 1
					seedRanges = append(
						seedRanges,
						newSeed1Start,
						newSeed1Length,
						newSeed2Start,
						newSeed2Length,
					)
					startCats = append(startCats, i, i)
					continue seedLoop
				} else if start >= catMap.srcStart && start <= catMap.srcEnd && end > catMap.srcEnd {
					// The seed range is partially within the category map's range,
					// extending further to the right.
					// Split the seed range into two smaller ranges, one of which is
					// entirely within the category map's range.
					newSeed1Start := start
					newSeed1Length := catMap.srcEnd - start + 1
					newSeed2Start := catMap.srcEnd + 1
					newSeed2Length := end - catMap.srcEnd
					seedRanges = append(
						seedRanges,
						newSeed1Start,
						newSeed1Length,
						newSeed2Start,
						newSeed2Length,
					)
					startCats = append(startCats, i, i)
					continue seedLoop
				}
			}
		}

		if lowestLocNum == -1 || start < lowestLocNum {
			lowestLocNum = start
		}
	}

	fmt.Println("part 2 result:", lowestLocNum)
}
