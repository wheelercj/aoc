// https://adventofcode.com/2022/day/3
package main

import (
	"fmt"
	"os"
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

	var totalPriority int
	for _, line := range lines {
		left, right := line[:len(line)/2], line[len(line)/2:]
		var misplaced rune
		for _, ch := range left {
			if strings.Contains(right, string(ch)) {
				misplaced = ch
				break
			}
		}
		totalPriority += getPriority(misplaced)
	}

	if totalPriority != 7727 {
		panic(totalPriority)
	}
	fmt.Println(totalPriority)
}

func getPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 'a' + 1
	} else {
		return int(item) - 'A' + 27
	}
}

func part2() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contentBytes), "\r\n")
	groups := make([][]string, len(lines)/3)
	groupNum := 0
	for i := 0; i < len(lines); i += 3 {
		groups[groupNum] = []string{
			lines[i],
			lines[i+1],
			lines[i+2],
		}
		groupNum++
	}

	var totalPriority int
	for _, group := range groups {
		elf1, elf2, elf3 := string(group[0]), string(group[1]), string(group[2])

		var badge rune
		for _, ch := range elf1 {
			if strings.Contains(elf2, string(ch)) && strings.Contains(elf3, string(ch)) {
				badge = ch
				break
			}
		}
		totalPriority += getPriority(badge)
	}

	if totalPriority != 2609 {
		panic(totalPriority)
	}
	fmt.Println(totalPriority)
}
