package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day5/supplystacks"
)

func main() {
	filePath := "day5/data/input.txt"
	inputString := getContentOfFile(filePath)

	// For part 1
	// result := supplystacks.SolveProblem(inputString)

	// For part 2
	result := supplystacks.SolveProblemPart2(inputString)
	fmt.Printf("\nSolution: %v\n", result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
