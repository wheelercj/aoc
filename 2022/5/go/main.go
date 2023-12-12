package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)
	stacksAndSteps := strings.Split(content, "\r\n\r\n")
	stacksStr, stepsStr := stacksAndSteps[0], stacksAndSteps[1]
	stacks := getStacks(stacksStr)
	steps := getSteps(stepsStr)
	moveCrates9000(stacks, steps)
	topCrates := getTopCrates(stacks)
	if topCrates != "VCTFTJQCG" {
		panic(topCrates)
	}
	fmt.Printf("%v\n", topCrates)
}

func getStacks(stacksStr string) [][]string {
	stacksLines := strings.Split(stacksStr, "\r\n")
	stackCount := strings.Count(stacksLines[len(stacksLines)-2], "]")
	stacks := make([][]string, stackCount)
	stacksTransposed := make([][]string, len(stacksLines)-1)
	k := 0
	for i := len(stacksLines) - 2; i >= 0; i-- {
		line := stacksLines[i]
		crates := make([]string, stackCount)
		m := 0
		for j := 1; j < len(line); j += 4 {
			crates[m] = string(line[j])
			m++
		}

		stacksTransposed[k] = crates
		k++
	}
	for i := 0; i < len(stacksTransposed[0]); i++ {
		for j := 0; j < len(stacksTransposed); j++ {
			if stacksTransposed[j][i] != " " {
				stacks[i] = append(stacks[i], stacksTransposed[j][i])
			}
		}
	}
	return stacks
}

func getSteps(stepsStr string) [][3]int {
	stepsLines := strings.Split(stepsStr, "\r\n")
	steps := make([][3]int, len(stepsLines))
	for i, line := range stepsLines {
		words := strings.Split(line, " ")
		n1, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(words[3])
		if err != nil {
			panic(err)
		}
		n3, err := strconv.Atoi(words[5])
		if err != nil {
			panic(err)
		}
		steps[i] = [3]int{n1, n2, n3}
	}
	return steps
}

func moveCrates9000(stacks [][]string, steps [][3]int) {
	for _, step := range steps {
		moveCount, sourceI, destI := step[0], step[1]-1, step[2]-1
		source := stacks[sourceI]
		dest := stacks[destI]
		for i := 0; i < moveCount; i++ {
			z := len(source) - 1
			dest = append(dest, source[z])
			source = source[:z]
		}
		stacks[sourceI] = source
		stacks[destI] = dest
	}
}

func getTopCrates(stacks [][]string) string {
	topCrates := make([]string, len(stacks))
	for i, stack := range stacks {
		topCrates[i] = stack[len(stack)-1]
	}
	return strings.Join(topCrates, "")
}

func part2() {
	contentBytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)
	stacksAndSteps := strings.Split(content, "\r\n\r\n")
	stacksStr, stepsStr := stacksAndSteps[0], stacksAndSteps[1]
	stacks := getStacks(stacksStr)
	steps := getSteps(stepsStr)
	moveCrates9001(stacks, steps)
	topCrates := getTopCrates(stacks)
	if topCrates != "GCFGLDNJZ" {
		panic(topCrates)
	}
	fmt.Printf("%v\n", topCrates)
}

func moveCrates9001(stacks [][]string, steps [][3]int) {
	for _, step := range steps {
		moveCount, sourceI, destI := step[0], step[1]-1, step[2]-1
		source := stacks[sourceI]
		dest := stacks[destI]
		
		z := len(source) - moveCount
		dest = append(dest, source[z:]...)
		source = source[:z]

		stacks[sourceI] = source
		stacks[destI] = dest
	}
}
