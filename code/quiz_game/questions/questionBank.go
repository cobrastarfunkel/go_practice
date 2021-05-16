package questions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type QuestionBank struct {
	Questions     []Question
	questionsFile string
}

func (q *QuestionBank) loadQuestions() {
	file, err := os.Open(q.questionsFile)
	defer file.Close()

	if err != nil {
		fmt.Println("File Not Found!")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	for _, line := range txtLines {
		items := strings.SplitN(line, " ", 2)
		q.Questions = append(q.Questions, Question{Text: items[1], Answer: items[0]})
	}
}

func NewQuestionBank(qFile string) *QuestionBank {
	q := new(QuestionBank)
	q.questionsFile = qFile
	q.loadQuestions()

	return q
}

func (q QuestionBank) PrintQuestions() {
	for _, q := range q.Questions {
		fmt.Println(q.Text)
	}
}
