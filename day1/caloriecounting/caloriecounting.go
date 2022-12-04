package caloriecounting

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInventoryFromInputString(inputData string) [][]int {
	inputData = strings.TrimSpace(inputData)
	if inputData == "" {
		return [][]int{}
	}

	splittedString := strings.Split(inputData, "\n\n")
	inventory := make([][]int, 0, len(splittedString))

	for _, v := range splittedString {
		d := strings.Split(v, "\n")
		calories := make([]int, 0, len(d))

		for _, v := range d {
			textInt, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			calories = append(calories, textInt)
		}
		inventory = append(inventory, calories)
	}
	return inventory
}

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
