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
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "Time limit for quiz in seconds")
	flag.Parse()
	// the above thing helps when we pass a binary to user if you try and run the binary with a --help flag it will give out the
	// message that we specify above
	// you could also add a -csv="abc.csv" flag if you want to open a different csv
	// "" are not really needed but if there is a space in the name and you dont use "" it will give an error
	// similarly you can use -limit to set a user defined time limit

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Pailed to parse the provided CSV file")
	}
	// fmt.Println(lines)
	problems := parseLines(lines)
	// fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()
		select {
		case <-timer.C:
			// fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			// return
			// this would work but in order to avoid duplicates we can break by introducing break segment
			fmt.Println()
			break problemloop
		case answer := <-ansCh:
			if answer == p.a {
				fmt.Println("Correct!")
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			// a: line[1],
			// this would work normally but in case the csv was passed to us with spaces before the answer
			// it can be corrected by using strings.TrimSpace(line[1])
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
