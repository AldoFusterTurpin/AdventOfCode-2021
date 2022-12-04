package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day1/caloriecounting"
)

func main() {
	filePath := "day1/data/input.txt"
	inputString := getContentOfFile(filePath)

	inventory := caloriecounting.GetAllElvesCalories(inputString)
	result := caloriecounting.GetTotalCaloriesOfTheElfCarryingMostCalories(inventory)
	fmt.Println(result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
