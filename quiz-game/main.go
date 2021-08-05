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
	q string
	a string
}
func main()  {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil{
		exit(fmt.Sprintf("Failed to open file %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit(fmt.Sprintf("Faild to parse file %s\n", *csvFilename))
	}
	problems := parseLines(lines)

	// this will send a message over the channel when the time is done for once
	// there is also "ticker" which will send a message everytime we complete the time
	// read about it here https://yourbasic.org/golang/time-reset-wait-stop-timeout-cancel-interval/
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctCount := 0
	for i,p := range problems {
		fmt.Printf("Problem #%d: %s =", i+1, p.q)
		answerCh := make(chan string)
		go func(answerCh chan string) {
			var answer string
			// Scanf blocks the program until it gets a result.
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}(answerCh)

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d problems!\n", correctCount, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a  {
				correctCount++
			}
		}
	}
	fmt.Printf("You scored %d out of %d problems!\n", correctCount, len(problems))
}

func parseLines(lines [][]string)  []problem {
	// whenever you know the size of your data use make(), it will better than using append()
	ret := make([]problem, len(lines))

	for i,line :=range lines{
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}