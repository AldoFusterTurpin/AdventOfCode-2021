package rucksack_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day3/rucksack"
)

func TestGetSolution(t *testing.T) {
	type testData struct {
		rucksacks []string
		expected  int
	}

	tests := map[string]testData{
		"example_1": {
			rucksacks: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			expected: 157,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := rucksack.GetSolution(tc.rucksacks)
			if err != nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
func TestGetCommonItemInRucksack(t *testing.T) {
	type testData struct {
		compartmensItems string
		expected         rune
	}

	tests := map[string]testData{
		"example_1": {
			compartmensItems: "vJrwpWtwJgWrhcsFMMfFFhFp",
			expected:         'p',
		},
		"example_2": {
			compartmensItems: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			expected:         'L',
		},
		"example_3": {
			compartmensItems: "PmmdzqPrVvPwwTWBwg",
			expected:         'P',
		},
		"example_4": {
			compartmensItems: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			expected:         'v',
		},
		"example_5": {
			compartmensItems: "ttgJtRGJQctTZtZT",
			expected:         't',
		},
		"example_6": {
			compartmensItems: "CrZsJsPPZsGzwwsLwLmpwMDw",
			expected:         's',
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := rucksack.GetCommonItemInRucksack(tc.compartmensItems)
			if err != nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestGetSumOfPrioritiesOfItems(t *testing.T) {
	type testData struct {
		items    []rune
		expected int
	}

	tests := map[string]testData{
		"example_1": {
			items:    []rune{'p', 'L', 'P', 'v', 't', 's'},
			expected: 157,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := rucksack.GetSumOfPrioritiesOfItems(tc.items)
			if err != nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
