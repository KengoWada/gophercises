package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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
	correctAnswers := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correctAnswers++
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
