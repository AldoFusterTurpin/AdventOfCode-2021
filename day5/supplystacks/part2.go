package supplystacks

import (
	"fmt"
	"math"
	"strings"
)

func SolveProblemPart2(input string) string {
	splitted := strings.Split(input, "\n\n")
	cratesDrawing := splitted[0]

	lines := strings.Split(cratesDrawing, "\n")

	linesLen := float64(len(lines[0]))
	nStacks := math.Ceil(linesLen / 4)

	stacks := createStacks(lines, int(nStacks))

	stacks = reverseStacks(stacks)

	rearrangment := splitted[1]
	stacks = performRearrangementPart2(rearrangment, stacks)

	return getMessageFromTopOfStacks(stacks)
}

func performRearrangementPart2(rearrangmentInstructions string, stacks [][]Crate) [][]Crate {
	for _, line := range strings.Split(rearrangmentInstructions, "\n") {
		lineSlice := splitLine(line)

		n, from, to := extractInfoFromLineSlice(lineSlice)

		stacks = moveCratePart2(stacks, n, from-1, to-1)
	}
	return stacks
}

func moveCratePart2(stacks [][]Crate, n, from, to int) [][]Crate {
	fmt.Printf("Move n: %v, from: %v, to: %v\n", n, from, to)

	fmt.Println("Before moving:")
	printStacks(stacks)

	// We need to create a tmp stack because "the ability to pick up and move multiple crates at once." in
	// (the problem statement) means that we need to preserve the order of the original stack, and we can do that
	// easily moving first the items to a temporal stack and then move them from the temporal stack to the
	// target stack.
	var tmp []Crate
	for i := 0; i < n; i++ {
		elementToMove := stacks[from][len(stacks[from])-1]
		stacks[from] = stacks[from][:len(stacks[from])-1]
		tmp = append(tmp, elementToMove)
	}

	for i := 0; i < n; i++ {
		elementToMove := tmp[len(tmp)-1]

		tmp = tmp[:len(tmp)-1]
		stacks[to] = append(stacks[to], elementToMove)

	}

	fmt.Println("\nAfter moved:")
	printStacks(stacks)
	fmt.Println("-----------------------------------------------------")

	return stacks
}
