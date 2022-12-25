package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day2/strategyguide"
)

func main() {
	filePath := "day2/data/input.txt"
	inputString := getContentOfFile(filePath)

	totalScore, err := strategyguide.ApplyStrategyGuide(inputString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("total score after applying strategy guide %v\n", totalScore)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
