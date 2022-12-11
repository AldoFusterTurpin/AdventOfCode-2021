package strategyguide_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day2/strategyguide"
)

func TestGetScoreRound(t *testing.T) {
	type testData struct {
		opponentsChoice    rune
		yourChoice         rune
		expectedRoundScore int
		expectedError      error
	}

	tests := map[string]testData{
		"example_1": {
			opponentsChoice:    'A',
			yourChoice:         'Y',
			expectedRoundScore: 8,
			expectedError:      nil,
		},
		"example_2": {
			opponentsChoice:    'B',
			yourChoice:         'X',
			expectedRoundScore: 1,
			expectedError:      nil,
		},
		"example_3": {
			opponentsChoice:    'C',
			yourChoice:         'Z',
			expectedRoundScore: 6,
			expectedError:      nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotValue, gotError := strategyguide.GetScoreRound(tc.opponentsChoice, tc.yourChoice)

			if tc.expectedRoundScore != gotValue {
				t.Fatalf("expected %v, but got %v", tc.expectedRoundScore, gotValue)
			}

			if tc.expectedError != gotError {
				t.Fatalf("expected %v, but got %v", tc.expectedRoundScore, gotError)
			}
		})
	}
}
