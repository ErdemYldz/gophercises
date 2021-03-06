package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const printTemplate = `The number of answered questions: %v
The number of correct answer: %v
The number of uncorrect answer: %v`

type record struct {
	question string
	answer   string
}
type result struct {
	correct   int
	uncorrect int
}

func main() {
	filepath := flag.String("fp", "problems.csv", "file path")
	totalTime := flag.Int("t", 10, "total exam duration in Seconds")
	flag.Parse()

	f, err := os.Open(*filepath)
	if err != nil {
		log.Fatalln("error while opening the file: ", err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	outcome := result{}
	done := make(chan bool)
	start := time.Now()
	go getRecords(reader, done, &outcome)

	timeIt(*totalTime, done)
	printResults(&outcome)
	ends := time.Since(start)
	fmt.Println("it takes ", ends)
}

func timeIt(t int, ch chan bool) {
	defer close(ch)
	select {
	case <-ch:
		return
	case <-time.Tick(time.Duration(t) * time.Second):
		fmt.Println()
		fmt.Println("The time is up!!")
		return
	}
}

func getRecords(reader *csv.Reader, ch chan bool, outcome *result) {
	for {
		columns, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error while getting the row: ", err)
		}
		var row record
		row.question = columns[0]
		row.answer = strings.TrimSpace(columns[1])
		userAnswer := askQuestion(row.question)
		if row.answer == userAnswer {
			outcome.correct++
		} else {
			outcome.uncorrect++
		}
	}
	ch <- true
}

func askQuestion(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s: ", question)
	scanner.Scan()
	answer := scanner.Text()
	answer = strings.TrimSpace(answer)
	return answer
}

func printResults(outcome *result) {
	fmt.Println()
	fmt.Printf(printTemplate, outcome.correct+outcome.uncorrect,
		outcome.correct, outcome.uncorrect)
	fmt.Println()
}
