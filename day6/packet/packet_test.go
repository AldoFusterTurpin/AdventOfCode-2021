package packet_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day6/packet"
)

func TestDetectStartOfPacketMarker(t *testing.T) {
	type TestData struct {
		input          string
		expectedResult int
		expectedErr    error
	}

	tests := map[string]TestData{
		"sample_input": TestData{
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expectedResult: 7,
			expectedErr:    nil,
		},
		"sample_input_2": TestData{
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			expectedResult: 6,
			expectedErr:    nil,
		},
		"sample_input_3": TestData{
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expectedResult: 10,
			expectedErr:    nil,
		},
		"sample_input_4": TestData{
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expectedResult: 11,
			expectedErr:    nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := packet.DetectStartOfPacketMarker(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.expectedResult {
				t.Fatalf("expected %v, but got %v", tc.expectedResult, got)
			}
		})
	}

}
