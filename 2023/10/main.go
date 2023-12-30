package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func parseInput(fileName string) [][]string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	return grid
}

func findS(grid [][]string) Coord {
	for y, line := range grid {
		for x, ch := range line {
			if ch == "S" {
				return Coord{x: x, y: y}
			}
		}
	}
	panic("unreachable")
}

func findLoop(sCoords Coord, grid [][]string) []Coord {
	loop := []Coord{sCoords}
	nextLoopPipeCoords := findNextLoopPipe(sCoords, nil, grid)
	prevLoopPipe := sCoords
	for nextLoopPipeCoords.x != sCoords.x || nextLoopPipeCoords.y != sCoords.y {
		loop = append(loop, nextLoopPipeCoords)
		tempLoopPipe := nextLoopPipeCoords
		nextLoopPipeCoords = findNextLoopPipe(nextLoopPipeCoords, &prevLoopPipe, grid)
		prevLoopPipe = tempLoopPipe
	}
	return loop
}

func findNextLoopPipe(currentCoords Coord, prevCoords *Coord, grid [][]string) Coord {
	currentPipe := grid[currentCoords.y][currentCoords.x]
	if currentPipe == "S" {
		currentPipe = getSPipe(currentCoords, grid)
	}
	switch currentPipe {
	case "|":
		if loopGoesUp(currentCoords, prevCoords, grid) {
			return Coord{x: currentCoords.x, y: currentCoords.y - 1}
		}
		return Coord{x: currentCoords.x, y: currentCoords.y + 1}
	case "-":
		if loopGoesLeft(currentCoords, prevCoords, grid) {
			return Coord{x: currentCoords.x - 1, y: currentCoords.y}
		}
		return Coord{x: currentCoords.x + 1, y: currentCoords.y}
	case "L":
		if loopGoesUp(currentCoords, prevCoords, grid) {
			return Coord{x: currentCoords.x, y: currentCoords.y - 1}
		}
		return Coord{x: currentCoords.x + 1, y: currentCoords.y}
	case "J":
		if loopGoesUp(currentCoords, prevCoords, grid) {
			return Coord{x: currentCoords.x, y: currentCoords.y - 1}
		}
		return Coord{x: currentCoords.x - 1, y: currentCoords.y}
	case "7":
		if loopGoesLeft(currentCoords, prevCoords, grid) {
			return Coord{x: currentCoords.x - 1, y: currentCoords.y}
		}
		return Coord{x: currentCoords.x, y: currentCoords.y + 1}
	case "F":
		if loopGoesRight(currentCoords, prevCoords, grid) {
			return Coord{x: currentCoords.x + 1, y: currentCoords.y}
		}
		return Coord{x: currentCoords.x, y: currentCoords.y + 1}
	}
	panic("unreachable")
}

func loopGoesUp(currentCoords Coord, prevCoords *Coord, grid [][]string) bool {
	upCoords := currentCoords
	upCoords.y--
	if upCoords.y >= 0 {
		if prevCoords == nil || *prevCoords != upCoords {
			upPipe := grid[upCoords.y][upCoords.x]
			if slices.Contains([]string{"|", "7", "F", "S"}, upPipe) {
				return true
			}
		}
	}
	return false
}

func loopGoesRight(currentCoords Coord, prevCoords *Coord, grid [][]string) bool {
	rightCoords := currentCoords
	rightCoords.x++
	if rightCoords.x < len(grid[0]) {
		if prevCoords == nil || *prevCoords != rightCoords {
			rightPipe := grid[rightCoords.y][rightCoords.x]
			if slices.Contains([]string{"-", "J", "7", "S"}, rightPipe) {
				return true
			}
		}
	}
	return false
}

func loopGoesDown(currentCoords Coord, prevCoords *Coord, grid [][]string) bool {
	downCoords := currentCoords
	downCoords.y++
	if downCoords.y < len(grid) {
		if prevCoords == nil || *prevCoords != downCoords {
			downPipe := grid[downCoords.y][downCoords.x]
			if slices.Contains([]string{"|", "L", "J", "S"}, downPipe) {
				return true
			}
		}
	}
	return false
}

func loopGoesLeft(currentCoords Coord, prevCoords *Coord, grid [][]string) bool {
	leftCoords := currentCoords
	leftCoords.x--
	if leftCoords.x >= 0 {
		if prevCoords == nil || *prevCoords != leftCoords {
			leftPipe := grid[leftCoords.y][leftCoords.x]
			if slices.Contains([]string{"-", "L", "F", "S"}, leftPipe) {
				return true
			}
		}
	}
	return false
}

// findContainedArea finds the number of tiles loop contains in grid. It does this by
// iterating through each line of the grid and counting the number of times a
// perpendicular pipe within loop is crossed. Crossing corner pipes in loop counts if
// they double back on the previous corner pipe in loop or if there is no previous
// corner pipe. Tiles with an odd loop crossing count are within the contained area.
func findContainedArea(loop []Coord, grid [][]string, sCoords Coord) int {
	var area int
	cornerPipes := []string{"L", "J", "7", "F"}
	for y, line := range grid {
		var loopCrossCount int
		var prevCorner string
		for x, ch := range line {
			if ch == "S" {
				ch = getSPipe(sCoords, grid)
			}
			partOfLoop := slices.Contains(loop, Coord{x: x, y: y})
			if !partOfLoop {
				insideLoop := loopCrossCount%2 == 1
				if insideLoop {
					area++
				}
			} else if ch != "-" {
				if ch == "|" {
					loopCrossCount++
				} else if slices.Contains(cornerPipes, ch) {
					switch prevCorner {
					case "":
						loopCrossCount++
						prevCorner = ch
					case "L":
						if ch == "J" {
							loopCrossCount++
						}
						prevCorner = ""
					case "F":
						if ch == "7" {
							loopCrossCount++
						}
						prevCorner = ""
					}
				}
			}
		}
	}
	return area
}

// getSPipe determines what kind of pipe the S symbol is covering up in grid.
func getSPipe(sCoords Coord, grid [][]string) string {
	var up, right, down, left bool
	if loopGoesUp(sCoords, nil, grid) {
		up = true
	}
	if loopGoesRight(sCoords, nil, grid) {
		right = true
	}
	if loopGoesDown(sCoords, nil, grid) {
		down = true
	}
	if loopGoesLeft(sCoords, nil, grid) {
		left = true
	}
	if up && right {
		return "L"
	} else if up && down {
		return "|"
	} else if up && left {
		return "J"
	} else if right && down {
		return "F"
	} else if right && left {
		return "-"
	} else {
		return "7"
	}
}

func part1(grid [][]string) {
	sCoords := findS(grid)
	loop := findLoop(sCoords, grid)
	fmt.Println("part 1 result:", len(loop)/2)
}

func part2(grid [][]string) {
	sCoords := findS(grid)
	loop := findLoop(sCoords, grid)
	area := findContainedArea(loop, grid, sCoords)
	fmt.Println("part 2 result:", area)
}
