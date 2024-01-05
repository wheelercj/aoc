package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Imagine one rod on an abacus. Its beads cannot intersect nor change order. Start by
// moving all the beads to the left edge of the rod. Then gradually move the last bead
// to the right edge of the rod while counting the number of places it can be. Next,
// move the second-to-last bead one place to the right, move the last bead left so it's
// next to the second-to-last one, and repeat the process. Repeat these steps with all
// the beads until all of them are at the right edge. Throughout all of this, beads must
// keep at least a little space between each other, and skip places they obviously
// cannot be.

type Row struct {
	conditions  []string
	damageSizes []int
}

func main() {
	rows := parseInput("input.txt")
	part1(rows)
}

func parseInput(fileName string) []Row {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(bytes), "\r\n"), "\r\n")

	rows := make([]Row, len(lines))
	for i, line := range lines {
		c := strings.Split(line, " ")
		conditions, damageSizesStr := c[0], c[1]
		rows[i].conditions = strings.Split(conditions, "")
		damageSizesStrs := strings.Split(damageSizesStr, ",")
		rows[i].damageSizes = make([]int, len(damageSizesStrs))
		for j, numStr := range damageSizesStrs {
			rows[i].damageSizes[j] = mustParseInt(numStr)
		}
	}

	return rows
}

func mustParseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

type Group struct {
	start, end, size, arrangementCount int
}

// initGroups sets each damaged spring group's start and end indexes to that group's
// first possible position without overlapping any other group.
func initGroups(groups []Group, row Row) []Group {
	// First, set each group's size.
	for i := 0; i < len(groups); i++ {
		groups[i].size = row.damageSizes[i]
	}

	// Next, move the spring groups just enough so they don't overlap and are where
	// damaged springs could be.
	var j int
	for i, group := range groups {
		// Where should this spring group go?
		for ; j < len(row.conditions)-group.size+1; j++ {
			// Can it go here (at j)?
			canBeHere := true
			damagedOnLeft := j-1 >= 0 && row.conditions[j-1] == "#"
			damagedOnRight := j+group.size < len(row.conditions) && row.conditions[j+group.size] == "#"
			if damagedOnLeft || damagedOnRight {
				canBeHere = false
			} else {
				for k := j; k < j+group.size; k++ {
					spring := row.conditions[k]
					if spring != "#" && spring != "?" {
						canBeHere = false
						break
					}
				}
			}
			if canBeHere {
				groups[i].start = j
				groups[i].end = j + group.size - 1
				j += group.size + 1
				break
			}
		}
	}

	return groups
}

func getArrangementCount(row Row) int {
	groups := make([]Group, len(row.damageSizes)) // groups of damaged springs
	groups = initGroups(groups, row)
	n, _ := _getArrangementCount(0, row, groups)
	return n
}

func _getArrangementCount(groupIndex int, row Row, groups []Group) (int, []Group) {
	var arrangementCount int
	var n int

	// Move the groupIndex group as far left as it can go.
	if groupIndex > 0 {
		groups[groupIndex].start = groups[groupIndex-1].end + 2
		groups[groupIndex].end = groups[groupIndex].start + groups[groupIndex].size - 1
		if groups[groupIndex].end > -2+groups[groupIndex+1].start {
			// The groupIndex group has run out of space.
			return arrangementCount, groups
		}
		for !groupCanBeHere(groupIndex, row, groups) {
			groups[groupIndex].start++
			groups[groupIndex].end++
			if groups[groupIndex].end > -2+groups[groupIndex+1].start {
				// The groupIndex group has run out of space.
				return arrangementCount, groups
			}
		}
	}

	for {
		// If the next group is not the last group.
		if groupIndex+1 < len(groups)-1 {
			n, groups = _getArrangementCount(groupIndex+1, row, groups)
			arrangementCount += n
		} else {
			// Move the last group as far right as it can go.
			n, groups = baseGetArrangementCount(row, groups)
			arrangementCount += n
		}

		// If the groupIndex group is already as far right as it can go, return.
		if groups[groupIndex].end == -2+groups[groupIndex+1].start {
			return arrangementCount, groups
		}

		// Move the groupIndex group one space to the right.
		groups[groupIndex].start++
		groups[groupIndex].end++
		if groups[groupIndex].end > -2+groups[groupIndex+1].start {
			// The groupIndex group has run out of space.
			return arrangementCount, groups
		}
		for !groupCanBeHere(groupIndex, row, groups) {
			groups[groupIndex].start++
			groups[groupIndex].end++
			if groups[groupIndex].end > -2+groups[groupIndex+1].start {
				// The groupIndex group has run out of space.
				return arrangementCount, groups
			}
		}
	}
}

func baseGetArrangementCount(row Row, groups []Group) (int, []Group) {
	// Move the last group as far left as it can go.
	z := len(groups) - 1
	groups[z].start = groups[z-1].end + 2
	groups[z].end = groups[z].start + groups[z].size - 1

	// Move the last group right gradually while counting valid arragements.
	var validArragements int
	if groupCanBeHere(z, row, groups) && damagedSpringsAreCovered(row.conditions, groups) {
		validArragements++
	}
	for i := groups[z].end; i < len(row.conditions)-1; i++ {
		groups[z].start++
		groups[z].end++
		if groupCanBeHere(z, row, groups) && damagedSpringsAreCovered(row.conditions, groups) {
			validArragements++
		}
	}

	// When the last group is as far right as it can go, return the count.
	return validArragements, groups
}

// groupCanBeHere determines whether the damaged spring group's current location
// contains any undamaged springs.
func groupCanBeHere(groupIndex int, row Row, groups []Group) bool {
	for i := groups[groupIndex].start; i <= groups[groupIndex].end; i++ {
		if row.conditions[i] == "." {
			return false
		}
	}
	return true
}

func damagedSpringsAreCovered(conditions []string, groups []Group) bool {
	for i, spring := range conditions {
		if spring == "#" {
			var found bool
			for _, group := range groups {
				if group.start <= i && i <= group.end {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func part1(rows []Row) {
	var sum int
	for _, row := range rows {
		sum += getArrangementCount(row)
	}
	fmt.Println("part 1 result:", sum)
}

func part2() {
	fmt.Println("part 2 result:")
}
