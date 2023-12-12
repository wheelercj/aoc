// https://adventofcode.com/2022/day/7

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileSystem := part1("../input.txt")
	part2(fileSystem)
}

func PrintMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%q -> %d\n", k, v)
	}
}

func part1(inputFilePath string) map[string]int {
	contentBytes, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contentBytes), "\r\n")

	fileSystem := make(map[string]int) // fileOrFolder path -> fileOrFolder size
	// Every path starts with a slash, and every folder path ends with a slash.

	fileSystem["/"] = 0
	var currentPath string
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd ") {
			newFolder := strings.TrimPrefix(line, "$ cd ")
			switch newFolder {
			case "/":
				currentPath = "/"
			case "..":
				currentPath = getParentFolder(currentPath)
			default:
				currentPath = openFolder(currentPath, newFolder)
			}
		} else if strings.HasPrefix(line, "dir ") {
			folderName := strings.TrimPrefix(line, "dir ")
			fileSystem[currentPath+folderName+"/"] = 0
		} else if line != "$ ls" {
			// The line starts with a number and ends with a file name.
			fileSizeAndName := strings.Split(line, " ")
			fileSizeStr, fileName := fileSizeAndName[0], fileSizeAndName[1]
			fileSize, err := strconv.Atoi(fileSizeStr)
			if err != nil {
				panic(err)
			}
			fileSystem[currentPath+fileName] = fileSize
		}
	}

	// get folder sizes
	for k1 := range fileSystem {
		if strings.HasSuffix(k1, "/") {
			// k1 is a folder path
			for k2, v2 := range fileSystem {
				if !strings.HasSuffix(k2, "/") && strings.HasPrefix(k2, k1) {
					fileSystem[k1] += v2
				}
			}
		}
	}

	sum := 0 // sum of sizes of small folders
	for k, v := range fileSystem {
		if strings.HasSuffix(k, "/") && v <= 100000 {
			sum += v
		}
	}

	fmt.Println(sum)
	return fileSystem
}

func openFolder(currentPath, newFolder string) string {
	return currentPath + newFolder + "/"
}

func getParentFolder(currentPath string) string {
	s := strings.Split(currentPath, "/")
	return strings.Join(s[:len(s)-2], "/") + "/"
}

func part2(fileSystem map[string]int) {
	totalDiskSpace := 70000000
	currentDiskSpace := totalDiskSpace - fileSystem["/"]
	totalSpaceNeeded := 30000000
	currentSpaceNeeded := totalSpaceNeeded - currentDiskSpace

	minName := ""
	minSize := 9999999999
	for k, v := range fileSystem {
		if v >= currentSpaceNeeded && v < minSize && strings.HasSuffix(k, "/") {
			minSize = v
			minName = k
		}
	}

	fmt.Printf("%v %v", minName, minSize)
}
