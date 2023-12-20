package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mm int // millimeters
type ms int // milliseconds

type Race struct {
	timeLimit   ms
	recordDist  mm
	winWayCount int
}

func main() {
	times, distances := parseInput("input.txt")
	part1(times, distances)
	part2(times, distances)
}

func parseInput(fileName string) (string, string) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\r\n")
	timesStr, distancesStr := lines[0], lines[1]
	return timesStr, distancesStr
}

// getInts gets a list of whitespace-separated integers from a string. The character(s)
// before the first whitespace are ignored.
func getInts(numsStr string) []int {
	var nums []int
	numsStrs := strings.Fields(numsStr)
	for i := 1; i < len(numsStrs); i++ {
		nums = append(nums, mustParseInt(numsStrs[i]))
	}
	return nums
}

func mustParseInt(numStr string) int {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func part1(timesStr, distancesStr string) {
	times, distances := getInts(timesStr), getInts(distancesStr)
	var races []Race
	for i := 0; i < len(times); i++ {
		races = append(
			races,
			Race{timeLimit: ms(times[i]), recordDist: mm(distances[i])},
		)
	}

	for i, race := range races {
		fmt.Printf("race %d\n", i+1)
		races[i].winWayCount = getWinWayCount(race)
	}

	product := 1
	for _, race := range races {
		product *= race.winWayCount
	}

	fmt.Println("\npart 1 result:", product)
}

func getWinWayCount(race Race) int {
	min, err := findMinChargeTimeNeeded(race)
	if err != nil {
		panic(err)
	}
	max, err := findMaxChargeTimeNeeded(race)
	if err != nil {
		panic(err)
	}
	return int(max-min) + 1
}

// findMinChargeTimeNeeded performs a binary search to find the minimum charge time
// required to beat the race's record.
func findMinChargeTimeNeeded(race Race) (ms, error) {
	startTime := ms(0)            // start of where to search, not necessarily start of race
	endTime := ms(race.timeLimit) // end of where to search, not necessarily end of race
	for startTime < endTime-1 {
		midTime := ms((endTime-startTime)/2 + startTime)
		dist1 := getDist(midTime, race.timeLimit)
		dist2 := getDist(midTime+1, race.timeLimit)
		if dist1 <= race.recordDist && dist2 > race.recordDist {
			return midTime + 1, nil
		} else if dist1 > race.recordDist {
			endTime = midTime + 1
		} else if dist2 < race.recordDist {
			startTime = midTime
		}
	}

	return 0, fmt.Errorf("No minimum charge time to beat the record found.")
}

// findMaxChargeTimeNeeded performs a binary search to find the maximum charge time (in
// milliseconds) required to beat the race's record.
func findMaxChargeTimeNeeded(race Race) (ms, error) {
	startTime := ms(0)            // start of where to search, not necessarily start of race
	endTime := ms(race.timeLimit) // end of where to search, not necessarily end of race
	for startTime < endTime-1 {
		midTime := ms((endTime-startTime)/2 + startTime)
		dist1 := getDist(midTime, race.timeLimit)
		dist2 := getDist(midTime+1, race.timeLimit)
		if dist1 > race.recordDist && dist2 <= race.recordDist {
			return midTime, nil
		} else if dist1 < race.recordDist {
			endTime = midTime + 1
		} else if dist2 > race.recordDist {
			startTime = midTime
		}
	}

	return 0, fmt.Errorf("No maximum charge time to beat the record found.")
}

// getDist determines the distance (in millimeters) the boat can travel after
// timeCharged milliseconds spent charging during a race that is a total of timeLimit
// milliseconds long.
func getDist(timeCharged, timeLimit ms) mm {
	return mm(timeCharged * (timeLimit - timeCharged))
}

func part2(timeStr, distanceStr string) {
	timeStr = strings.ReplaceAll(timeStr, " ", "")
	distanceStr = strings.ReplaceAll(distanceStr, " ", "")
	timeStr = strings.TrimPrefix(timeStr, "Time:")
	distanceStr = strings.TrimPrefix(distanceStr, "Distance:")

	timeLimit := ms(mustParseInt(timeStr))
	recordDist := mm(mustParseInt(distanceStr))
	race := Race{timeLimit: timeLimit, recordDist: recordDist}
	winWayCount := getWinWayCount(race)

	fmt.Println("part 2 result:", winWayCount)
}
