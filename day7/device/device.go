package device

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

const defaultDirSeparator = "/"

/*
You can find a complete "manual" exercise I performed to think about how to solve the problem.
Understand the problem -> think -> TDD -> problem solved.
Sometimes it is not straightforward :)
*/

// Path contains the absolute path of the directory.
// Content contains the subelements (files and folders) of the directory.
type Dir struct {
	Path    string
	Content []string
}

func getSmallestDirSizeWithAtLeastSize(minDirSizeToBeDeleted int, dirsInfo map[string]int) int {
	sizeOfSmallestDir := 0
	firstDirFound := false

	for _, dirSize := range dirsInfo {
		if dirSize >= minDirSizeToBeDeleted {
			if !firstDirFound {
				sizeOfSmallestDir = dirSize
				firstDirFound = true
			} else if dirSize < sizeOfSmallestDir {
				sizeOfSmallestDir = dirSize
			}
		}
	}

	return sizeOfSmallestDir
}

// Part 1
func SolveProblem(input string, maxSizePerDir int) int {
	input = strings.TrimSpace(input)
	sum := 0
	sizeOfDirs := GetDirsInfo(input)
	for _, v := range sizeOfDirs {
		if v <= maxSizePerDir {
			sum += v
		}
	}
	return sum
}

func GetDirsInfo(input string) map[string]int {
	dirsInfo, dirsStack := GetDirsInfoAndDirsStack(input)
	TraverseStackAndUpdateDirsInfo(dirsInfo, dirsStack)
	return dirsInfo
}

func TraverseStackAndUpdateDirsInfo(dirsInfo map[string]int, dirsStack []Dir) {
	dirsStackSize := len(dirsStack)
	for dirsStackSize > 0 {
		dirInfoInStack := dirsStack[dirsStackSize-1]
		dirsStack = dirsStack[:dirsStackSize-1]
		dirsStackSize--

		dirsInfo[dirInfoInStack.Path] = getSizeOfDir(dirsInfo, dirInfoInStack)
	}
}

func getSizeOfDir(dirsInfo map[string]int, dirInfo Dir) int {
	sizeOfCurrentFolder := 0

	for _, line := range dirInfo.Content {
		lineSlice := strings.Split(line, " ")
		firstString := strings.TrimSpace(lineSlice[0])

		if firstString == "dir" {
			dirName := strings.TrimSpace(lineSlice[1])
			sizeOfCurrentFolder += dirsInfo[dirInfo.Path+defaultDirSeparator+dirName]
		} else {
			runeSlice := []rune(firstString)
			char := runeSlice[0]
			if unicode.IsDigit(char) {
				fileSize, err := strconv.Atoi(firstString)

				// I am assuming correct input avoiding handling errors (do not try this at home/production).
				if err != nil {
					log.Fatal(err)
				}

				sizeOfCurrentFolder += fileSize
			}
		}
	}
	return sizeOfCurrentFolder
}

func GetDirsInfoAndDirsStack(input string) (map[string]int, []Dir) {
	dirsInfo := make(map[string]int)
	var cwdStack []string
	var dirsStack []Dir

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		lineStr := strings.TrimSpace(lines[i])
		lineSlice := strings.Split(lineStr, " ")

		if strings.TrimSpace(lineSlice[0]) == "$" {
			if strings.TrimSpace(lineSlice[1]) == "cd" {
				if strings.TrimSpace(lineSlice[2]) == ".." {
					l := len(cwdStack)
					cwdStack = cwdStack[:l-1]
				} else {
					dirNameToMoveTo := strings.TrimSpace(lineSlice[2])
					cwdStack = append(cwdStack, dirNameToMoveTo)
				}
			} else if strings.TrimSpace(lineSlice[1]) == "ls" {
				var nItemsInDir int
				currentDirName := ConcatenatePaths(cwdStack, "/")
				nItemsInDir, dirsStack, dirsInfo = AnalyseLsCommand(currentDirName, dirsInfo, lines[i:], dirsStack)
				i += nItemsInDir

				// We need to substract one because we are already adding 1 to "i" in each iteration of the "for".
				i--
			}
		}
	}
	return dirsInfo, dirsStack
}

// "lines" includes "$ ls" as the first line
func AnalyseLsCommand(currentDirName string, dirsInfo map[string]int, lines []string, dirsStack []Dir) (nItemsInDir int, dirsStackOut []Dir, dirsInfoOut map[string]int) {
	nItemsInDir, currentDirContainsSubdir, dirsInfo := UpdateDirInfoIfItIsALeafDir(currentDirName, dirsInfo, lines[1:])
	if currentDirContainsSubdir {
		d := Dir{
			Path:    currentDirName,
			Content: lines[1 : nItemsInDir+1],
		}
		dirsStack = append(dirsStack, d)
	}
	return nItemsInDir, dirsStack, dirsInfo
}

func UpdateDirInfoIfItIsALeafDir(currentDirName string, dirsInfo map[string]int, lines []string) (nItemsInDir int, currentDirContainsSubdir bool, dirsInfoOut map[string]int) {
	if dirsInfo == nil {
		dirsInfo = make(map[string]int)
	}

	sizeOfCurrentFolder := 0
	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		firstString := strings.TrimSpace(lineParts[0])

		cliCommand := firstString == "$"
		if cliCommand {
			break
		}

		nItemsInDir++
		subdirectoryFound := firstString == "dir"
		if subdirectoryFound {
			currentDirContainsSubdir = true
		} else {
			runeSlice := []rune(firstString)
			char := runeSlice[0]
			if unicode.IsDigit(char) {
				fileSize, err := strconv.Atoi(firstString)
				if err != nil {
					log.Fatal(err)
				}
				sizeOfCurrentFolder += fileSize
			}
		}
	}

	if !currentDirContainsSubdir {
		dirsInfo[currentDirName] += sizeOfCurrentFolder
	}

	return nItemsInDir, currentDirContainsSubdir, dirsInfo
}

func ConcatenatePaths(path []string, delimiter string) string {
	if delimiter == "" {
		delimiter = defaultDirSeparator
	}
	var sb strings.Builder
	for _, v := range path {
		sb.WriteString(v + delimiter)
	}

	s := sb.String()
	return s[:len(s)-1]
}
