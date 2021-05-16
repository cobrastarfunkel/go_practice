package game

import (
	"fmt"

	"example.com/quiz_game/questions"
)

type Game struct {
	Questions     *questions.QuestionBank
	AmountWagered int
	CurrentScore  int
	MaxQuestions  int
}

func (g *Game) setupWager() {
	fmt.Printf("\nYour Current Balance is: %d\nHow Much would you like to wager? ", g.CurrentScore)
	var curWager int
	fmt.Scanln(&curWager)
	g.SetWager(curWager)
}

func getGuess(q questions.Question, incorrectGuesses *int) string {
	fmt.Println(q.Text)
	var guess string
	fmt.Scanln(&guess)

	if !q.AnswerContainsDigit([]rune(guess)[0]) {
		*incorrectGuesses += 1
	}

	return guess
}

func (g *Game) AskQuestion() int {
	answeredCorrectly := false
	incorrectGuesses := 0
	curQ := g.Questions.Questions[0]
	var curAnswer string
	g.setupWager()

	for incorrectGuesses < 4 {
		curAnswer += getGuess(curQ, &incorrectGuesses)
		curAnswer = curQ.GetAnswerWithPlaceholder(curAnswer)

		if curQ.AllDigitsGuessed(curAnswer) {
			answeredCorrectly = true
			fmt.Printf("Congratulations %s was the answer!\n", curQ.Answer)
			break
		}

		fmt.Printf("Current Answers: %s\n", curAnswer)
		fmt.Printf("Incorrect Guesses: %d\n\n", incorrectGuesses)
	}
	if !answeredCorrectly {
		fmt.Printf("Sorry the correct Answer was: %s", curQ.Answer)
	}
	g.AdjustPoints(answeredCorrectly)
	g.Questions.Questions = g.Questions.Questions[1:]

	return g.CurrentScore
}

func (g *Game) AdjustPoints(corretGuess bool) {
	if corretGuess {
		g.CurrentScore += g.AmountWagered
	} else {
		g.CurrentScore -= g.AmountWagered
	}
}

func (g *Game) SetWager(wager int) {
	if wager <= g.CurrentScore && wager > -1 {
		g.AmountWagered = wager
	} else {
		g.AmountWagered = g.CurrentScore
	}
}
