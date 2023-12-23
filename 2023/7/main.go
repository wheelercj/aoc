package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Part int

const (
	isPart1 Part = iota
	isPart2
)

type Type int

const (
	highCard Type = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

// handTypeNames is the print-friendly names of hand types. This list should be in the
// same order as the "enum" above.
var handTypeNames []string = []string{
	"high card",
	"one pair",
	"two pair",
	"three of a kind",
	"full house",
	"four of a kind",
	"five of a kind",
}

type Hand struct {
	hand  string
	bid   int
	type_ Type
}

// compareHands returns a function that returns -1 if a < b, 0 if a == b, and 1 if a >
// b.
func compareHands(part Part) func(a, b Hand) int {
	return func(a, b Hand) int {
		if a.type_ < b.type_ {
			return -1
		}
		if a.type_ > b.type_ {
			return 1
		}
		for i := 0; i < len(a.hand); i++ {
			if a.hand[i] == b.hand[i] {
				continue
			}
			return compareLabels(a.hand[i], b.hand[i], part)
		}
		return 0
	}
}

// compareLabels returns -1 if a < b, 0 if a == b, and 1 if a > b.
func compareLabels(a, b byte, part Part) int {
	var labels string
	if part == isPart1 {
		labels = "23456789TJQKA"
	} else {
		labels = "J23456789TQKA"
	}
	ai := strings.Index(labels, string(a))
	bi := strings.Index(labels, string(b))
	if ai < bi {
		return -1
	} else if ai == bi {
		return 0
	} else {
		return 1
	}
}

func main() {
	hands := parseInput("input.txt")
	part1(hands)
	part2(hands)
}

func parseInput(fileName string) []Hand {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")

	hands := make([]Hand, len(lines))
	for i, line := range lines {
		handAndBid := strings.Split(line, " ")
		hand := handAndBid[0]
		bid := mustParseInt(handAndBid[1])
		hands[i] = Hand{hand: hand, bid: bid}
	}

	return hands
}

func mustParseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func getType1(hand string) Type {
	handSlice := strings.Split(hand, "")
	slices.Sort(handSlice)
	unique := slices.Compact(handSlice)
	switch len(unique) {
	case 1:
		return fiveOfAKind
	case 2:
		switch strings.Count(hand, handSlice[0]) {
		case 1, 4:
			return fourOfAKind
		case 2, 3:
			return fullHouse
		}
	case 3:
		for _, ch := range unique {
			switch strings.Count(hand, ch) {
			case 3:
				return threeOfAKind
			case 2:
				return twoPair
			case 1:
				continue
			}
		}
	case 4:
		return onePair
	case 5:
		return highCard
	}

	panic("unreachable")
}

func getType2(hand string) Type {
	handSlice := strings.Split(hand, "")
	slices.Sort(handSlice)
	unique := slices.Compact(handSlice)
	switch len(unique) {
	case 1: // 1 unique label
		return fiveOfAKind // AAAAA
	case 2: // 2 unique labels
		if strings.Contains(hand, "J") {
			return fiveOfAKind // AAAAJ
		}
		switch strings.Count(hand, handSlice[0]) {
		case 1, 4:
			return fourOfAKind // AAAAQ
		case 2, 3:
			return fullHouse // AAAQQ
		}
	case 3: // 3 unique labels
		for _, ch := range unique {
			switch strings.Count(hand, ch) {
			case 3: // 3 instances of one label, the other two labels are unique
				if strings.Contains(hand, "J") {
					return fourOfAKind // AJJJQ, AAAJQ
				}
				return threeOfAKind // A999Q
			case 2: // 2 instances of one label, so there must be two pairs
				switch strings.Count(hand, "J") {
				case 2:
					return fourOfAKind // AAJJQ
				case 1:
					return fullHouse // AAJQQ
				}
				return twoPair // AA9QQ
			case 1: // unknown, keep searching
				continue
			}
		}
	case 4: // 4 unique labels
		if strings.Contains(hand, "J") {
			return threeOfAKind
		}
		return onePair
	case 5: // 5 unique labels
		if strings.Contains(hand, "J") {
			return onePair
		}
		return highCard
	}

	panic("unreachable")
}

func printHands(hands []Hand) {
	fmt.Println("\thand\tbid\ttype")
	for i, hand := range hands {
		fmt.Print(i)
		printHand(hand)
	}
}

func printHand(hand Hand) {
	fmt.Printf("\t%v\t%d\t%v\n", hand.hand, hand.bid, handTypeNames[hand.type_])
}

func part1(hands []Hand) {
	for i, hand := range hands {
		hands[i].type_ = getType1(hand.hand)
	}
	slices.SortFunc(hands, compareHands(isPart1))
	// printHands(hands)

	var sum int
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	fmt.Println("part 1 result:", sum)
}

func part2(hands []Hand) {
	for i, hand := range hands {
		hands[i].type_ = getType2(hand.hand)
	}
	slices.SortFunc(hands, compareHands(isPart2))
	// printHands(hands)

	var sum int
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	fmt.Println("part 2 result:", sum)
}
