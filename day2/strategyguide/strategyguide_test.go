package strategyguide_test

import (
	"reflect"
	"testing"
)

func TestPlayRound(t *testing.T) {
	type testData struct {
		opponentsChoice    string
		yourChoice         string
		expectedRoundScore int
	}

	tests := map[string]testData{
		"example": {
			opponentsChoice: "A",
			yourChoice:      "Y",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PlayRound(tc.opponentsChoice, tc.yourChoice)

			if !reflect.DeepEqual(tc.expectedRoundScore, got) {
				t.Fatalf("expected %v, but got %v", tc.expectedRoundScore, got)
			}
		})
	}
}
