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
	for i, line := range grid {
		for j, ch := range line {
			if ch == "S" {
				return Coord{x: i, y: j}
			}
		}
	}
	panic("unreachable")
}

func findLoop(sCoords Coord, grid [][]string) []Coord {
	loop := []Coord{sCoords}
	nextPipe := findNextPipe(sCoords, nil, grid)
	prevPipe := sCoords
	for nextPipe.x != sCoords.x || nextPipe.y != sCoords.y {
		loop = append(loop, nextPipe)
		tempPipe := nextPipe
		nextPipe = findNextPipe(nextPipe, &prevPipe, grid)
		prevPipe = tempPipe
	}
	return loop
}

func findNextPipe(currentCoords Coord, prevCoords *Coord, grid [][]string) Coord {
	currentPipe := grid[currentCoords.y][currentCoords.x]
	switch currentPipe {
	case "|":
		if nextPipe := lookUp(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookDown(currentCoords, prevCoords, grid)
	case "-":
		if nextPipe := lookLeft(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookRight(currentCoords, prevCoords, grid)
	case "L":
		if nextPipe := lookUp(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookRight(currentCoords, prevCoords, grid)
	case "J":
		if nextPipe := lookUp(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookLeft(currentCoords, prevCoords, grid)
	case "7":
		if nextPipe := lookLeft(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookDown(currentCoords, prevCoords, grid)
	case "F":
		if nextPipe := lookRight(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookDown(currentCoords, prevCoords, grid)
	case "S":
		if nextPipe := lookUp(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		if nextPipe := lookRight(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		if nextPipe := lookDown(currentCoords, prevCoords, grid); nextPipe != nil {
			return *nextPipe
		}
		return *lookLeft(currentCoords, prevCoords, grid)
	}
	panic("unreachable")
}

func lookUp(currentCoords Coord, prevCoords *Coord, grid [][]string) *Coord {
	if currentCoords.y-1 >= 0 {
		if prevCoords == nil || prevCoords.y != currentCoords.y-1 || prevCoords.x != currentCoords.x {
			if slices.Contains([]string{"|", "7", "F", "S"}, grid[currentCoords.y-1][currentCoords.x]) {
				return &Coord{x: currentCoords.x, y: currentCoords.y - 1}
			}
		}
	}
	return nil
}

func lookRight(currentCoords Coord, prevCoords *Coord, grid [][]string) *Coord {
	if currentCoords.x+1 < len(grid[currentCoords.y]) {
		if prevCoords == nil || prevCoords.y != currentCoords.y || prevCoords.x != currentCoords.x+1 {
			if slices.Contains([]string{"-", "J", "7", "S"}, grid[currentCoords.y][currentCoords.x+1]) {
				return &Coord{x: currentCoords.x + 1, y: currentCoords.y}
			}
		}
	}
	return nil
}

func lookDown(currentCoords Coord, prevCoords *Coord, grid [][]string) *Coord {
	if currentCoords.y+1 < len(grid) {
		if prevCoords == nil || prevCoords.y != currentCoords.y+1 || prevCoords.x != currentCoords.x {
			if slices.Contains([]string{"|", "L", "J", "S"}, grid[currentCoords.y+1][currentCoords.x]) {
				return &Coord{x: currentCoords.x, y: currentCoords.y + 1}
			}
		}
	}
	return nil
}

func lookLeft(currentCoords Coord, prevCoords *Coord, grid [][]string) *Coord {
	if currentCoords.x-1 >= 0 {
		if prevCoords == nil || prevCoords.y != currentCoords.y || prevCoords.x != currentCoords.x-1 {
			if slices.Contains([]string{"-", "L", "F", "S"}, grid[currentCoords.y][currentCoords.x-1]) {
				return &Coord{x: currentCoords.x - 1, y: currentCoords.y}
			}
		}
	}
	return nil
}

func part1(grid [][]string) {
	sCoords := findS(grid)
	loop := findLoop(sCoords, grid)
	fmt.Println("part 1 result:", len(loop)/2)
}
