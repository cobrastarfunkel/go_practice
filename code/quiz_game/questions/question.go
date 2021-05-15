package questions

import (
	"strings"
)

type Question struct {
	Text   string
	Answer string
}

func (q Question) AnswerContainsDigit(digit rune) bool {
	if strings.ContainsRune(q.Answer, digit) {
		return true
	}
	return false
}

func (q Question) allDigitsGuessed(answer string) bool {
	if strings.ContainsRune(answer, '_') {
		return false
	}
	return true
}

func (q Question) GetAnswerWithPlaceholder(answer string) string {
	retAnswer := []rune(answer)

	for i, char := range q.Answer {
		if char == []rune(answer)[i] {
			retAnswer[i] = char
		} else {
			retAnswer[i] = '_'
		}
	}
	if q.allDigitsGuessed(string(retAnswer)) {
		return "Congratulations! You guessed all the numbers."
	}
	return string(retAnswer)
}
