package rucksack

import (
	"errors"
	"strings"
	"unicode"
)

func GetSolution(rucksacks []string) (int, error) {
	commonItems := make([]rune, len(rucksacks))

	for i, rucksack := range rucksacks {
		commonItem, err := GetCommonItemInRucksack(rucksack)
		if err != nil {
			return 0, err
		}
		commonItems[i] = commonItem
	}

	return GetSumOfPrioritiesOfItems(commonItems)
}

func GetCommonItemInRucksack(rucksuck string) (rune, error) {
	size := len(rucksuck)
	firstHalf := rucksuck[:size/2]
	secondHalf := rucksuck[size/2:]

	for _, item := range firstHalf {
		if itemInCompartment(item, secondHalf) {
			return item, nil
		}
	}

	for _, item := range secondHalf {
		if itemInCompartment(item, firstHalf) {
			return item, nil
		}
	}
	return rune(0), errors.New("no common item found")
}

func itemInCompartment(item rune, compartment string) bool {
	return strings.ContainsRune(compartment, item)
}

func GetSumOfPrioritiesOfItems(items []rune) (int, error) {
	sum := 0
	for _, item := range items {
		// fmt.Printf("%T\n", item)
		// fmt.Printf("%v\n\n", string(item))
		// fmt.Printf("%v\n\n", item)
		sum += getPriority(item)
	}
	return sum, nil
}

func getPriority(item rune) int {
	if unicode.IsLower(item) {
		// Lowercase item types a through z have priorities 1 through 26.
		// 'a' maps to 97:
		// 97-1 == x -> x = 96 -> amount that should be substracted
		const lowerOffset = 96
		return int(item) - lowerOffset
	}

	// Uppercase item types A through Z have priorities 27 through 52.
	// 'A' maps to 65:
	// 65-27 == x -> x = 38 -> amount that should be substracted
	const upperOffset = 38
	return int(item) - upperOffset
}
