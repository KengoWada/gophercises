package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)

	rows, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse csv file")
	}

	problems := parseRows(rows)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correctAnswers := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.answer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("You got %d correct answers out %d questions\n", correctAnswers, len(problems))
}

func parseRows(rows [][]string) []problem {
	problems := make([]problem, len(rows))
	for i, row := range rows {
		problems[i] = problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}

	return problems
}

type problem struct {
	question string
	answer   string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
