// https://adventofcode.com/2022/day/8
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.Trim(string(contentBytes), "\r\n")
	lines := strings.Split(content, "\r\n")

	heights := make([][]int, len(lines))
	isVisible := make([][]bool, len(lines))
	for i, line := range lines {
		heights[i] = make([]int, len(line))
		isVisible[i] = make([]bool, len(line))
		for j, ch := range line {
			n, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			heights[i][j] = n
		}
	}

	part1(heights, isVisible)
	part2(heights)
}

func part1(heights [][]int, isVisible [][]bool) {
	// right
	for y := 0; y < len(heights); y++ {
		highestSoFar := -1
		for x := 0; x < len(heights[y]); x++ {
			height := heights[y][x]
			if x == 0 {
				isVisible[y][x] = true
				highestSoFar = height
			} else if height > highestSoFar {
				isVisible[y][x] = true
				highestSoFar = height
			}
		}
	}

	// left
	for y := 0; y < len(heights); y++ {
		highestSoFar := -1
		for x := len(heights[y]) - 1; x >= 0; x-- {
			height := heights[y][x]
			if x == len(heights[y])-1 {
				isVisible[y][x] = true
				highestSoFar = height
			} else if height > highestSoFar {
				isVisible[y][x] = true
				highestSoFar = height
			}
		}
	}

	// down
	for x := 0; x < len(heights[0]); x++ {
		highestSoFar := -1
		for y := 0; y < len(heights); y++ {
			height := heights[y][x]
			if y == 0 {
				isVisible[y][x] = true
				highestSoFar = height
			} else if height > highestSoFar {
				isVisible[y][x] = true
				highestSoFar = height
			}
		}
	}

	// up
	for x := 0; x < len(heights[0]); x++ {
		highestSoFar := -1
		for y := len(heights) - 1; y >= 0; y-- {
			height := heights[y][x]
			if y == len(heights)-1 {
				isVisible[y][x] = true
				highestSoFar = height
			} else if height > highestSoFar {
				isVisible[y][x] = true
				highestSoFar = height
			}
		}
	}

	// fmt.Println(heights)
	// fmt.Println(isVisible)

	sum := 0
	for _, row := range isVisible {
		for _, e := range row {
			if e {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func part2(heights [][]int) {
	maxScore := 0
	for y := 1; y < len(heights)-1; y++ {
		for x := 1; x < len(heights[0])-1; x++ {
			score := getScenicScore(x, y, heights)
			// fmt.Printf("... (%v, %v) score: %v\n", x, y, score)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)
}

func getScenicScore(x int, y int, heights [][]int) int {
	height := heights[y][x]

	down := 0
	for i := y + 1; i < len(heights); i++ {
		down++
		if heights[i][x] >= height {
			break
		}
	}

	up := 0
	for i := y - 1; i >= 0; i-- {
		up++
		if heights[i][x] >= height {
			break
		}
	}

	right := 0
	for i := x + 1; i < len(heights[0]); i++ {
		right++
		if heights[y][i] >= height {
			break
		}
	}

	left := 0
	for i := x - 1; i >= 0; i-- {
		left++
		if heights[y][i] >= height {
			break
		}
	}

	return down * up * right * left
}
