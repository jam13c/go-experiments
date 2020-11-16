package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	file    = flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	limit   = flag.Int("limit", 30, "Time limit for the quiz in seconds")
	shuffle = flag.Bool("shuffle", false, "Shuffle the questions")
)

func main() {
	flag.Parse()

	questions := getQuestions(*file)
	if *shuffle {
		shuffleQuestions(questions)
	}

	fmt.Printf("This quiz has %d questions and you have %d seconds to complete it. Hit enter to start", len(questions), *limit)
	fmt.Scanf("\n")

	timeout := time.After(time.Duration(*limit) * time.Second)

	correct := 0
	for i, question := range questions {
		ansCh := makeQuestion(i, question.q)

		select {
		case ans := <-ansCh:
			if ans == question.a {
				correct++
			}
		case <-timeout:
			fmt.Printf("\nTime is up. You scored %d out of %d", correct, len(questions))
			return
		}
	}
	fmt.Printf("\nYou scored %d out of %d", correct, len(questions))

}

func shuffleQuestions(q []question) {
	for i := len(q) - 1; i > 0; i-- {
		j := rand.Intn(i)
		q[i], q[j] = q[j], q[i]
	}
}

func makeQuestion(index int, text string) <-chan string {
	c := make(chan string)
	fmt.Printf("\nProblem %d: %s = ", index+1, text)
	go func() {
		var a string
		fmt.Scanf("%s\n", &a)
		c <- a
	}()
	return c
}

func getQuestions(file string) []question {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Failed to open file. (%s)\n", err.Error())
		os.Exit(1)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read csv. (%s)\n", err.Error())
		os.Exit(1)
	}
	q := make([]question, len(records))
	for i, record := range records {

		q[i] = question{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}

	}
	return q
}

type question struct {
	q string
	a string
}
