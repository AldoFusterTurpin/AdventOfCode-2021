package supplystacks

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Crate represents a crate. In the below example,
// [D]
// [N] [C]
// [Z] [M] [P]
//  0   1   2
// we find:
// - One crate with letter D in stack 0
// - One crate with letter N in stack 0
// - One crate with letter Z in stack 0
// - One crate with letter C in stack 1
// - One crate with letter M in stack 1
// - One crate with letter P in stack 2
// The puzzle input stacks are not zero indexed but we will treat them as they were.
type Crate struct {
	letter string
}

// Every time we find a char that is a letter, divide the index of it
// by the number of columns we have and round it up.
func SolveProblem(input string) string {
	splitted := strings.Split(input, "\n\n")
	cratesDrawing := splitted[0]

	lines := strings.Split(cratesDrawing, "\n")

	linesLen := float64(len(lines[0]))
	nStacks := math.Ceil(linesLen / 4)

	// stacks[i] represents the stack number i of Crates and
	// a stack is of type []Crate.
	stacks := createStacks(lines, int(nStacks))

	// Wee need to reverse the stacks because we read the input text
	// from top to bottom (left to right). We could also add the elements to the head
	// while reading the initial stacks but I prefer this approach as append() is optimized.
	// When Go performs 'append' it reserves the double of capacity of the needed, something that is not
	// done if we add the elements to the head.
	// It is also cleaner and simple for me.
	stacks = reverseStacks(stacks)

	rearrangment := splitted[1]
	stacks = performRearrangement(rearrangment, stacks)

	return getMessageFromTopOfStacks(stacks)
}

func performRearrangement(rearrangmentInstructions string, stacks [][]Crate) [][]Crate {
	for _, line := range strings.Split(rearrangmentInstructions, "\n") {
		lineSlice := splitLine(line)
		// fmt.Println(lineSlice)

		n, from, to := extractInfoFromLineSlice(lineSlice)
		stacks = moveCrate(stacks, n, from-1, to-1)
	}
	return stacks
}

func extractInfoFromLineSlice(lineSplitted []string) (int, int, int) {
	nStr := lineSplitted[0]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		log.Fatal(err)
	}

	fromStr := lineSplitted[1]
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		log.Fatal(err)
	}

	toStr := lineSplitted[2]
	to, err := strconv.Atoi(toStr)
	if err != nil {
		log.Fatal(err)
	}
	return n, from, to
}

// given the input "move 1 from 9 to 3"
// it returns the slice
// [1, 9, 3]
func splitLine(line string) []string {
	fmt.Println(line)
	const sentinel = "-"
	line = strings.ReplaceAll(line, "move ", "")
	line = strings.ReplaceAll(line, " from ", sentinel)
	line = strings.ReplaceAll(line, " to ", sentinel)
	line = strings.TrimSpace(line)

	return strings.Split(line, sentinel)
}

func getMessageFromTopOfStacks(stacks [][]Crate) string {
	var sb strings.Builder
	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		element := stack[len(stack)-1]
		sb.WriteString(element.letter)
	}

	return sb.String()
}

func printStacks(stacks [][]Crate) {
	for i, stack := range stacks {
		fmt.Printf("stack %v of type %T: %v\n", i, stack, stack)
	}
}

func reverseStacks(stacks [][]Crate) [][]Crate {
	for i := 0; i < len(stacks); i++ {
		stacks[i] = reverseSlice(stacks[i])
	}
	return stacks
}

func reverseSlice(stack []Crate) []Crate {
	var tmp []Crate
	for i := len(stack) - 1; i >= 0; i-- {
		tmp = append(tmp, stack[i])
	}
	return tmp
}

func createStacks(lines []string, n int) [][]Crate {
	stacks := make([][]Crate, n)
	for i := 0; i < n; i++ {
		stacks[i] = make([]Crate, 0)
	}

	for _, line := range lines {
		for j, c := range line {
			// We need to substract one because the index in the statement is NOT zero indexed
			processCharacter(c, j, stacks)
		}
	}

	return stacks
}

func processCharacter(character rune, j int, stacks [][]Crate) {
	if unicode.IsLetter(character) {
		crate := Crate{
			letter: string(character),
		}

		indexFloat := math.Ceil(float64(j) / 4)
		index := int(indexFloat)

		// stacks in the problem statement/input are not zero-indexed.
		index--
		stacks[index] = append(stacks[index], crate)
	}
}

func moveCrate(stacks [][]Crate, n, from, to int) [][]Crate {
	fmt.Printf("Move n: %v, from: %v, to: %v\n", n, from, to)

	// fmt.Println("Before moving:")
	// printStacks(stacks)

	for i := 0; i < n; i++ {
		elementToMove := stacks[from][len(stacks[from])-1]

		stacks[from] = stacks[from][:len(stacks[from])-1]
		stacks[to] = append(stacks[to], elementToMove)

	}

	// fmt.Println("\nAfter moved:")
	// printStacks(stacks)
	// fmt.Println("-----------------------------------------------------")

	return stacks
}
