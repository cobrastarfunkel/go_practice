package main

import (
	"fmt"
	"os"

	"example.com/quiz_game/game"
	QuestionBank "example.com/quiz_game/questions"
)

func playAgain() {
	fmt.Println("Play again? [y/n]")
	var answer string
	fmt.Scanln(&answer)
	if answer == "y" {
		run()
	} else if answer == "n" {
		os.Exit(0)
	} else {
		playAgain()
	}
}

func gameLoop(curGame *game.Game) {
	totalQs := len(curGame.Questions.Questions)
	for i := 0; i < totalQs; i++ {
		curGame.AskQuestion()
	}

	fmt.Printf("Final Score: %d\n", curGame.CurrentScore)
	playAgain()
}

func loadQuestions(filename string) {
	qb := QuestionBank.NewQuestionBank(filename)
	curGame := game.Game{
		Questions:     qb,
		CurrentScore:  200,
		MaxQuestions:  len(qb.Questions),
		AmountWagered: 0,
	}
	gameLoop(&curGame)
}

func run() {
	fmt.Println("Enter the name of the quiz file: ")
	var filename string
	fmt.Scanln(&filename)

	loadQuestions(filename)
}

func main() {
	run()
}
