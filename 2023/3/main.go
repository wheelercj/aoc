package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	input := parseInput("input.txt")
	part1(input) // < 632038
}

func parseInput(fileName string) [][]string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	text := string(bytes)
	lines := strings.Split(text, "\r\n")

	input := make([][]string, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		input[i] = make([]string, len(line))
		for j, ch := range line {
			input[i][j] = string(ch)
		}
	}

	return input
}

func part1(input [][]string) {
	var sum int
	for i, line := range input {
		for j, ch := range line {
			if isSpecialChar(ch) {
				nums := getAdjacentNums(j, i, input)
				for _, num := range nums {
					sum += num
				}
			}
		}
	}
	fmt.Println("sum:", sum)
}

func isNum(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func isSpecialChar(ch string) bool {
	return !isNum(ch) && ch != "." && ch != " "
}

// getAdjacentNums gets all the numbers next to (including diagonally) the chosen
// coordinates in a 2-D array of strings. x and y are expected to be the coordinates of
// a special character. No number in input is added twice, even if it has multiple
// special characters next to it.
func getAdjacentNums(x, y int, input [][]string) []int {
	var nums []int
	var foundNumCoords []Coord

	if x-1 >= 0 && isNum(input[y][x-1]) { // left
		num, coords := getNum(x-1, y, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	}
	if x+1 < len(input[y]) && isNum(input[y][x+1]) { // right
		num, coords := getNum(x+1, y, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	}

	var topMidIsNum bool
	if x-1 >= 0 && y-1 >= 0 && isNum(input[y-1][x-1]) { // top left
		num, coords := getNum(x-1, y-1, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	} else if y-1 >= 0 && isNum(input[y-1][x]) { // top middle
		topMidIsNum = true
		num, coords := getNum(x, y-1, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	}
	if !topMidIsNum && y-1 >= 0 && x+1 < len(input[y-1]) && isNum(input[y-1][x+1]) { // top right
		num, coords := getNum(x+1, y-1, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	}

	var bottomMidIsNum bool
	if x-1 >= 0 && y+1 < len(input) && isNum(input[y+1][x-1]) { // bottom left
		num, coords := getNum(x-1, y+1, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	} else if y+1 < len(input) && isNum(input[y+1][x]) { // bottom middle
		bottomMidIsNum = true
		num, coords := getNum(x, y+1, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	}
	if !bottomMidIsNum && y+1 < len(input) && x+1 < len(input[y+1]) && isNum(input[y+1][x+1]) { // bottom right
		num, coords := getNum(x+1, y+1, input)
		if !slices.Contains(foundNumCoords, coords) {
			nums = append(nums, num)
			foundNumCoords = append(foundNumCoords, coords)
		}
	}

	return nums
}

// getNum gets all consecutive numeric digits that make up one number from a 2-D array
// of strings, and returns the integer of that number and the coordinates of its
// leftmost digit. The x and y inputs are coordinates for any one part of the number.
func getNum(x, y int, input [][]string) (int, Coord) {
	start := x
	end := x

	// left
	for i := x - 1; i >= 0 && isNum(input[y][i]); i-- {
		start = i
	}
	// right
	for i := x + 1; i < len(input[y]) && isNum(input[y][i]); i++ {
		end = i
	}

	numStr := strings.Join(input[y][start:end+1], "")
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(num), Coord{start, y}
}

func part2() {

}
