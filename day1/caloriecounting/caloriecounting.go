package caloriecounting

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTotalCaloriesOfTheElfCarryingMostCalories(elvesItemsCalories [][]int) int {
	maxCalories := 0
	for _, elfInventory := range elvesItemsCalories {
		sum := getSumOfCaloriesOfElfInventory(elfInventory)
		if sum > maxCalories {
			maxCalories = sum
		}
	}
	return maxCalories
}

func getSumOfCaloriesOfElfInventory(elfInventory []int) int {
	sumCalories := 0
	for _, calorie := range elfInventory {
		sumCalories += calorie
	}
	return sumCalories
}

func GetInventoryFromInputString(inputData string) [][]int {
	inputData = strings.TrimSpace(inputData)
	if inputData == "" {
		return [][]int{}
	}

	return getInventoryFromString(inputData)
}

func getInventoryFromString(inputData string) [][]int {
	splittedString := strings.Split(inputData, "\n\n")
	inventory := make([][]int, 0, len(splittedString))

	for _, caloriesOfSingleElfStr := range splittedString {
		caloriesOfSingleElfSlice := strings.Split(caloriesOfSingleElfStr, "\n")

		caloriesOfSingleElf := getCaloriesFromSlice(caloriesOfSingleElfSlice)
		inventory = append(inventory, caloriesOfSingleElf)
	}

	return inventory
}

func getCaloriesFromSlice(inputSlice []string) []int {
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
