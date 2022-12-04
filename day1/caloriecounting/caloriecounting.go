package caloriecounting

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTotalCaloriesOfTheElfCarryingMostCalories(elvesCalories [][]int) int {
	maxSumCalories := 0
	for _, elfCalories := range elvesCalories {
		sum := getSumOfCaloriesOfElf(elfCalories)
		if sum > maxSumCalories {
			maxSumCalories = sum
		}
	}

	return maxSumCalories
}

func getSumOfCaloriesOfElf(elfCalories []int) int {
	sumCalories := 0
	for _, calories := range elfCalories {
		sumCalories += calories
	}
	return sumCalories
}

func GetAllElvesCalories(inputData string) [][]int {
	inputData = strings.TrimSpace(inputData)
	if inputData == "" {
		return [][]int{}
	}

	return getAllElvesCaloriesFromString(inputData)
}

func getAllElvesCaloriesFromString(inputData string) [][]int {
	splittedString := strings.Split(inputData, "\n\n")
	inventory := make([][]int, 0, len(splittedString))

	for _, caloriesOfSingleElfStr := range splittedString {
		caloriesOfSingleElfSlice := strings.Split(caloriesOfSingleElfStr, "\n")

		elfCalories := getElfCaloriesFromSlice(caloriesOfSingleElfSlice)
		inventory = append(inventory, elfCalories)
	}

	return inventory
}

func getElfCaloriesFromSlice(inputSlice []string) []int {
	calories := make([]int, 0, len(inputSlice))

	for _, v := range inputSlice {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		calories = append(calories, intValue)
	}
	return calories
}
