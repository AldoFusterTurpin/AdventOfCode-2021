package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day1/caloriecounting"
)

func main() {
	filePath := "day1/data/input.txt"
	inputString := getStringFromFilePath(filePath)

	inventory := caloriecounting.GetInventoryFromInputString(inputString)
	fmt.Println(inventory)
}

func getStringFromFilePath(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
