package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func read(filename string) [][]string {
	file, _ := os.Open(filename)
	reader := csv.NewReader(file)

	lines, _ := reader.ReadAll()

	return lines
}

func parseProblems(problems [][]string) []problem {
	var record []problem
	for _, a := range problems {
		var p problem
		p.question = a[0]
		p.answer = a[1]
		record = append(record, p)
	}

	return record
}

func game(problems []problem, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	var correct int
	for _, problem := range problems {
		fmt.Printf("%s > ", problem.question)
		answerCh := make(chan string)
		go func() {
			var userAns string
			fmt.Scanf("%s\n", &userAns)
			answerCh <- userAns
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				fmt.Println("You are smart")
				correct++
			} else {
				fmt.Println("Try again!")
			}
		}
	}
	if correct == len(problems) {
		fmt.Println("You won!")
	}
}

func main() {
	filename := flag.String("csv", "problems.csv", "the filename with problems")
	limit := flag.Int("limit", 30, "the time for quiz in seconds")
	flag.Parse()
	lines := read(*filename)
	problems := parseProblems(lines)
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
	game(problems, *limit)
}
