package rucksack

import (
	"errors"
	"strings"
)

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
