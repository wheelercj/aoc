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

type Direction int

const (
	noDirection Direction = iota
	up
	right
	down
	left
)

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

// getSPipe determines what kind of pipe the S symbol is covering up in grid.
func getSPipe(sCoords Coord, grid [][]string) string {
	var (
		up    = loopGoesUp(sCoords, grid)
		right = loopGoesRight(sCoords, grid)
		down  = loopGoesDown(sCoords, grid)
		left  = loopGoesLeft(sCoords, grid)
	)
	switch {
	case up && right:
		return "L"
	case up && down:
		return "|"
	case up && left:
		return "J"
	case right && down:
		return "F"
	case right && left:
		return "-"
	case down && left:
		return "7"
	}
	panic("unreachable")
}

func findLoop(sCoords Coord, grid [][]string) []Coord {
	loop := []Coord{sCoords}
	coords := sCoords
	dir := noDirection
	for {
		dir = findNextLoopDir(coords, dir, grid)
		switch dir {
		case up:
			coords.y--
		case right:
			coords.x++
		case down:
			coords.y++
		case left:
			coords.x--
		}
		if coords == sCoords {
			break
		}
		loop = append(loop, coords)
	}
	return loop
}

// findNextLoopDir finds the direction of the next part of the loop. prevCoords is the
// coordinates of the most recently found part of the loop. prevDir is the direction in
// which prevCoords was found.
func findNextLoopDir(prevCoords Coord, prevDir Direction, grid [][]string) Direction {
	prevPipe := grid[prevCoords.y][prevCoords.x]
	if prevPipe == "S" {
		prevPipe = getSPipe(prevCoords, grid)
	}
	switch prevPipe {
	case "|":
		if prevDir == up {
			return up
		}
		return down
	case "-":
		if prevDir == left {
			return left
		}
		return right
	case "L":
		if prevDir == down {
			return right
		}
		return up
	case "J":
		if prevDir == right {
			return up
		}
		return left
	case "7":
		if prevDir == up {
			return left
		}
		return down
	case "F":
		if prevDir == up {
			return right
		}
		return down
	}
	panic("unreachable")
}

func loopGoesUp(currentCoords Coord, grid [][]string) bool {
	upCoords := currentCoords
	upCoords.y--
	if upCoords.y >= 0 {
		upPipe := grid[upCoords.y][upCoords.x]
		if slices.Contains([]string{"|", "7", "F", "S"}, upPipe) {
			return true
		}
	}
	return false
}

func loopGoesRight(currentCoords Coord, grid [][]string) bool {
	rightCoords := currentCoords
	rightCoords.x++
	if rightCoords.x < len(grid[0]) {
		rightPipe := grid[rightCoords.y][rightCoords.x]
		if slices.Contains([]string{"-", "J", "7", "S"}, rightPipe) {
			return true
		}
	}
	return false
}

func loopGoesDown(currentCoords Coord, grid [][]string) bool {
	downCoords := currentCoords
	downCoords.y++
	if downCoords.y < len(grid) {
		downPipe := grid[downCoords.y][downCoords.x]
		if slices.Contains([]string{"|", "L", "J", "S"}, downPipe) {
			return true
		}
	}
	return false
}

func loopGoesLeft(currentCoords Coord, grid [][]string) bool {
	leftCoords := currentCoords
	leftCoords.x--
	if leftCoords.x >= 0 {
		leftPipe := grid[leftCoords.y][leftCoords.x]
		if slices.Contains([]string{"-", "L", "F", "S"}, leftPipe) {
			return true
		}
	}
	return false
}

// findContainedArea finds the number of tiles loop contains in grid. It does this by
// iterating through each line of the grid and counting the number of times a
// perpendicular pipe within loop is crossed. Tiles with an odd loop crossing count are
// within the contained area. Crossing corner pipes in loop counts if they double back
// on the previous corner pipe in loop or if there is no previous corner pipe.
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
