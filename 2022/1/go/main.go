// https://adventofcode.com/2022/day/1
package main

import (
	"fmt"
	"os"
	"sort"
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
	content := string(contentBytes)

	elvesStrs := strings.Split(content, "\r\n\r\n")
	elves := make([][]int, len(elvesStrs))
	for i, elfStr := range elvesStrs {
		itemStrs := strings.Split(elfStr, "\r\n")
		elves[i] = make([]int, len(itemStrs))
		for j, itemStr := range itemStrs {
			itemCalories, err := strconv.Atoi(itemStr)
			if err != nil {
				panic(err)
			}
			elves[i][j] = itemCalories
		}
	}

	elfSums := make([]int, len(elves))
	for i, elf := range elves {
		for _, itemCalories := range elf {
			elfSums[i] += itemCalories
		}
	}
	sort.Ints(elfSums)
	result := elfSums[len(elfSums)-1]
	if result != 69836 {
		panic(result)
	}
	fmt.Println(result)
}

func part2() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)

	elvesStrs := strings.Split(content, "\r\n\r\n")
	elves := make([][]int, len(elvesStrs))
	for i, elfStr := range elvesStrs {
		itemStrs := strings.Split(elfStr, "\r\n")
		elves[i] = make([]int, len(itemStrs))
		for j, itemStr := range itemStrs {
			itemCalories, err := strconv.Atoi(itemStr)
			if err != nil {
				panic(err)
			}
			elves[i][j] = itemCalories
		}
	}

	elfSums := make([]int, len(elves))
	for i, elf := range elves {
		for _, itemCalories := range elf {
			elfSums[i] += itemCalories
		}
	}
	sort.Ints(elfSums)
	top3 := elfSums[len(elfSums)-3:]
	result := top3[0] + top3[1] + top3[2]
	if result != 207968 {
		panic(result)
	}
	fmt.Println(result)
}
