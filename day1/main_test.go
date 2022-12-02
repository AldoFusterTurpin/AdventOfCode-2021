package main_test

import "testing"

type elfInventory []int

type elvesItemsCalories []elfInventory

func TestGetTotalCaloriesOfTheElfCarryingMostCalories(t *testing.T) {
	expected := 24000
	elvesItemsCalories := elvesItemsCalories{
		elfInventory{
			1000,
			2000,
			3000,
		},
		elfInventory{4000},
		elfInventory{5000, 6000},
		elfInventory{7000, 8000, 9000},
		elfInventory{10000},
	}

	got := GetTotalCaloriesOfTheElfCarryingMostCalories(elvesItemsCalories)
	if expected != got {
		t.Fatalf("expected %v, but got %v", expected, got)
	}
}
