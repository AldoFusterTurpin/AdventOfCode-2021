package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day3/rucksack"
)

func main() {
	filePath := "day3/data/input.txt"
	inputString := getContentOfFile(filePath)
	splitted := strings.Split(inputString, "\n")

	solution, err := rucksack.GetSolutionPart2(splitted)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(splitted)
	fmt.Printf("sum of the priorities is %v\n", solution)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
