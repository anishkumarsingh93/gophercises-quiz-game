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

//Part1 of the Quiz game
func main() {

	//Reading the csv filename using flags
	csvFileName := flag.String("csv", "problems.csv", "A csv fle containing the quiz in <question,answer> format")
	//Reading time duration in seconds
	timeLimit := flag.Int("duration", 30, "Time duration of the quiz in seconds")
	flag.Parse()

	//Opening the file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Could not open csv file: %s\n", *csvFileName))
	}

	//Reading the whole file at a time as csv file wont be large
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Could not read the csv file" + err.Error()))
	}

	//Getting the problems
	problems := getProblems(lines)

	timer := time.After(time.Duration(*timeLimit) * time.Second)

	//Timed Q&A
	correct := 0
	for i, p := range problems {

		//channel to get answer
		ansChan := make(chan string)
		//Printing the problem
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		//go routine for reading the answers
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansChan <- ans
		}()

		//Waiting for timer channel to respond
		select {
		case <-timer:
			//Final Result
			fmt.Printf("\nYou scored %d out of %d!\n", correct, len(problems))
			return
		case answer := <-ansChan:
			if answer == p.a {
				correct++
			}
		}
	}
	//If user ompletes the quiz  before  time runs out
	fmt.Printf("You scored %d out of %d!\n", correct, len(problems))
}

//getProblems parses the the file data into a list of problem struct
func getProblems(lines [][]string) []problem {
	//Its better to make a defnes slice and not use append as we already know the length. More performance
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), //trimming extra spaces from csv
		}
	}
	return result
}

//exit prints a message and exits the program
func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
