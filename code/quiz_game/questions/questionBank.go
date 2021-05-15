package questions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type questionBank struct {
	questions     []Question
	questionsFile string
}

func (q *questionBank) loadQuestions() {
	file, err := os.Open(q.questionsFile)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	for _, line := range txtLines {
		items := strings.SplitN(line, " ", 2)
		q.questions = append(q.questions, Question{Text: items[1], Answer: items[0]})
	}
}

func NewQuestionBank(qFile string) *questionBank {
	q := new(questionBank)
	q.questionsFile = qFile
	q.loadQuestions()

	return q
}

func (q questionBank) PrintQuestions() {
	for _, q := range q.questions {
		fmt.Println(q.Text)
	}
}
