package packet

import "errors"

// For Part 2
func DetectStartOfMessageMarker(input string) (int, error) {
	inputLenght := len(input)

	for i := 0; i < inputLenght; i++ {
		leftIndexOfStartOfMessageMarker := i
		rightIndexOfStartOfMessageMarker := i + 13

		if rightIndexOfStartOfMessageMarker >= inputLenght {
			return 0, errors.New("no start of packer marker found")
		}

		if areAllCharsDifferent(input[leftIndexOfStartOfMessageMarker : rightIndexOfStartOfMessageMarker+1]) {
			return rightIndexOfStartOfMessageMarker + 1, nil
		}
	}

	return 0, errors.New("no start of packer marker found")
}
