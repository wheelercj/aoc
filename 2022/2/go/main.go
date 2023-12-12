// https://adventofcode.com/2022/day/2
package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota
	paper
	scissors
)

const (
	rockPoints     = 1
	paperPoints    = 2
	scissorsPoints = 3
)

const (
	winPoints  = 6
	drawPoints = 3
	losePoints = 0
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
	lines := strings.Split(content, "\r\n")

	score := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		elfChoice, myChoice := toEnum(choices[0]), toEnum(choices[1])
		if iWin(elfChoice, myChoice) {
			score += winPoints
		} else if elfChoice == myChoice {
			score += drawPoints
		} else {
			score += losePoints
		}

		switch myChoice {
		case rock:
			score += rockPoints
		case paper:
			score += paperPoints
		case scissors:
			score += scissorsPoints
		}
	}

	if score != 13009 {
		panic(score)
	}
	fmt.Println(score)
}

func toEnum(choice string) int {
	if choice == "A" || choice == "X" {
		return rock
	} else if choice == "B" || choice == "Y" {
		return paper
	} else {
		return scissors
	}
}

func playerOutcomeToEnum(elfChoiceStr, myOutcomeStr string) int {
	switch elfChoiceStr {
	case "A":  // rock
		switch myOutcomeStr {
		case "X":  // lose
			return scissors
		case "Y":  // draw
			return rock
		case "Z":  // win
			return paper
		}
	case "B":  // paper
		switch myOutcomeStr {
		case "X":  // lose
			return rock
		case "Y":  // draw
			return paper
		case "Z":  // win
			return scissors
		}
	case "C":  // scissors
		switch myOutcomeStr {
		case "X":  // lose
			return paper
		case "Y":  // draw
			return scissors
		case "Z":  // win
			return rock
		}
	}
	panic("Invalid input")
}

func iWin(elfChoice, myChoice int) bool {
	return elfChoice == rock && myChoice == paper || elfChoice == paper && myChoice == scissors || elfChoice == scissors && myChoice == rock
}

func part2() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)
	lines := strings.Split(content, "\r\n")

	score := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		elfChoice, myChoice := toEnum(choices[0]), playerOutcomeToEnum(choices[0], choices[1])
		if iWin(elfChoice, myChoice) {
			score += winPoints
		} else if elfChoice == myChoice {
			score += drawPoints
		} else {
			score += losePoints
		}

		switch myChoice {
		case rock:
			score += rockPoints
		case paper:
			score += paperPoints
		case scissors:
			score += scissorsPoints
		}
	}

	if score != 10398 {
		panic(score)
	}
	fmt.Println(score)
}
