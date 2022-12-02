package day1

type ElfInventory []int
type ElvesItemsCalories []ElfInventory

func GetTotalCaloriesOfTheElfCarryingMostCalories(elvesItemsCalories ElvesItemsCalories) int {
	maxCalories := 0
	for _, elfInventory := range elvesItemsCalories {
		sum := getSumOfCaloriesOfElfInventory(elfInventory)
		if sum > maxCalories {
			maxCalories = sum
		}
	}
	return maxCalories
}

func getSumOfCaloriesOfElfInventory(elfInventory ElfInventory) int {
	sumCalories := 0
	for _, calorie := range elfInventory {
		sumCalories += calorie
	}
	return sumCalories
}
