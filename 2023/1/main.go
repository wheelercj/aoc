package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numWords []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	bytes, err := os.ReadFile("sample2.txt")
	if err != nil {
		panic(err)
	}
	text := string(bytes)
	// fmt.Println("text:", text)
	lines := strings.Split(text, "\r\n")

	sum1 := part1(lines)
	fmt.Printf("part 1 result: %d\n", sum1)
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
