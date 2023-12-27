package main

import (
	"fmt"
	"os"
	"strings"
)

type Turn struct {
	left, right string
}

type EndNode struct {
	name      string
	stepCount int // the number of steps at which the end node was first found
	period    int // the number of steps between each encounter of the end node
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

	endNodes := make([]EndNode, len(currentNodes))

	stepCount := 0
	periodCount := 0
	turnIndex := 0
	for periodCount < 6 {
		if string(turns[turnIndex]) == "L" {
			for i, node := range currentNodes {
				currentNodes[i] = network[node].left
				if strings.HasSuffix(currentNodes[i], "Z") {
					if len(endNodes[i].name) == 0 {
						endNodes[i] = EndNode{name: currentNodes[i], stepCount: stepCount}
					} else if endNodes[i].period == 0 {
						endNodes[i].period = stepCount - endNodes[i].stepCount
						periodCount++
					}
				}
			}
		} else {
			for i, node := range currentNodes {
				currentNodes[i] = network[node].right
				if strings.HasSuffix(currentNodes[i], "Z") {
					if len(endNodes[i].name) == 0 {
						endNodes[i] = EndNode{name: currentNodes[i], stepCount: stepCount}
					} else if endNodes[i].period == 0 {
						endNodes[i].period = stepCount - endNodes[i].stepCount
						periodCount++
					}
				}
			}
		}

		stepCount++
		turnIndex = (turnIndex + 1) % len(turns)
		fmt.Print("\r                                     \r\tstep count: ", stepCount)
	}
	fmt.Print("\r                                                 \r")

	periods := make([]int, len(currentNodes))
	for i, endNode := range endNodes {
		periods[i] = endNode.period
	}

	fmt.Println("\tresult:", lcm(periods...))
}

// lcm finds the least common multiple of two or more integers.
func lcm(integers ...int) int {
	if len(integers) < 2 {
		panic("The lcm function requires 2 or more integers.")
	}
	result := integers[0] * integers[1] / gcd(integers[0], integers[1])
	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

// gcd finds the greatest common denominator of two integers.
func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
