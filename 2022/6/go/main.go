// https://adventofcode.com/2022/day/6
package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
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
	line := string(contentBytes)
	if strings.Contains(line, "\n") {
		panic("There's more than one line")
	}

	last4Letters := make([]string, 4)
	i := 0

	letterCount := 0
	for _, letter := range line {
		last4Letters[i] = string(letter)
		i = (i + 1) % 4
		letterCount++
		if isPacketStartMarker(last4Letters) {
			break
		}
	}
	if letterCount != 1848 {
		panic(letterCount)
	}
	fmt.Println(letterCount)
}

func isPacketStartMarker(fourLetters []string) bool {
	strs := make([]string, 4)
	copy(strs, fourLetters)
	sort.Strings(strs)
	return strs[0] != strs[1] &&
		strs[1] != strs[2] &&
		strs[2] != strs[3] &&
		!slices.Contains(strs, "")
}

func isMessageStartMarker(fourteenLetters []string) bool {
	strs := make([]string, 14)
	copy(strs, fourteenLetters)
	if slices.Contains(strs, "") {
		return false
	}
	sort.Strings(strs)
	for i := 0; i < len(strs)-1; i++ {
		if strs[i] == strs[i+1] {
			return false
		}
	}
	return true
}

func part2() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	line := string(contentBytes)
	if strings.Contains(line, "\n") {
		panic("There's more than one line")
	}

	last14Letters := make([]string, 14)
	i := 0

	letterCount := 0
	for _, letter := range line {
		last14Letters[i] = string(letter)
		i = (i + 1) % 14
		letterCount++
		if isMessageStartMarker(last14Letters) {
			break
		}
	}
	if letterCount != 2308 {
		panic(letterCount)
	}
	fmt.Println(letterCount)
}
