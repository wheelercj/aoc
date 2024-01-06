package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cache map[string]int = make(map[string]int)

type Row struct {
	conditions  string
	damageSizes []int
}

func main() {
	rows := parseInput("input.txt")
	part1(rows)
	part2(rows)
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
		rows[i].conditions = conditions
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

func stringify(nums []int) string {
	s := make([]string, len(nums))
	for i, num := range nums {
		s[i] = fmt.Sprint(num)
	}
	return strings.Join(s, ",")
}

func getArrangementCount(conditions string, damageSizes []int) int {
	if len(conditions) == 0 {
		if len(damageSizes) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(damageSizes) == 0 {
		if strings.Contains(conditions, "#") {
			return 0
		} else {
			return 1
		}
	}

	hash := conditions + stringify(damageSizes)
	if v, ok := cache[hash]; ok {
		return v
	}

	var result int
	if strings.Contains(".?", string(conditions[0])) {
		result += getArrangementCount(conditions[1:], damageSizes)
	}
	if strings.Contains("#?", string(conditions[0])) {
		if damageSizes[0] <= len(conditions) && !strings.Contains(conditions[:damageSizes[0]], ".") {
			if damageSizes[0] == len(conditions) {
				result += getArrangementCount("", damageSizes[1:])
			} else if string(conditions[damageSizes[0]]) != "#" {
				result += getArrangementCount(conditions[damageSizes[0]+1:], damageSizes[1:])
			}
		}
	}

	cache[hash] = result
	return result
}

func quintuple(rows []Row) {
	for i := range rows {
		c := rows[i].conditions
		d := slices.Clone(rows[i].damageSizes)
		rows[i].conditions = c + "?" + c + "?" + c + "?" + c + "?" + c
		for j := 0; j < 4; j++ {
			rows[i].damageSizes = append(rows[i].damageSizes, d...)
		}
	}
}

func part1(rows []Row) {
	var sum int
	for _, row := range rows {
		sum += getArrangementCount(row.conditions, row.damageSizes)
	}
	fmt.Println("part 1 result:", sum)
}

func part2(rows []Row) {
	quintuple(rows)
	var sum int
	for _, row := range rows {
		sum += getArrangementCount(row.conditions, row.damageSizes)
	}
	fmt.Println("part 2 result:", sum)
}
