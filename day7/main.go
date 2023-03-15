package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day7/device"
)

func main() {
	filePath := "day7/data/input.txt"
	inputString := getContentOfFile(filePath)
	maxSizePerDirectory := 100000
	result := device.SolveProblem(inputString, maxSizePerDirectory)
	fmt.Printf("Result: %v\n", result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
