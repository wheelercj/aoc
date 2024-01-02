package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	lines := parseInput("input.txt")
	extraEmptyRows := findExtraEmptyRows(lines)
	extraEmptyCols := findExtraEmptyCols(lines)
	universe := expandUniverse(lines, extraEmptyRows, extraEmptyCols)
	galaxies := findGalaxies(universe)
	part1(universe, galaxies)
}

func parseInput(fileName string) []string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")
	return lines
}

// findExtraEmptyRows finds the indexes of the rows in lines that do not contain any #
// characters, and returns the indexes to where the added empty rows will be when the
// universe is expanded.
func findExtraEmptyRows(lines []string) []int {
	var extraEmptyRows []int
	var foundCount int
	for i, line := range lines {
		if !strings.Contains(line, "#") {
			extraEmptyRows = append(extraEmptyRows, i+foundCount+1)
			foundCount++
		}
	}
	return extraEmptyRows
}

// findExtraEmptyCols finds the indexes of the columns in lines that do not contain any
// # characters, and returns the indexes to where the added empty columns will be when
// the universe is expanded.
func findExtraEmptyCols(lines []string) []int {
	var extraEmptyCols []int
	var foundCount int
	for x := 0; x < len(lines[0]); x++ {
		var foundHashtag bool
		for _, line := range lines {
			if string(line[x]) == "#" {
				foundHashtag = true
				break
			}
		}
		if !foundHashtag {
			extraEmptyCols = append(extraEmptyCols, x+foundCount+1)
			foundCount++
		}
	}
	return extraEmptyCols
}

// expandUniverse takes everything in lines and adds extra "empty" (period-filled) rows
// and columns at the indexes specified by extraEmptyRows and extraEmptyCols. These
// indexes must be sorted (ascending) and be the indexes of where the rows and columns
// will be after the expansion, not before.
func expandUniverse(lines []string, extraEmptyRows, extraEmptyCols []int) [][]string {
	universeHeight := len(lines) + len(extraEmptyRows)
	universeWidth := len(lines[0]) + len(extraEmptyCols)
	universe := make([][]string, universeHeight)

	linesY := 0
	linesX := 0
	for y := range universe {
		universe[y] = make([]string, universeWidth)
		if slices.Contains(extraEmptyRows, y) {
			for x := range universe[y] {
				universe[y][x] = "."
			}
			continue
		}

		for x := range universe[y] {
			if slices.Contains(extraEmptyCols, x) {
				universe[y][x] = "."
				continue
			}
			if string(lines[linesY][linesX]) == "#" {
				universe[y][x] = "#"
			} else {
				universe[y][x] = "."
			}
			linesX++
		}
		linesY++
		linesX = 0
	}

	return universe
}

func findGalaxies(universe [][]string) []Coord {
	var galaxies []Coord
	for y, row := range universe {
		for x, ch := range row {
			if ch == "#" {
				galaxies = append(galaxies, Coord{x: x, y: y})
			}
		}
	}
	return galaxies
}

// findDistance finds the character distance between two galaxies in universe.
func findDistance(universe [][]string, galaxy1, galaxy2 Coord) int {
	a := int(math.Abs(float64(galaxy2.x - galaxy1.x)))
	b := int(math.Abs(float64(galaxy2.y - galaxy1.y)))
	return a + b
}

func part1(universe [][]string, galaxies []Coord) {
	var sum int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += findDistance(universe, galaxies[i], galaxies[j])
		}
	}
	fmt.Println("part 1 result:", sum)
}

func part2() {
	fmt.Println("part 2 result:")
}
