package packet

import "errors"

// Detects a start-of-packet marker in the datastream.
// In the protocol being used by the Elves, the start of a packet is
// indicated by a sequence of four characters that are all different.
// Packet marker means string of 4 chars which are all different.
// left index is were the packet marker starts and right index is where it ends, in regars to the original input.
// e, g: in "mjqjpqmgbljsphdztnvjfqwrcgsmlb", we have the start of packet marker "jpqm" has
// left index == 3 and right index == 6
func DetectStartOfPacketMarker(input string) (int, error) {
	inputLenght := len(input)

	for i := 0; i < inputLenght; i++ {
		leftIndexOfStartOfPacketMarker := i
		rightIndexOfStartOfPacketMarker := i + 3

		if rightIndexOfStartOfPacketMarker >= inputLenght {
			return 0, errors.New("no start of packer marker found")
		}

		if areAllCharsDifferent(input[leftIndexOfStartOfPacketMarker : rightIndexOfStartOfPacketMarker+1]) {
			return rightIndexOfStartOfPacketMarker + 1, nil
		}
	}

	return 0, errors.New("no start of packer marker found")
}

func areAllCharsDifferent(in string) bool {
	visited := make(map[rune]struct{})
	for _, v := range in {
		_, isPresent := visited[v]
		if isPresent {
			return false
		}
		visited[v] = struct{}{}
	}
	return true
}
