package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day5/supplystacks"
)

func main() {
	filePath := "day5/data/input.txt"
	inputString := getContentOfFile(filePath)

	result := supplystacks.SolveProblem(inputString)
	fmt.Printf("\nSolution: %v\n", result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
