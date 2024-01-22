package domain

// countLetterInWord counts how many times a letter appears in a word
func countLetterInWord(word, letter *string) (count *int) {
	count = new(int)

	for _, l := range *word {
		if string(l) == *letter {
			*count++
		}
	}

	return count
}

// countLetterInCorrectPosition counts how many times a letter appears in a word in the correct position
func countLetterInCorrectPosition(reference, word, letter *string) (count *int) {
	count = new(int)

	for i, l := range *word {
		if string(l) == *letter && *(checkLetterInCorrectPosition(reference, letter, &i)) {
			*count++
		}
	}

	return count
}

// countLetterInAnotherPosition counts how many times a letter appears in a word in another position
func countLetterInAnotherPosition(checkWord []CheckWord) (count *int) {
	count = new(int)

	for _, l := range checkWord {
		if l.Check != nil && *l.Check == Elsewhere {
			*count++
		}
	}

	return count
}

// checkLetterInCorrectPosition checks if a letter is in the correct position
func checkLetterInCorrectPosition(word, letter *string, position *int) (correct *bool) {
	correct = new(bool)

	if len(*word) <= *position {
		return correct
	}

	if string((*word)[*position]) == *letter {
		*correct = true
	}

	return correct
}
