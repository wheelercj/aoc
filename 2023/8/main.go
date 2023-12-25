package main

import (
	"fmt"
	"os"
	"strings"
)

type Turn struct {
	left, right string
}

func main() {
	fileName := "input.txt"
	fmt.Println("file name:", fileName)
	turns, network := parseInput(fileName)
	part1(turns, network)
	part2(turns, network)
}

func parseInput(fileName string) (string, map[string]Turn) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")

	turns := lines[0]
	network := make(map[string]Turn, 0)

	for i := 2; i < len(lines); i++ {
		keyAndValues := strings.Split(lines[i], " = ")
		key := keyAndValues[0]
		values := strings.Split(strings.Trim(keyAndValues[1], "()"), ", ")
		network[key] = Turn{left: values[0], right: values[1]}
	}

	return turns, network
}

func allEndNodes(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

func part1(turns string, network map[string]Turn) {
	fmt.Println("part 1")
	startNode := "AAA"
	if _, ok := network[startNode]; !ok {
		fmt.Println("\tnode", startNode, "not in the network")
		return
	}

	currentNode := startNode
	stepCount := 0

	for currentNode != "ZZZ" {
		for _, turn := range turns {
			stepCount++
			if string(turn) == "L" {
				currentNode = network[currentNode].left
			} else {
				currentNode = network[currentNode].right
			}
		}
	}

	fmt.Println("\tresult:", stepCount)
}

func part2(turns string, network map[string]Turn) {
	fmt.Println("part 2")

	var currentNodes []string
	for key := range network {
		if strings.HasSuffix(key, "A") {
			currentNodes = append(currentNodes, key)
		}
	}

	stepCount := 0
	turnIndex := 0
	for !allEndNodes(currentNodes) {
		if string(turns[turnIndex]) == "L" {
			for i, node := range currentNodes {
				currentNodes[i] = network[node].left
			}
		} else {
			for i, node := range currentNodes {
				currentNodes[i] = network[node].right
			}
		}

		stepCount++
		turnIndex = (turnIndex + 1) % len(turns)
		fmt.Print("\r                                     \r\tstep count: ", stepCount)
	}
	fmt.Print("\r                                                 \r")

	fmt.Println("\tresult:", stepCount) // > 22199
}
