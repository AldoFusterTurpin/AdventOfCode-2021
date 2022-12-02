package day1_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day1"
)

func TestGetTotalCaloriesOfTheElfCarryingMostCalories(t *testing.T) {
	expected := 24000
	elvesItemsCalories := day1.ElvesItemsCalories{
		day1.ElfInventory{
			1000,
			2000,
			3000,
		},
		day1.ElfInventory{4000},
		day1.ElfInventory{5000, 6000},
		day1.ElfInventory{7000, 8000, 9000},
		day1.ElfInventory{10000},
	}

	got := day1.GetTotalCaloriesOfTheElfCarryingMostCalories(elvesItemsCalories)
	if expected != got {
		t.Fatalf("expected %v, but got %v", expected, got)
	}
}
