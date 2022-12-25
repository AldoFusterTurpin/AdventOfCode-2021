package strategyguide_test

import (
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day2/strategyguide"
)

func TestApplyStrategyGuide(t *testing.T) {
	type testData struct {
		input              string
		expectedTotalScore int
	}

	tests := map[string]testData{
		"example_1": {
			input: `A Y
B X
C Z`,
			expectedTotalScore: 15,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := strategyguide.ApplyStrategyGuide(tc.input)
			if err != nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if tc.expectedTotalScore != got {
				t.Fatalf("expected %v, but got %v", tc.expectedTotalScore, got)
			}
		})
	}
}

func TestApplyStrategyGuidePart2(t *testing.T) {
	type testData struct {
		input              string
		expectedTotalScore int
	}

	tests := map[string]testData{
		"example_1": {
			input: `A Y
B X
C Z`,
			expectedTotalScore: 12,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := strategyguide.ApplyStrategyGuidePart2(tc.input)
			if err != nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if tc.expectedTotalScore != got {
				t.Fatalf("expected %v, but got %v", tc.expectedTotalScore, got)
			}
		})
	}
}

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

func TestGetTotalScore(t *testing.T) {
	type testData struct {
		scoreRounds        []int
		expectedTotalScore int
	}

	tests := map[string]testData{
		"example_1": {
			scoreRounds:        []int{1, 1, 1, 2},
			expectedTotalScore: 5,
		},
		"empty_array": {
			scoreRounds:        []int{},
			expectedTotalScore: 0,
		},
		"example_from_statement": {
			scoreRounds:        []int{8, 1, 6},
			expectedTotalScore: 15,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotTotalScore := strategyguide.GetTotalScore(tc.scoreRounds)

			if tc.expectedTotalScore != gotTotalScore {
				t.Fatalf("expected %v, but got %v", tc.expectedTotalScore, gotTotalScore)
			}
		})
	}
}
