package rucksack

import "errors"

func GetSolutionPart2(rucksacks []string) (int, error) {
	var commonItems []rune
	var groupsOfThree []string

	for i, rucksack := range rucksacks {
		groupsOfThree = append(groupsOfThree, rucksack)

		if (i+1)%3 == 0 {
			commonItem, err := GetCommonItemInGroupsOfThree(groupsOfThree)
			if err != nil {
				return 0, err
			}
			commonItems = append(commonItems, commonItem)
			groupsOfThree = make([]string, 0)
		}
	}

	return GetSumOfPrioritiesOfItems(commonItems)
}

// This could be implemented with a more generic approach iterating over
// groupsOfThree and use the index 'i'(current group) to look for the items of this current group
// in the other ones. Not needed at this point but the idea is simple.
func GetCommonItemInGroupsOfThree(groupsOfThree []string) (rune, error) {
	firstGroup := groupsOfThree[0]
	secondGroup := groupsOfThree[1]
	thirdGroup := groupsOfThree[2]

	for _, item := range firstGroup {
		if itemInCompartment(item, secondGroup) && itemInCompartment(item, thirdGroup) {
			return item, nil
		}
	}

	for _, item := range secondGroup {
		if itemInCompartment(item, firstGroup) && itemInCompartment(item, thirdGroup) {
			return item, nil
		}
	}

	for _, item := range thirdGroup {
		if itemInCompartment(item, firstGroup) && itemInCompartment(item, secondGroup) {
			return item, nil
		}
	}

	return rune(0), errors.New("no common item found")
}
