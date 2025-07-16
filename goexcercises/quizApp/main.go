package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	timerflag := flag.Int("timelimit", 30, "Timer for quiz")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file: %s\n", err))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read csv file.")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timerflag) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Question #%d:%s= ", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if p.answer == answer {
			correct++
		} else {
			fmt.Println("Incorrect answer")
		}
	}
	fmt.Printf("Your final score is: %d/%d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
