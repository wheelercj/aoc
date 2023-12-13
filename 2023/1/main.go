package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var numWords []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	text := string(bytes)
	// fmt.Println("text:", text)
	lines := strings.Split(text, "\r\n")

	sum1 := part1(lines)
	fmt.Printf("part 1 result: %d\n", sum1)

	sum2 := part2(lines)
	fmt.Printf("part 2 result: %d\n", sum2)
}

func part1(lines []string) int {
	var sum int

	for _, line := range lines {
		var tensDigit rune
		var onesDigit rune

		for _, ch := range line {
			if ch >= '0' && ch <= '9' {
				tensDigit = ch
				// fmt.Println("tens digit string:", string(tensDigit))
				// fmt.Println("tens digit rune:", tensDigit)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			ch := rune(line[i])
			if ch >= '0' && ch <= '9' {
				onesDigit = ch
				// fmt.Println("ones digit string:", string(onesDigit))
				// fmt.Println("ones digit rune:", onesDigit)
				break
			}
		}

		calibrationValue := MustParseInt(tensDigit)*10 + MustParseInt(onesDigit)
		// fmt.Printf("line %d calibration value: %d\n", j, calibrationValue)
		sum += calibrationValue
	}

	return sum
}

func part2(lines []string) int {
	var sum int

	for _, line := range lines {
		var tensDigit int
		var onesDigit int

		for i, ch := range line {
			numWord := numWordStartsAt(i, line)
			if ch >= '0' && ch <= '9' {
				tensDigit = MustParseInt(ch)
				break
			} else if numWord != "" {
				// fmt.Println(j+1, "tens digit numWord:", numWord)
				tensDigit = slices.Index(numWords, numWord) + 1
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			numWord := numWordEndsAt(i, line)
			ch := rune(line[i])
			if ch >= '0' && ch <= '9' {
				onesDigit = MustParseInt(ch)
				break
			} else if numWord != "" {
				// fmt.Println(j+1, "ones digit numWord:", numWord)
				onesDigit = slices.Index(numWords, numWord) + 1
				break
			}
		}

		calibrationValue := tensDigit*10 + onesDigit
		// fmt.Printf("line %d calibration value: %d\n", j, calibrationValue)
		sum += calibrationValue
	}

	return sum
}

func MustParseInt(num rune) int {
	if num == '\x00' {
		num = '0'
	}
	result, err := strconv.ParseInt(string(num), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(result)
}

func numWordStartsAt(i int, line string) string {
	for _, numWord := range numWords {
		if i+len(numWord) < len(line) && line[i:i+len(numWord)] == numWord {
			return numWord
		}
	}
	return ""
}

func numWordEndsAt(i int, line string) string {
	for _, numWord := range numWords {
		if i-len(numWord)+1 >= 0 && line[i-len(numWord)+1:i+1] == numWord {
			return numWord
		}
	}
	return ""
}
