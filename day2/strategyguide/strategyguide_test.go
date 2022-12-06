package strategyguide_test

import (
	"reflect"
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day2/strategyguide"
)

func TestPlayRound(t *testing.T) {
	type testData struct {
		opponentsChoice    rune
		yourChoice         rune
		expectedRoundScore int
	}

	tests := map[string]testData{
		"example": {
			opponentsChoice:    'A',
			yourChoice:         'Y',
			expectedRoundScore: 8,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := strategyguide.PlayRound(tc.opponentsChoice, tc.yourChoice)

			if !reflect.DeepEqual(tc.expectedRoundScore, got) {
				t.Fatalf("expected %v, but got %v", tc.expectedRoundScore, got)
			}
		})
	}
}
