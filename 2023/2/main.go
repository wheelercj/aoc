package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handful struct {
	red, green, blue int64
}

func main() {
	games := parseInput("input.txt")
	part1(games)
	part2(games)
}

func parseInput(fileName string) [][]Handful {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	text := string(bytes)
	lines := strings.Split(text, "\r\n")

	games := make([][]Handful, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		games[i] = make([]Handful, 0)

		gameIdAndHandfuls := strings.Split(line, ": ")
		_, handfulsStr := gameIdAndHandfuls[0], gameIdAndHandfuls[1]
		handfuls := strings.Split(handfulsStr, "; ")
		for _, handful := range handfuls {
			games[i] = append(games[i], Handful{})
			last := len(games[i]) - 1

			colorNums := strings.Split(handful, ", ")
			for _, colorNum := range colorNums {
				numAndColor := strings.Split(colorNum, " ")
				numStr, color := numAndColor[0], numAndColor[1]
				num, err := strconv.ParseInt(numStr, 10, 0)
				if err != nil {
					panic(err)
				}
				if color == "red" {
					games[i][last].red += num
				} else if color == "green" {
					games[i][last].green += num
				} else if color == "blue" {
					games[i][last].blue += num
				}
			}
		}
	}

	// fmt.Println("games:", games)
	return games
}

func part1(games [][]Handful) {
	var possibleGameIdSum int
	for i, game := range games {
		possible := true
		for _, handful := range game {
			if handful.red > 12 || handful.green > 13 || handful.blue > 14 {
				possible = false
				break
			}
		}
		if possible {
			possibleGameIdSum += i + 1
		}
	}

	fmt.Println("possible game ID sum:", possibleGameIdSum)
}

func part2(games [][]Handful) {
	var minCubesPowerSum int64
	for _, game := range games {
		var minCubes Handful
		for _, handful := range game {
			if handful.red > minCubes.red {
				minCubes.red = handful.red
			}
			if handful.green > minCubes.green {
				minCubes.green = handful.green
			}
			if handful.blue > minCubes.blue {
				minCubes.blue = handful.blue
			}
		}

		minCubesPower := minCubes.red * minCubes.green * minCubes.blue
		minCubesPowerSum += minCubesPower
	}

	fmt.Println("min cubes power sum:", minCubesPowerSum)
}
