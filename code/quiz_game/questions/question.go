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

func (q Question) AllDigitsGuessed(answer string) bool {
	if strings.ContainsRune(answer, '_') {
		return false
	}
	return true
}

func (q Question) GetAnswerWithPlaceholder(answer string) string {
	var retAnswer []rune

	for i := 0; i < len(q.Answer); i++ {
		retAnswer = append(retAnswer, '_')
	}

	for _, char := range answer {
		for j, ansChar := range q.Answer {
			if char == ansChar {
				retAnswer[j] = char
			}
		}
	}
	return string(retAnswer)
}
