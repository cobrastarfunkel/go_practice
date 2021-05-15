package main

import (
	"fmt"

	question "example.com/quiz_game/questions"
	questionBank "example.com/quiz_game/questions"
)

func main() {
	q := question.Question{Text: "Test", Answer: "1234"}
	fmt.Println(q.GetAnswerWithPlaceholder("1234"))

	qb := questionBank.NewQuestionBank("./test.txt")
	qb.PrintQuestions()
}
