package rucksack_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day3/rucksack"
)

func TestGetCommonItemInCompartments(t *testing.T) {
	type testData struct {
		compartmensItems []string
		expected         string
	}

	tests := map[string]testData{
		"example_1": {
			compartmensItems: []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			expected:         "p",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := rucksack.GetCommonItemInCompartments(tc.compartmensItems)
			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
