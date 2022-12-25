package strategyguide

import (
	"errors"
	"strings"
)

// Part 2, meaning of the second column
const youNeedToLoose = 'X'
const youNeedToDraw = 'Y'
const youNeedToWin = 'Z'

func ApplyStrategyGuidePart2(input string) (int, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	nRounds := len(lines)
	scoreRounds := make([]int, nRounds)

	for i, line := range lines {
		var err error
		scoreRounds, err = treatLinePart2(line, scoreRounds, i)
		if err != nil {
			return 0, err
		}
	}

	totalScore := GetTotalScore(scoreRounds)
	return totalScore, nil
}

func treatLinePart2(line string, scoreRounds []int, index int) ([]int, error) {
	parts := strings.Split(line, " ")

	// a string with just one character, but still a string, needs the conversion below
	opponentsChoiceStr := parts[0]
	opponentsChoiceRune := rune(opponentsChoiceStr[0])

	resultRoundStr := parts[1]
	resultRoundRune := rune(resultRoundStr[0])

	yourChoiceRune, err := whatShouldYouPlay(opponentsChoiceRune, resultRoundRune)
	if err != nil {
		return nil, err
	}

	scoreRound, err := GetScoreRound(opponentsChoiceRune, yourChoiceRune)
	if err != nil {
		return nil, err
	}
	scoreRounds[index] = scoreRound
	return scoreRounds, nil
}

// used in Part 2
func whatShouldYouPlay(opponentsChoice, resultRound rune) (rune, error) {
	if opponentsChoice == opponentsRock {
		if resultRound == youNeedToDraw {
			return yourRock, nil
		}
		if resultRound == youNeedToLoose {
			return yourScissors, nil
		}
		if resultRound == youNeedToWin {
			return yourPaper, nil
		} else {
			return 0, errors.New("invalid round")
		}
	}
	if opponentsChoice == opponentsScissors {
		if resultRound == youNeedToDraw {
			return yourScissors, nil
		}
		if resultRound == youNeedToLoose {
			return yourPaper, nil
		}
		if resultRound == youNeedToWin {
			return yourRock, nil
		} else {
			return 0, errors.New("invalid round")
		}
	}
	if opponentsChoice == opponentsPaper {
		if resultRound == youNeedToDraw {
			return yourPaper, nil
		}
		if resultRound == youNeedToLoose {
			return yourRock, nil
		}
		if resultRound == youNeedToWin {
			return yourScissors, nil
		} else {
			return 0, errors.New("invalid round")
		}
	}

	return 0, errors.New("invalid round")
}
