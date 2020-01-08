package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	csvr := csv.NewReader(file)

	correctCounter := 0
	questionCounter := 0

	for {
		row, err := csvr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		question := row[0]
		answer := row[1]

		fmt.Printf("%v = ", question)

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if text == answer {
			correctCounter++
		}
		questionCounter++
	}

	fmt.Printf("You got %v correct out of %v\n", correctCounter, questionCounter)
}
