package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	correct := 0
	problems := parseLines(lines)
	for index, problem := range problems {
		fmt.Printf("Problem %d: %s\n", index, problem.question)
		var input string
		fmt.Scanln(&input)
		if input == problem.answer {
			correct++
		}
	}
	fmt.Printf("You scored %v out of %v\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for index, line := range lines {
		problems[index] = problem{question: line[0], answer: line[1]}
	}
	return problems
}

type problem struct {
	question string
	answer   string
}
