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
	universe := parseInput("input.txt")
	emptyRows := findEmptyRows(universe)
	emptyCols := findEmptyCols(universe)
	galaxies := findGalaxies(universe)
	part1(universe, galaxies, emptyRows, emptyCols)
	part2(universe, galaxies, emptyRows, emptyCols)
}

func parseInput(fileName string) [][]string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")
	universe := make([][]string, len(lines))
	for i := range universe {
		universe[i] = strings.Split(lines[i], "")
	}
	return universe
}

// findEmptyRows finds the indexes of the rows in universe that do not contain any
// # characters.
func findEmptyRows(universe [][]string) []int {
	var emptyRows []int
	for y, row := range universe {
		if !slices.Contains(row, "#") {
			emptyRows = append(emptyRows, y)
		}
	}
	return emptyRows
}

// findEmptyCols finds the indexes of the columns in universe that do not contain
// any # characters.
func findEmptyCols(universe [][]string) []int {
	var emptyCols []int
	for x := 0; x < len(universe[0]); x++ {
		var foundHashtag bool
		for _, row := range universe {
			if row[x] == "#" {
				foundHashtag = true
				break
			}
		}
		if !foundHashtag {
			emptyCols = append(emptyCols, x)
		}
	}
	return emptyCols
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

// findDistance finds the character distance between two galaxies in universe. Each
// empty row and empty column between the galaxies add emptyWidth distance to the total.
func findDistance(
	universe [][]string, galaxy1, galaxy2 Coord, emptyRows, emptyCols []int, emptyWidth int,
) int {
	dist := int(math.Abs(float64(galaxy2.x - galaxy1.x)))
	dist += int(math.Abs(float64(galaxy2.y - galaxy1.y)))

	for _, emptyRow := range emptyRows {
		isBetween := galaxy1.y < emptyRow && emptyRow < galaxy2.y ||
			galaxy2.y < emptyRow && emptyRow < galaxy1.y
		if isBetween {
			dist += emptyWidth - 1
			// 1 is subtracted here because the original empty space was already added
			// above.
		}
	}

	for _, emptyCol := range emptyCols {
		isBetween := galaxy1.x < emptyCol && emptyCol < galaxy2.x ||
			galaxy2.x < emptyCol && emptyCol < galaxy1.x
		if isBetween {
			dist += emptyWidth - 1
		}
	}

	return dist
}

func findSumOfDistances(
	universe [][]string, galaxies []Coord, emptyRows, emptyCols []int, emptyWidth int,
) int {
	var sum int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += findDistance(
				universe, galaxies[i], galaxies[j], emptyRows, emptyCols, emptyWidth,
			)
		}
	}
	return sum
}

func part1(universe [][]string, galaxies []Coord, emptyRows, emptyCols []int) {
	emptyWidth := 2
	sum := findSumOfDistances(universe, galaxies, emptyRows, emptyCols, emptyWidth)
	fmt.Println("part 1 result:", sum)
}

func part2(universe [][]string, galaxies []Coord, emptyRows, emptyCols []int) {
	emptyWidth := 1000000
	sum := findSumOfDistances(universe, galaxies, emptyRows, emptyCols, emptyWidth)
	fmt.Println("part 2 result:", sum)
}
