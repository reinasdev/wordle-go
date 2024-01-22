package domain

import (
	"wordle/infra"
	"wordle/utils"
)

// CheckWord represents a letter check
type CheckWord struct {
	Letter *string      `json:"letter"`
	Check  *CheckLetter `json:"check"`
}

// CheckLetter represents a letter check type
type CheckLetter string

// CheckLetter types
const (
	Correct   CheckLetter = "correct"
	Wrong     CheckLetter = "wrong"
	Elsewhere CheckLetter = "elsewhere"
)

// VerifyWord checks if a word is correct
func VerifyWord(word *string) (attempt []CheckWord, success *bool, err error) {
	attempt = make([]CheckWord, len(*word))
	success = utils.GetPointer(true)

	dailyWord, err := infra.GetDailyWord()
	if err != nil {
		return nil, nil, err
	}

	for i, letter := range *word {
		attempt[i].Letter = utils.GetPointer(string(letter))

		if *(checkLetterInCorrectPosition(dailyWord, attempt[i].Letter, &i)) {
			attempt[i].Check = utils.GetPointer(Correct)
			continue
		} else {
			success = utils.GetPointer(false)

			if *countLetterInWord(dailyWord, attempt[i].Letter) > (*countLetterInCorrectPosition(dailyWord, word, attempt[i].Letter) + *countLetterInAnotherPosition(attempt)) {
				attempt[i].Check = utils.GetPointer(Elsewhere)
				continue
			}
		}

		attempt[i].Check = utils.GetPointer(Wrong)
	}

	return attempt, success, nil
}
