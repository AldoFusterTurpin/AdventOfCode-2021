package caloriecounting_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day1/caloriecounting"
)

func TestGetTotalCaloriesOfTheElfCarryingMostCalories(t *testing.T) {
	expected := 24000
	elvesItemsCalories := caloriecounting.ElvesItemsCalories{
		caloriecounting.ElfInventory{
			1000,
			2000,
			3000,
		},
		caloriecounting.ElfInventory{4000},
		caloriecounting.ElfInventory{5000, 6000},
		caloriecounting.ElfInventory{7000, 8000, 9000},
		caloriecounting.ElfInventory{10000},
	}

	got := caloriecounting.GetTotalCaloriesOfTheElfCarryingMostCalories(elvesItemsCalories)
	if expected != got {
		t.Fatalf("expected %v, but got %v", expected, got)
	}
}
