package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	winningNums []int
	yourNums    []int
	matchCount  int
}

func main() {
	input := parseInput("input.txt")
	part1(input)
	part2(input)
}

func parseInput(fileName string) []Card {
	var cards []Card
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\r\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineParts := strings.Split(line, ": ")
		numbersStr := lineParts[1]
		numbersParts := strings.Split(numbersStr, " | ")
		winningNums, yourNums := getNums(numbersParts[0]), getNums(numbersParts[1])

		cards = append(
			cards,
			Card{winningNums: winningNums, yourNums: yourNums},
		)
	}

	return cards
}

func getNums(numsStr string) []int {
	var nums []int
	numsStr = strings.Trim(numsStr, " ")
	numsStrs := strings.Split(numsStr, " ")
	for _, numStr := range numsStrs {
		numStr = strings.Trim(numStr, " ")
		if len(numStr) == 0 {
			continue
		}
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(num))
	}
	return nums
}

func part1(cards []Card) {
	var totalPoints int
	for _, card := range cards {
		var cardPoints int
		for _, winningNum := range card.winningNums {
			if slices.Contains(card.yourNums, winningNum) {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}

		totalPoints += cardPoints
	}

	fmt.Println("points:", totalPoints)
}

func part2(cards []Card) {
	addMatchCounts(cards)

	cardCopyIds := make([]int, len(cards))
	for i := range cardCopyIds {
		cardCopyIds[i] = i + 1
	}

	for i := 0; i < len(cardCopyIds); i++ {
		cardCopyId := cardCopyIds[i]
		start := cardCopyId + 1
		end := cardCopyId + 1 + cards[cardCopyId-1].matchCount
		for j := start; j < end; j++ {
			cardCopyIds = append(cardCopyIds, j)
		}
	}

	fmt.Println("card count:", len(cardCopyIds))
}

func addMatchCounts(cards []Card) {
	for i, card := range cards {
		for _, winningNum := range card.winningNums {
			if slices.Contains(card.yourNums, winningNum) {
				cards[i].matchCount++
			}
		}
	}
}
