package caloriecounting

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// used in part 2
func GetTotalCaloriesOfTheTopNElvesCarryingTheMostCalories(elvesCalories [][]int, n int) int {
	sumOfCaloriesOfEachElf := getSumOfCaloriesOfEachElf(elvesCalories)

	sort.Slice(sumOfCaloriesOfEachElf, func(i, j int) bool {
		return sumOfCaloriesOfEachElf[i] > sumOfCaloriesOfEachElf[j]
	})

	return getSumOfCalories(sumOfCaloriesOfEachElf[0:n])
}

func getSumOfCaloriesOfEachElf(elvesCalories [][]int) []int {
	sumOfCaloriesOfEachElf := make([]int, 0, len(elvesCalories))

	for _, elfCalories := range elvesCalories {
		sum := getSumOfCalories(elfCalories)
		sumOfCaloriesOfEachElf = append(sumOfCaloriesOfEachElf, sum)
	}
	return sumOfCaloriesOfEachElf
}

// used in part 1
func GetTotalCaloriesOfTheElfCarryingMostCalories(elvesCalories [][]int) int {
	maxSumCalories := 0
	for _, elfCalories := range elvesCalories {
		sum := getSumOfCalories(elfCalories)
		if sum > maxSumCalories {
			maxSumCalories = sum
		}
	}

	return maxSumCalories
}

func getSumOfCalories(calories []int) int {
	sumCalories := 0
	for _, calories := range calories {
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
