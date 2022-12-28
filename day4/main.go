package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
For the line
2-4,6-8
we have:
x1=2 y1=4
x2=6 y2=8
*/

func main() {
	filePath := "day4/data/input.txt"
	inputString := getContentOfFile(filePath)
	lines := strings.Split(inputString, "\n")

	n, err := solveProblem(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Answer: %v\n", n)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}

func solveProblem(lines []string) (int, error) {
	n := 0
	for _, line := range lines {
		lineSplitted := strings.Split(line, ",")
		firstRangeStr := lineSplitted[0]
		secondRangeStr := lineSplitted[1]

		x1, y1, err := extractRangeFromStr(firstRangeStr)
		if err != nil {
			return 0, err
		}

		x2, y2, err := extractRangeFromStr(secondRangeStr)
		if err != nil {
			return 0, err
		}

		// part 1
		// if doesOneRangeFullyContainTheOther(x1, y1, x2, y2) {
		// 	n++
		// }

		// part 2
		if doRangesOverlap(x1, y1, x2, y2) {
			n++
		}
	}
	return n, nil
}

func extractRangeFromStr(line string) (int, int, error) {
	rangeSplitted := strings.Split(line, "-")
	xStr := rangeSplitted[0]
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return 0, 0, err
	}

	yStr := rangeSplitted[1]
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

// Part 1
func doesOneRangeFullyContainTheOther(x1, y1, x2, y2 int) bool {
	if x1 >= x2 && y1 <= y2 ||
		x2 >= x1 && y2 <= y1 {
		return true
	}
	return false
}

// Part 2
func doRangesOverlap(x1, y1, x2, y2 int) bool {
	if x1 >= x2 && x1 <= y2 || x2 >= x1 && x2 <= y1 {
		return true
	}
	return false
}
