package device

import "log"

// Part 2
func SolveProblemPart2(input string, maxSizePerDir int) int {
	dirsInfo := GetDirsInfo(input)
	return GetSizeOfSmallestDirectoryThatShouldBeRemovedToRunTheUpdate(dirsInfo)
}

// Part 2
func GetSizeOfSmallestDirectoryThatShouldBeRemovedToRunTheUpdate(dirsInfo map[string]int) int {
	totalDiskSpaceAvailable := 70000000
	neededUnusedSpaceToPerformUpdate := 30000000

	rootFolder := "/"
	sizeOfRoot, present := dirsInfo[rootFolder]

	// this should never happen
	if !present {
		log.Fatal("root folder is not present in 'dirsInfo', aborting")
	}

	currentUnusedSpace := totalDiskSpaceAvailable - sizeOfRoot
	if currentUnusedSpace >= neededUnusedSpaceToPerformUpdate {
		// Enough space to perform update. No need to delete any folder.
		return 0
	}

	minDirSizeToBeDeleted := neededUnusedSpaceToPerformUpdate - currentUnusedSpace
	return getSmallestDirSizeWithAtLeastSize(minDirSizeToBeDeleted, dirsInfo)
}
